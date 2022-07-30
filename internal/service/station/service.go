package station

import (
	"context"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"

	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
)

type StationService interface {
	CreateStation(ctx context.Context, station stationDomain.Station) error
	GetStationsByPlaceID(context.Context, placeDomain.PlaceID) ([]stationDomain.Station, error)
}

type StationStorage interface {
	CreateStation(ctx context.Context, station stationDomain.Station) error
	GetStationsByPlaceID(context.Context, placeDomain.PlaceID) ([]stationDomain.Station, error)
}

type service struct {
	stationStorage StationStorage
}

func NewStationService(stationStorage StationStorage) StationService {
	return &service{
		stationStorage: stationStorage,
	}
}

func (s *service) CreateStation(ctx context.Context, station stationDomain.Station) error {
	return s.stationStorage.CreateStation(ctx, station)
}

func (s *service) GetStationsByPlaceID(ctx context.Context, placeID placeDomain.PlaceID) ([]stationDomain.Station, error) {
	return s.stationStorage.GetStationsByPlaceID(ctx, placeID)
}
