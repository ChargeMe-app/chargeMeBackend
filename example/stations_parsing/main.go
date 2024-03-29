package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	amenityDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/amenity"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"

	"github.com/poorfrombabylon/chargeMeBackend/internal/config"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"

	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

type LocationDTOJson struct {
	PlaceID                      int                `json:"id,omitempty"`
	Name                         string             `json:"name"`
	Access                       int                `json:"access,omitempty"`
	Address                      string             `json:"address,omitempty"`
	Latitude                     float32            `json:"latitude"`
	Longitude                    float32            `json:"longitude"`
	Score                        *float32           `json:"score,omitempty"`
	IconType                     string             `json:"icon_type,omitempty"`
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
	LocationID int `json:"location_id,omitempty"`
	Form       int `json:"type,omitempty"`
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

	jsonFile, err := os.Open("/Users/almazkhayrullin/Desktop/jsons/full.json")
	if err != nil {
		log.Fatal("failed to parse json:", err.Error())
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	jsonFile.Close()

	json.Unmarshal(byteValue, &dto)

	for i := 0; i < len(dto); i++ {
		err = NewLocationFromDTO(ctx, dto[i], storageRegistry)
		if err != nil {
			return err
		}

		err = NewStationFromDTO(ctx, dto[i], storageRegistry)
		if err != nil {
			return err
		}

		err = NewReviewFromDTO(ctx, dto[i], storageRegistry)
		if err != nil {
			return err
		}

		err = NewAmenityFromDTO(ctx, dto[i], storageRegistry)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewLocationFromDTO(ctx context.Context, dto LocationDTOJson, storageRegistry *storage.Storages) error {
	place := placeDomain.NewPlaceWithID(
		placeDomain.PlaceID(strconv.Itoa(dto.PlaceID)),
		dto.Name,
		dto.Score,
		dto.Longitude,
		dto.Latitude,
		&dto.Access,
		dto.IconType,
		dto.Address,
		dto.Description,
		dto.Cost,
		dto.CostDescription,
		dto.Hours,
		dto.Open247,
		dto.ComingSoon,
		dto.PhoneNumber,
		domain.NewModel(),
	)

	err := storageRegistry.PlaceStorage.CreatePlace(ctx, place)
	if err != nil {
		return err
	}

	return nil
}

func NewStationFromDTO(ctx context.Context, dto LocationDTOJson, storageRegistry *storage.Storages) error {
	var err error

	for _, stationDTO := range dto.Stations {
		station := stationDomain.NewStationWithID(
			stationDomain.StationID(strconv.Itoa(stationDTO.Id)),
			placeDomain.PlaceID(strconv.Itoa(dto.PlaceID)),
			stationDTO.Available,
			stationDTO.Cost,
			stationDTO.Name,
			stationDTO.Manufacturer,
			stationDTO.CostDescription,
			stationDTO.Hours,
			stationDTO.Kilowatts,
			domain.NewModel(),
		)

		err = storageRegistry.StationStorage.CreateStation(ctx, station)
		if err != nil {
			return err
		}

		err = NewOutletsFromDTO(ctx, stationDTO, storageRegistry)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewOutletsFromDTO(ctx context.Context, stationDTO StationsDTOJson, storageRegistry *storage.Storages) error {
	var err error

	for _, outletDTO := range stationDTO.Outlets {
		outlet := outletDomain.NewOutletWithID(
			outletDomain.OutletID(strconv.Itoa(outletDTO.Id)),
			stationDomain.StationID(strconv.Itoa(stationDTO.Id)),
			outletDTO.ConnectorType,
			outletDTO.Kilowatts,
			outletDTO.Power,
			domain.NewModel(),
		)

		err = storageRegistry.OutletStorage.CreateOutlet(ctx, outlet)
		if err != nil {
			fmt.Println("outlet:", outlet)

			return err
		}

	}

	return nil
}

func NewReviewFromDTO(ctx context.Context, dto LocationDTOJson, storageRegistry *storage.Storages) error {
	var err error

	for _, reviewDTO := range dto.Reviews {

		createdAt, _ := time.Parse("2006-01-02T15:04:05Z", reviewDTO.CreatedAT)

		userID := userDomain.UserID(uuid.UUID{})

		review := reviewDomain.NewReviewWithID(
			reviewDomain.ReviewID(uuid.New().String()),
			stationDomain.StationID(strconv.Itoa(reviewDTO.StationID)),
			outletDomain.OutletID(strconv.Itoa(reviewDTO.OutletID)),
			&userID,
			reviewDTO.Comment,
			reviewDTO.Rating,
			reviewDTO.ConnectorType,
			reviewDTO.Kilowatts,
			reviewDTO.User.FirstName,
			reviewDTO.VehicleName,
			reviewDTO.VehicleType,
			domain.NewModelFrom(createdAt, nil),
		)

		err = storageRegistry.ReviewStorage.CreateReview(ctx, review)
		if err != nil {
			fmt.Println("review:", review)

			return err
		}
	}

	return nil
}

func NewAmenityFromDTO(ctx context.Context, dto LocationDTOJson, storageRegistry *storage.Storages) error {
	var err error

	for _, amenityDTO := range dto.Amenities {
		amenity := amenityDomain.NewAmenityWithID(
			amenityDomain.AmenityID(uuid.New().String()),
			placeDomain.PlaceID(strconv.Itoa(amenityDTO.LocationID)),
			amenityDTO.Form,
			domain.NewModel(),
		)

		err = storageRegistry.AmenityStorage.CreateAmenity(ctx, amenity)
		if err != nil {
			fmt.Println("amenity:", amenity)

			return err
		}
	}

	return nil
}
