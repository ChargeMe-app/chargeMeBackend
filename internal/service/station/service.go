package station

import (
	"context"

	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
)

type StationService interface {
	CreateStation(ctx context.Context, station stationDomain.Station) error
	GetStationsByPlaceID(context.Context, placeDomain.PlaceID) ([]stationDomain.Station, error)
	DeleteStationsByPlaceID(context.Context, placeDomain.PlaceID) error
}

type StationStorage interface {
	CreateStation(ctx context.Context, station stationDomain.Station) error
	GetStationsByPlaceID(context.Context, placeDomain.PlaceID) ([]stationDomain.Station, error)
	DeleteStationByID(context.Context, stationDomain.StationID) error
}

type OutletStorage interface {
	DeleteOutletsByStationID(context.Context, stationDomain.StationID) error
}

type service struct {
	stationStorage StationStorage
	outletStorage  OutletStorage
}

func NewStationService(stationStorage StationStorage, outletStorage OutletStorage) StationService {
	return &service{
		stationStorage: stationStorage,
		outletStorage:  outletStorage,
	}
}

func (s *service) CreateStation(ctx context.Context, station stationDomain.Station) error {
	return s.stationStorage.CreateStation(ctx, station)
}

func (s *service) GetStationsByPlaceID(ctx context.Context, placeID placeDomain.PlaceID) ([]stationDomain.Station, error) {
	return s.stationStorage.GetStationsByPlaceID(ctx, placeID)
}

func (s *service) DeleteStationByID(ctx context.Context, stationID stationDomain.StationID) error {
	return s.stationStorage.DeleteStationByID(ctx, stationID)
}

func (s *service) DeleteStationsByPlaceID(
	ctx context.Context,
	placeID placeDomain.PlaceID,
) error {
	stationsList, err := s.stationStorage.GetStationsByPlaceID(ctx, placeID)
	if err != nil {
		return err
	}

	for i := range stationsList {
		err = s.outletStorage.DeleteOutletsByStationID(ctx, stationsList[i].GetStationID())
		if err != nil {
			return err
		}

		err = s.stationStorage.DeleteStationByID(ctx, stationsList[i].GetStationID())
		if err != nil {
			return err
		}
	}

	return nil
}
