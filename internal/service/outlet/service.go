package outlet

import (
	"context"

	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
)

type OutletService interface {
	CreateOutlet(context.Context, outletDomain.Outlet) error
	GetOutletsByStationID(context.Context, stationDomain.StationID) ([]outletDomain.Outlet, error)
	CreateOutletsList(context.Context, []outletDomain.Outlet) error
}

type OutletStorage interface {
	CreateOutlet(context.Context, outletDomain.Outlet) error
	GetOutletsByStationID(context.Context, stationDomain.StationID) ([]outletDomain.Outlet, error)
	DeleteOutletByID(context.Context, outletDomain.OutletID) error
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
	return s.outletStorage.CreateOutlet(ctx, outlet)
}

func (s *service) CreateOutletsList(ctx context.Context, outletsList []outletDomain.Outlet) error {
	for i := range outletsList {
		err := s.outletStorage.CreateOutlet(ctx, outletsList[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) GetOutletsByStationID(ctx context.Context, stationID stationDomain.StationID) ([]outletDomain.Outlet, error) {
	outlets, err := s.outletStorage.GetOutletsByStationID(ctx, stationID)
	if err != nil {
		return nil, err
	}

	return outlets, nil
}
