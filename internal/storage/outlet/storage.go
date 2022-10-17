package outlet

import (
	"context"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"log"

	"github.com/Masterminds/squirrel"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	tableOutlets = "outlets"
)

type Storage interface {
	CreateOutlet(context.Context, outletDomain.Outlet) error
	GetOutletsByStationID(context.Context, stationDomain.StationID) ([]outletDomain.Outlet, error)
	GetOutletByID(context.Context, outletDomain.OutletID) (outletDomain.Outlet, error)
}

func NewOutletStorage(db libdb.DB) Storage {
	return &outletStorage{db: db}
}

type outletStorage struct {
	db libdb.DB
}

func (o *outletStorage) CreateOutlet(ctx context.Context, outlet outletDomain.Outlet) error {
	log.Println("outlet.storage.CreateOutlet")

	query := squirrel.Insert(tableOutlets).
		Columns(
			"id",
			"station_id",
			"connector",
			"kilowatts",
			"power",
			"created_at",
		).
		Values(
			outlet.GetOutletID().String(),
			outlet.GetStationID().String(),
			outlet.GetConnector(),
			outlet.GetKilowatts(),
			outlet.GetPower(),
			outlet.GetCreatedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := o.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (o *outletStorage) GetOutletsByStationID(ctx context.Context, stationID stationDomain.StationID) ([]outletDomain.Outlet, error) {
	log.Println("outlet.storage.GetOutletsByStationID")

	query := squirrel.Select(
		"id",
		"station_id",
		"connector",
		"kilowatts",
		"power",
	).
		From(tableOutlets).
		Where(squirrel.Eq{"station_id": stationID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var result []OutletDTO

	err := o.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	outlets := NewOutletListFromDTO(result)

	return outlets, nil
}

func (o *outletStorage) GetOutletByID(ctx context.Context, outletId outletDomain.OutletID) (outletDomain.Outlet, error) {
	log.Println("outlet.storage.GetOutletByID")

	query := squirrel.Select(
		"id",
		"station_id",
		"connector",
		"kilowatts",
		"power",
	).
		From(tableOutlets).
		Where(squirrel.Eq{"id": outletId.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var result OutletDTO

	err := o.db.Get(ctx, query, &result)
	if err != nil {
		return outletDomain.Outlet{}, err
	}

	outlet := NewOutletFromDTO(result)

	return outlet, nil
}
