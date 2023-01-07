package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/poorfrombabylon/chargeMeBackend/internal/config"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type LocationDTOJson struct {
	PlaceID                      int                `json:"id,omitempty"`
	Name                         string             `json:"name"`
	Access                       int                `json:"access,omitempty"`
	Address                      string             `json:"address,omitempty"`
	Latitude                     float32            `json:"latitude"`
	Longitude                    float32            `json:"longitude"`
	Score                        *float32           `json:"score,omitempty"`
	IconType                     *string            `json:"icon_type,omitempty"`
	Description                  *string            `json:"description,omitempty"`
	AccessRestriction            *string            `json:"access_restriction,omitempty"`
	AccessRestrictionDescription *string            `json:"access_restriction_description,omitempty"`
	Cost                         *bool              `json:"cost,omitempty"`
	CostDescription              *string            `json:"cost_description,omitempty"`
	Hours                        *string            `json:"hours,omitempty"`
	Open247                      *bool              `json:"open247,omitempty"`
	ComingSoon                   *bool              `json:"coming_soon,omitempty"`
	PhoneNumber                  *string            `json:"e164_phone_number"`
	Stations                     []StationsDTOJson  `json:"stations"`
	Reviews                      []ReviewDTOJson    `json:"reviews"`
	Amenities                    []AmenitiesDTOJson `json:"amenities"`
	Photos                       []PhotoDTOJson     `json:"photos"`
}

type StationsDTOJson struct {
	Id              int             `json:"id"`
	LocationID      int             `json:"location_id"`
	Available       *int            `json:"available,omitempty"`
	Cost            *int            `json:"cost,omitempty"`
	Name            *string         `json:"name,omitempty"`
	Manufacturer    *string         `json:"manufacturer,omitempty"`
	CostDescription *string         `json:"cost_description,omitempty"`
	Hours           *string         `json:"hours,omitempty"`
	Kilowatts       *float32        `json:"kilowatts,omitempty"`
	Outlets         []OutletDTOJson `json:"outlets"`
}

type OutletDTOJson struct {
	Id            int      `json:"id"`
	ConnectorType int      `json:"connector_type"`
	Kilowatts     *float32 `json:"kilowatts,omitempty"`
	Power         int      `json:"power"`
}

type AmenitiesDTOJson struct {
	LocationID int  `json:"location_id,omitempty"`
	Form       *int `json:"type,omitempty"`
}

type ReviewDTOJson struct {
	StationID     int         `json:"station_id"`
	OutletID      int         `json:"outlet_id"`
	Comment       *string     `json:"comment,omitempty"`
	Rating        *int        `json:"rating,omitempty"`
	ConnectorType *int        `json:"connector_type"`
	Kilowatts     *float32    `json:"kilowatts"`
	VehicleName   *string     `json:"vehicle_name"`
	VehicleType   *int        `json:"vehicle_type"`
	User          UserDTOJSOn `json:"user"`
	CreatedAT     string      `json:"created_at"`
}

type UserDTOJSOn struct {
	FirstName *string `json:"first_name"`
}

type PhotoDTOJson struct {
	Caption *string `json:"caption"`
	ID      int     `json:"id"`
}

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("error while init config")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect to database:", err.Error())
	} else {
		log.Println("connected to db")
	}

	dbx := sqlx.NewDb(db, "pgx")
	libDBWrapper := libdb.NewSQLXDB(dbx)

	storageRegistry := storage.NewStorageRegistry(libDBWrapper)

	err = startJob(ctx, storageRegistry)
	if err != nil {
		log.Fatal("failure:", err.Error())
	}
}

func startJob(ctx context.Context, storageRegistry *storage.Storages) error {
	var dto []LocationDTOJson

	jsonFile, err := os.Open("/Users/almazkhayrullin/Desktop/full.json")
	if err != nil {
		log.Fatal("failed to parse json:", err.Error())
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	jsonFile.Close()

	json.Unmarshal(byteValue, &dto)

	for i := 0; i < len(dto); i++ {
		err = NewPhotosFromDTO(ctx, dto[i], storageRegistry)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewPhotosFromDTO(ctx context.Context, dto LocationDTOJson, storageRegistry *storage.Storages) error {

	for _, p := range dto.Photos {
		photo := placeDomain.NewPhoto(
			placeDomain.PhotoID(strconv.Itoa(p.ID)),
			userDomain.UserID(uuid.UUID{}),
			placeDomain.PlaceID(strconv.Itoa(dto.PlaceID)),
			strconv.Itoa(p.ID)+".jpg",
			p.Caption,
		)

		err := storageRegistry.PhotoStorage.CreatePhoto(ctx, photo)
		if err != nil {
			return err
		}
	}

	return nil
}
