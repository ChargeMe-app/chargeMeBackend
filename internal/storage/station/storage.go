package station

import (
	"context"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"

	"github.com/Masterminds/squirrel"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	tableStations = "stations"
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
	query := squirrel.Insert(tableStations).
		Columns(
			"id",
			"location_id",
		).
		Values(
			station.GetStationID().String(),
			station.GetPlaceID().String(),
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
		From(tableStations).
		Where(squirrel.Eq{"location_id": placeID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var result []StationDTO

	err := s.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewStationListFromDTO(result), nil
}
