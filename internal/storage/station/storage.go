package station

import (
	"context"

	"github.com/Masterminds/squirrel"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	tableStations = "stations"
)

type Storage interface {
	CreateStation(context.Context, stationDomain.Station) error
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
