package outlet

import (
	"context"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"log"

	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
)

type OutletService interface {
	CreateOutlet(context.Context, outletDomain.Outlet) error
	GetOutletsByStationID(context.Context, stationDomain.StationID) ([]outletDomain.Outlet, error)
}

type OutletStorage interface {
	CreateOutlet(context.Context, outletDomain.Outlet) error
	GetOutletsByStationID(context.Context, stationDomain.StationID) ([]outletDomain.Outlet, error)
}

type service struct {
	outletStorage OutletStorage
}

func NewOutletService(outletStorage OutletStorage) OutletService {
	return &service{
		outletStorage: outletStorage,
	}
}

func (s *service) CreateOutlet(ctx context.Context, outlet outletDomain.Outlet) error {
	log.Println("outlet.service.CreateOutlet")

	return s.outletStorage.CreateOutlet(ctx, outlet)
}

func (s *service) GetOutletsByStationID(ctx context.Context, stationID stationDomain.StationID) ([]outletDomain.Outlet, error) {
	log.Println("outlet.service.GetOutletsByStationID")

	outlets, err := s.outletStorage.GetOutletsByStationID(ctx, stationID)
	if err != nil {
		return nil, err
	}

	return outlets, nil
}
