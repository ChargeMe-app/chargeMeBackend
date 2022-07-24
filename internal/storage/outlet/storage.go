package outlet

import (
	"context"

	"github.com/Masterminds/squirrel"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	tableOutlets = "outlets"
)

type Storage interface {
	CreateOutlet(context.Context, outletDomain.Outlet) error
}

func NewOutletStorage(db libdb.DB) Storage {
	return &outletStorage{db: db}
}

type outletStorage struct {
	db libdb.DB
}

func (o *outletStorage) CreateOutlet(ctx context.Context, outlet outletDomain.Outlet) error {
	query := squirrel.Insert(tableOutlets).
		Columns(
			"id",
			"station_id",
			"connector",
			"kilowatts",
			"power",
		).
		Values(
			outlet.GetOutletID().String(),
			outlet.GetStationID().String(),
			outlet.GetConnector(),
			outlet.GetKilowatts(),
			outlet.GetPower(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := o.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
