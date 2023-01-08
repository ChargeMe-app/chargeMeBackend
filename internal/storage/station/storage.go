package station

import (
	"context"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/place"
	"log"

	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"

	"github.com/Masterminds/squirrel"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	TableStations = "stations"
)

type Storage interface {
	CreateStation(context.Context, stationDomain.Station) error
	GetStationsByPlaceID(context.Context, placeDomain.PlaceID) ([]stationDomain.Station, error)
	GetPlaceIdByStationID(context.Context, stationDomain.StationID) (placeDomain.PlaceID, error)
	DeleteStationByID(context.Context, stationDomain.StationID) error
	HideCompanyStations(context.Context, string) error
	CreateOrUnhideStation(context.Context, stationDomain.Station) error
}

func NewStationStorage(db libdb.DB) Storage {
	return &stationStorage{db: db}
}

type stationStorage struct {
	db libdb.DB
}

func (s *stationStorage) CreateStation(ctx context.Context, station stationDomain.Station) error {
	log.Println("station.storage.CreateStation")
	query := squirrel.Insert(TableStations).
		Columns(
			"id",
			"location_id",
			"available",
			"cost",
			"name",
			"manufacturer",
			"cost_description",
			"hours",
			"kilowatts",
			"created_at",
		).
		Values(
			station.GetStationID().String(),
			station.GetPlaceID().String(),
			station.GetStationAvailability(),
			station.GetStationCost(),
			station.GetStationName(),
			station.GetStationManufacturer(),
			station.GetStationCostDescription(),
			station.GetStationWorkingHours(),
			station.GetStationKilowatts(),
			station.GetCreatedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (s *stationStorage) GetStationsByPlaceID(
	ctx context.Context,
	placeID placeDomain.PlaceID,
) ([]stationDomain.Station, error) {
	query := squirrel.Select(
		"id",
		"location_id",
	).
		From(TableStations).
		Where(squirrel.Eq{"location_id": placeID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var result []StationDTO

	err := s.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewStationListFromDTO(result), nil
}

func (s *stationStorage) GetPlaceIdByStationID(ctx context.Context, stationID stationDomain.StationID) (placeDomain.PlaceID, error) {
	query := squirrel.Select(
		"location_id",
	).
		From(TableStations).
		Where(squirrel.Eq{"id": stationID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var result string

	err := s.db.Get(ctx, query, &result)
	if err != nil {
		return "", err
	}

	return placeDomain.PlaceID(result), nil
}

func (s *stationStorage) DeleteStationByID(ctx context.Context, stationID stationDomain.StationID) error {
	query := squirrel.Delete(TableStations).
		Where(squirrel.Eq{"id": stationID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Delete(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (s *stationStorage) HideCompanyStations(ctx context.Context, companyName string) error {
	query := squirrel.Update(TableStations+" s").
		Set("hide", true).
		Where(squirrel.Expr(
			"s.location_id IN (SELECT p.id FROM "+place.TablePlaces+" p "+
				"WHERE p.company_name = ?)", companyName)).
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Update(ctx, query)
	if err != nil {
		return err
	}

	return err
}

func (s *stationStorage) CreateOrUnhideStation(ctx context.Context, station stationDomain.Station) error {
	query := squirrel.Insert(TableStations).
		Columns(
			"id",
			"location_id",
			"available",
			"cost",
			"name",
			"manufacturer",
			"cost_description",
			"hours",
			"kilowatts",
			"created_at",
		).
		Values(
			station.GetStationID().String(),
			station.GetPlaceID().String(),
			station.GetStationAvailability(),
			station.GetStationCost(),
			station.GetStationName(),
			station.GetStationManufacturer(),
			station.GetStationCostDescription(),
			station.GetStationWorkingHours(),
			station.GetStationKilowatts(),
			station.GetCreatedAt(),
		).
		Suffix("ON CONFLICT (id) DO UPDATE SET hide = false").
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
