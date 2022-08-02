package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"

	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"

	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "postgres"
)

type LocationDTOJson struct {
	PlaceID   int               `json:"id,omitempty"`
	Name      string            `json:"name"`
	Access    int               `json:"access,omitempty"`
	Address   string            `json:"address,omitempty"`
	Latitude  float32           `json:"latitude"`
	Longitude float32           `json:"longitude"`
	Score     *float32          `json:"score,omitempty"`
	IconLink  *string           `json:"icon"`
	Stations  []StationsDTOJson `json:"stations"`
}

type StationsDTOJson struct {
	Id         int             `json:"id"`
	LocationID int             `json:"location_id"`
	Outlets    []OutletDTOJson `json:"outlets"`
}

type OutletDTOJson struct {
	Id            int      `json:"id"`
	ConnectorType int      `json:"connector_type"`
	Kilowatts     *float32 `json:"kilowatts,omitempty"`
	Power         int      `json:"power"`
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

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect to database:", err.Error())
	}

	dbx := sqlx.NewDb(db, "pgx")
	libDBWrapper := libdb.NewSQLXDB(dbx)

	storageRegistry := storage.NewStorageRegistry(libDBWrapper)

	err = startJob(ctx, storageRegistry)
	if err != nil {
		log.Fatal("failure:")
	}
}

func startJob(ctx context.Context, storageRegistry *storage.Storages) error {
	var dto []LocationDTOJson

	jsonFile, err := os.Open("/Users/almazkhayrullin/Desktop/stationData.json")
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
		dto.IconLink,
		&dto.Address,
	)

	err := storageRegistry.PlaceStorage.CreatePlace(ctx, place)
	if err != nil {
		return err
	}

	return nil
}

func NewStationFromDTO(ctx context.Context, dto LocationDTOJson, storageRegistry *storage.Storages) error {
	var err error

	for _, i := range dto.Stations {
		station := stationDomain.NewStationWithID(
			stationDomain.StationID(strconv.Itoa(i.Id)),
			placeDomain.PlaceID(strconv.Itoa(dto.PlaceID)),
		)

		err = storageRegistry.StationStorage.CreateStation(ctx, station)
		if err != nil {
			return err
		}

		err = NewOutletsFromDTO(ctx, i, storageRegistry)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewOutletsFromDTO(ctx context.Context, stationDTO StationsDTOJson, storageRegistry *storage.Storages) error {
	var err error

	for _, i := range stationDTO.Outlets {
		outlet := outletDomain.NewOutletWithID(
			outletDomain.OutletID(strconv.Itoa(i.Id)),
			stationDomain.StationID(strconv.Itoa(stationDTO.Id)),
			i.ConnectorType,
			i.Kilowatts,
			i.Power,
		)

		err = storageRegistry.OutletStorage.CreateOutlet(ctx, outlet)
		if err != nil {
			return err
		}

	}

	return nil
}
