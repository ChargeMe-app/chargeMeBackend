package station

import (
	"context"
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
}

func NewStationStorage(db libdb.DB) Storage {
	return &stationStorage{db: db}
}

type stationStorage struct {
	db libdb.DB
}

func (s *stationStorage) CreateStation(ctx context.Context, station stationDomain.Station) error {
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
