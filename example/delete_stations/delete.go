package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"
	"os/signal"
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

	allPlaces, err := storageRegistry.PlaceStorage.GetAllPlaces(ctx)
	if err != nil {
		log.Fatal("failed to get all places ", err.Error())
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

			if getDistanceBetweenPoints(latitudeBase, longitudeBase, latitudeNew, longitudeNew) <= 0.02 {
				err := storageRegistry.AmenityStorage.DeleteAmenitiesByLocationID(ctx, allPlaces[i].GetPlaceID())
				if err != nil {
					fmt.Println("placeID", allPlaces[i].GetPlaceID().String())
					log.Fatal("failed to delete amenity ", err.Error())
				}

				stations, err := storageRegistry.StationStorage.GetStationsByPlaceID(ctx, allPlaces[i].GetPlaceID())
				if err != nil {
					fmt.Println("placeID", allPlaces[i].GetPlaceID().String())
					log.Fatal("failed to get stations ", err.Error())
				}

				for l := range stations {
					err = storageRegistry.ReviewStorage.DeleteReviewsByStationID(ctx, stations[l].GetStationID())
					if err != nil {
						fmt.Println("stationID", stations[l].GetStationID().String())
						log.Fatal("failed to delete review ", err.Error())
					}
				}

				err = storageRegistry.PlaceStorage.DeletePlaceByID(ctx, allPlaces[i].GetPlaceID())
				if err != nil {
					fmt.Println("placeID", allPlaces[i].GetPlaceID().String())
					log.Fatal("failed to delete place ", err.Error())
				}
			}

		}
	}
}
