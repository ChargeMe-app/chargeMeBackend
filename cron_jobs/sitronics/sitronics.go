package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/poorfrombabylon/chargeMeBackend/internal/config"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration/sitronics"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

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

	integrationRegistry := integration.NewIntegrationRegistry(cfg)

	err = startJob(ctx, storageRegistry, integrationRegistry)
	if err != nil {
		log.Fatal("failed while doing job: ", err.Error())
	}
}

func startJob(
	ctx context.Context,
	storageRegistry *storage.Storages,
	integrationRegistry integration.Integration,
) error {
	owner := placeDomain.Sitronics
	var priceUnit string

	err := storageRegistry.PlaceStorage.HideCompanyPlaces(ctx, owner)
	if err != nil {
		return err
	}

	err = storageRegistry.StationStorage.HideCompanyStations(ctx, owner)
	if err != nil {
		return err
	}

	err = storageRegistry.OutletStorage.HideCompanyOutlets(ctx, owner)
	if err != nil {
		return err
	}

	sitronicsStations, err := integrationRegistry.SitronicsIntegration.GetAllStations(ctx)
	if err != nil {
		return err
	}

	for _, s := range sitronicsStations.CPList {
		var costDescription *string
		cost := false

		if s.ServiceList[0].Price != 0 {
			cost = true

			priceUnit = fmt.Sprintf("%s/%s", s.ServiceList[0].CurrencyName, s.ServiceList[0].Unit)

			description := fmt.Sprintf("%v %s", s.ServiceList[0].Price/100, priceUnit)
			costDescription = &description
		}

		iconType := "G"
		yIconType := map[int]bool{3: true, 4: true, 13: true}
		for _, outlet := range s.Connectors {
			connectorType := convertConnectorType(outlet.Type)
			if yIconType[connectorType] {
				iconType = "Y"
				break
			}
		}

		if s.Status == 1 {
			iconType += "R"
		}

		access := 1

		placeID := strings.Replace(s.Id, ";", ".", -1)

		place := placeDomain.NewPlaceWithID(
			placeDomain.PlaceID(placeID),
			s.Name,
			nil,
			s.Longitude,
			s.Latitude,
			&access,
			iconType,
			s.Address,
			&s.PublicDescription,
			&cost,
			costDescription,
			&s.WorkingTime,
			nil,
			nil,
			s.PhoneNumber,
			domain.NewModel(),
		)

		place.SetCompanyName(&owner)

		err = storageRegistry.PlaceStorage.CreateOrUnhidePlace(ctx, place)
		if err != nil {
			fmt.Println(fmt.Sprintf("placeID %s error: %s", place.GetPlaceID().String(), err.Error()))
			return err
		}

		station := convertSitronicsStation(s)

		err = storageRegistry.StationStorage.CreateOrUnhideStation(ctx, station)
		if err != nil {
			fmt.Println(fmt.Sprintf("stationID %s errot: %s", station.GetStationID().String(), err.Error()))
			return err
		}

		outlets, err := convertSitronicsOutlets(s.Connectors, station.GetStationID(), s.ServiceList[0].Price, priceUnit)
		if err != nil {
			return err
		}

		err = storageRegistry.OutletStorage.CreateOrUnhideOutletsList(ctx, outlets)
		if err != nil {
			fmt.Println("outletId error CreateOrUnhideOutletsList: ", err.Error())
			return err
		}
	}

	return nil
}

func convertSitronicsStation(station sitronics.SitronicsStation) stationDomain.Station {
	stationID := strings.Replace(station.Id, ";", ".", -1)

	return stationDomain.NewStationWithID(
		stationDomain.StationID(stationID),
		placeDomain.PlaceID(stationID),
		nil,
		nil,
		&station.Name,
		nil,
		nil,
		nil,
		nil,
		domain.NewModel(),
	)
}

func convertSitronicsOutlets(
	sitronicsOutlets []sitronics.SitronicsConnector,
	stationID stationDomain.StationID,
	price int,
	priceUnit string,
) ([]outletDomain.Outlet, error) {
	result := make([]outletDomain.Outlet, 0, len(sitronicsOutlets))

	for i, outlet := range sitronicsOutlets {
		maxP, err := strconv.Atoi(outlet.MaxPower)
		if err != nil {
			return nil, err
		}

		o := outletDomain.NewOutletWithID(
			outletDomain.OutletID(stationID.String()+strconv.Itoa(i)),
			stationID,
			convertConnectorType(outlet.Type),
			nil,
			maxP,
			domain.NewModel(),
		)

		if price != 0 {
			outletPrice := float32(price / 100)

			o.SetPrice(&outletPrice)
			o.SetPriceUnit(&priceUnit)
		}

		result = append(result, o)
	}

	return result, nil
}

func convertConnectorType(connectorType string) int {
	result := 0

	switch connectorType {
	case "3":
		result = 13
	case "4":
		result = 3
	case "5":
		result = 16
	case "2":
		result = 7
	case "1":
		result = 1
	}

	return result
}
