package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/poorfrombabylon/chargeMeBackend/internal/config"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

func radiansToDegrees(radians float64) float64 {
	degrees := radians * 180 / math.Pi

	return degrees
}

func degreesToRadians(degrees float64) float64 {
	radians := degrees * math.Pi / 180

	return radians
}

// distance in kilometers
func getDistanceBetweenPoints(latitudeBase, longitudeBase, latitudeNew, longitudeNew float64) float64 {
	theta := longitudeBase - longitudeNew

	distance := 60 * 1.1515 * radiansToDegrees(
		math.Acos(
			(math.Sin(degreesToRadians(latitudeBase))*math.Sin(degreesToRadians(latitudeNew)))+
				(math.Cos(degreesToRadians(latitudeBase))*math.Cos(degreesToRadians(latitudeNew))*math.Cos(degreesToRadians(theta))),
		),
	)

	return distance * 1.609344
}

func deleteSitronicsStations(
	ctx context.Context,
	storageRegistry *storage.Storages,
	integrationRegistry integration.Integration,
) error {
	allPlaces, err := storageRegistry.PlaceStorage.GetAllPlaces(ctx)
	if err != nil {
		log.Fatal("failed to get all places sitronics", err.Error())
	}

	sitronicsResp, err := integrationRegistry.SitronicsIntegration.GetAllStations(ctx)
	if err != nil {
		log.Fatal("failed to get sitronics places ", err.Error())
	}

	sitronicsPlaces := sitronicsResp.CPList

	for i := range allPlaces {
		for j := range sitronicsPlaces {

			latitudeBase := float64(allPlaces[i].GetPlaceLatitude())
			longitudeBase := float64(allPlaces[i].GetPlaceLongitude())

			latitudeNew := float64(sitronicsPlaces[j].Latitude)
			longitudeNew := float64(sitronicsPlaces[j].Longitude)

			if getDistanceBetweenPoints(latitudeBase, longitudeBase, latitudeNew, longitudeNew) <= 0.04 {
				err := storageRegistry.AmenityStorage.DeleteAmenitiesByLocationID(ctx, allPlaces[i].GetPlaceID())
				if err != nil {
					fmt.Println("placeID", allPlaces[i].GetPlaceID().String())
					log.Fatal("failed to delete amenity sitronics ", err.Error())
				}

				stations, err := storageRegistry.StationStorage.GetStationsByPlaceID(ctx, allPlaces[i].GetPlaceID())
				if err != nil {
					fmt.Println("placeID", allPlaces[i].GetPlaceID().String())
					log.Fatal("failed to get stations sitronics ", err.Error())
				}

				for l := range stations {
					err = storageRegistry.ReviewStorage.DeleteReviewsByStationID(ctx, stations[l].GetStationID())
					if err != nil {
						fmt.Println("stationID", stations[l].GetStationID().String())
						log.Fatal("failed to delete review sitronics ", err.Error())
					}
				}

				err = storageRegistry.PlaceStorage.DeletePlaceByID(ctx, allPlaces[i].GetPlaceID())
				if err != nil {
					fmt.Println("placeID", allPlaces[i].GetPlaceID().String())
					log.Fatal("failed to delete place sitronics ", err.Error())
				}
			}

		}
	}

	return nil
}

func deleteMyECarsStations(
	ctx context.Context,
	storageRegistry *storage.Storages,
	integrationRegistry integration.Integration,
) error {
	allPlaces, err := storageRegistry.PlaceStorage.GetAllPlaces(ctx)
	if err != nil {
		log.Fatal("failed to get all places my.eCars", err.Error())
	}

	myECarsResp, err := integrationRegistry.MyECarsIntegration.GetAllStations(ctx)
	if err != nil {
		log.Fatal("failed to get my.eCars places ", err.Error())
	}

	myECarsPlaces := myECarsResp.Evse

	for i := range allPlaces {
		for j := range myECarsPlaces {

			coordsNew := strings.Split(myECarsPlaces[j].Location, ",")

			latitudeBase := float64(allPlaces[i].GetPlaceLatitude())
			longitudeBase := float64(allPlaces[i].GetPlaceLongitude())

			latitudeNew, err := strconv.ParseFloat(coordsNew[0], 64)
			if err != nil {
				log.Fatal("failed to get my.eCars latitude ", err.Error())
			}

			longitudeNew, err := strconv.ParseFloat(coordsNew[1], 64)
			if err != nil {
				log.Fatal("failed to get my.eCars longitude ", err.Error())
			}

			if getDistanceBetweenPoints(latitudeBase, longitudeBase, latitudeNew, longitudeNew) <= 0.04 {
				err := storageRegistry.AmenityStorage.DeleteAmenitiesByLocationID(ctx, allPlaces[i].GetPlaceID())
				if err != nil {
					fmt.Println("placeID", allPlaces[i].GetPlaceID().String())
					log.Fatal("failed to delete amenity my.eCars ", err.Error())
				}

				stations, err := storageRegistry.StationStorage.GetStationsByPlaceID(ctx, allPlaces[i].GetPlaceID())
				if err != nil {
					fmt.Println("placeID", allPlaces[i].GetPlaceID().String())
					log.Fatal("failed to get stations my.eCars ", err.Error())
				}

				for l := range stations {
					err = storageRegistry.ReviewStorage.DeleteReviewsByStationID(ctx, stations[l].GetStationID())
					if err != nil {
						fmt.Println("stationID", stations[l].GetStationID().String())
						log.Fatal("failed to delete review my.eCars ", err.Error())
					}
				}

				err = storageRegistry.PlaceStorage.DeletePlaceByID(ctx, allPlaces[i].GetPlaceID())
				if err != nil {
					fmt.Println("placeID", allPlaces[i].GetPlaceID().String())
					log.Fatal("failed to delete place my.eCars ", err.Error())
				}
			}

		}
	}

	return nil
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

	integrationRegistry := integration.NewIntegrationRegistry(cfg)

	err = deleteSitronicsStations(ctx, storageRegistry, integrationRegistry)
	if err != nil {
		fmt.Println("deleteSitronicsStations error: ", err.Error())
	}

	err = deleteMyECarsStations(ctx, storageRegistry, integrationRegistry)
	if err != nil {
		fmt.Println("deleteMyECarsStations error: ", err.Error())
	}
}
