package outlet

import (
	"context"

	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
)

type OutletService interface {
	CreateOutlet(context.Context, outletDomain.Outlet) error
}

type OutletStorage interface {
	CreateOutlet(context.Context, outletDomain.Outlet) error
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
