package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/poorfrombabylon/chargeMeBackend/internal/config"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration/my_ecars"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
	"log"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
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
	owner := placeDomain.MyECars

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

	myecarsStations, err := integrationRegistry.MyECarsIntegration.GetAllStations(ctx)
	if err != nil {
		return err
	}

	for _, s := range myecarsStations.Evse {
		coords := strings.Split(s.Location, ",")

		latitude, err := strconv.ParseFloat(coords[0], 32)
		if err != nil {
			return err
		}

		longitude, err := strconv.ParseFloat(coords[1], 32)
		if err != nil {
			return err
		}

		if latitude < 0 || longitude < 0 {
			continue
		}

		priceString := fmt.Sprintf("%v", s.Connectors[0].Cost)
		priceString = priceString[1 : len(priceString)-2]
		priceFields := strings.Split(priceString, " ")

		price, err := strconv.ParseFloat(priceFields[0], 64)
		if err != nil {
			return err
		}

		priceBool := false
		if price > 0 {
			priceBool = true
		}

		access := 1

		placeID := strings.Replace(s.Id, ";", ".", -1)

		place := placeDomain.NewPlaceWithID(
			placeDomain.PlaceID(placeID),
			s.Name,
			nil,
			float32(longitude),
			float32(latitude),
			&access,
			"G",
			s.Address,
			&s.Access,
			&priceBool,
			nil,
			nil,
			nil,
			nil,
			&s.Phone,
			domain.NewModel(),
		)

		place.SetCompanyName(&owner)

		err = storageRegistry.PlaceStorage.CreateOrUnhidePlace(ctx, place)
		if err != nil {
			fmt.Println(fmt.Sprintf("placeID %s error: %s", place.GetPlaceID().String(), err.Error()))
			return err
		}

		station := convertMyeCarsStation(s)

		err = storageRegistry.StationStorage.CreateOrUnhideStation(ctx, station)
		if err != nil {
			fmt.Println(fmt.Sprintf("stationID %s errot: %s", station.GetStationID().String(), err.Error()))
			return err
		}

		outlets := convertMyeCarsOutlets(s.Connectors, station.GetStationID())
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

func convertMyeCarsStation(station my_ecars.MyECarsStation) stationDomain.Station {
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

func convertMyeCarsOutlets(
	myeCarsConnectors []my_ecars.MyECarsConnector,
	stationID stationDomain.StationID,
) []outletDomain.Outlet {
	result := make([]outletDomain.Outlet, 0, len(myeCarsConnectors))

	for i, outlet := range myeCarsConnectors {
		connectorType := convertInterfaceToConnectorType(outlet.Type)
		kilowatts := convertInterfaceToKilowatts(outlet.Power)
		price, units := convertInterfaceToPriceInfo(outlet.Cost)

		o := outletDomain.NewOutletWithID(
			outletDomain.OutletID(stationID.String()+strconv.Itoa(i)),
			stationID,
			convertConnectorType(connectorType),
			&kilowatts,
			0,
			domain.NewModel(),
		)

		if price != 0 {
			o.SetPrice(&price)
			o.SetPriceUnit(&units)
		}

		result = append(result, o)

	}

	return result
}

func convertConnectorType(connectorType int) int {
	result := 0

	switch connectorType {
	case 1:
		result = 2
	case 2, 5:
		result = 7
	case -1:
		result = 1
	case 7, 8:
		result = 16
	case 9:
		result = 3
	case 10, 11:
		result = 13
	}

	return result
}

func convertInterfaceToConnectorType(info interface{}) int {
	fieldsString := fmt.Sprintf("%v", info)
	fieldsString = fieldsString[1 : len(fieldsString)-1]

	fields := strings.Split(fieldsString, " ")

	result, _ := strconv.Atoi(fields[0])

	return result
}

func convertInterfaceToKilowatts(info interface{}) float32 {
	fieldsString := fmt.Sprintf("%v", info)
	fieldsString = fieldsString[1 : len(fieldsString)-1]

	fields := strings.Split(fieldsString, " ")

	result, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		fmt.Println("err while parsing float for kilowatts", err.Error())
	}

	return float32(result)
}

func convertInterfaceToPriceInfo(info interface{}) (float32, string) {
	fieldsString := fmt.Sprintf("%v", info)
	fieldsString = fieldsString[1 : len(fieldsString)-1]

	fields := strings.Split(fieldsString, " ")

	result, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		fmt.Println("err while parsing float for price", err.Error())
	}

	return float32(result), fields[1]
}
