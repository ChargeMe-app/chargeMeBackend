package outlet

import (
	"context"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/place"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/station"

	"github.com/Masterminds/squirrel"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	tableOutlets = "outlets"
)

type Storage interface {
	CreateOutlet(context.Context, outletDomain.Outlet) error
	GetOutletsByStationID(context.Context, stationDomain.StationID) ([]outletDomain.Outlet, error)
	GetOutletByID(context.Context, outletDomain.OutletID) (outletDomain.Outlet, error)
	DeleteOutletByID(context.Context, outletDomain.OutletID) error
	DeleteOutletsByStationID(context.Context, stationDomain.StationID) error
	HideCompanyOutlets(context.Context, string) error
	CreateOrUnhideOutletsList(context.Context, []outletDomain.Outlet) error
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
			"price",
			"price_unit",
			"created_at",
		).
		Values(
			outlet.GetOutletID().String(),
			outlet.GetStationID().String(),
			outlet.GetConnector(),
			outlet.GetKilowatts(),
			outlet.GetPower(),
			outlet.GetPrice(),
			outlet.GetPriceUnit(),
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
	query := squirrel.Select(
		"id",
		"station_id",
		"connector",
		"kilowatts",
		"power",
		"price",
		"price_unit",
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
	query := squirrel.Select(
		"id",
		"station_id",
		"connector",
		"kilowatts",
		"power",
		"price",
		"price_unit",
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

func (o *outletStorage) DeleteOutletsByStationID(ctx context.Context, stationID stationDomain.StationID) error {
	query := squirrel.Delete(tableOutlets).
		Where(squirrel.Eq{"station_id": stationID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	err := o.db.Delete(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (o *outletStorage) DeleteOutletByID(ctx context.Context, outletID outletDomain.OutletID) error {
	query := squirrel.Delete(tableOutlets).
		Where(squirrel.Eq{"id": outletID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	err := o.db.Delete(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (o *outletStorage) HideCompanyOutlets(ctx context.Context, companyName string) error {
	query := squirrel.Update(tableOutlets+" o").
		Set("hide", true).
		Where(squirrel.Expr(
			"o.station_id IN (SELECT s.id FROM "+station.TableStations+" s "+
				"JOIN "+place.TablePlaces+" p ON s.location_id = p.id "+
				"WHERE p.company_name = ?)", companyName)).
		PlaceholderFormat(squirrel.Dollar)

	err := o.db.Update(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (o *outletStorage) CreateOrUnhideOutletsList(ctx context.Context, outlets []outletDomain.Outlet) error {
	query := squirrel.Insert(tableOutlets).
		Columns(
			"id",
			"station_id",
			"connector",
			"kilowatts",
			"power",
			"price",
			"price_unit",
			"created_at",
		).
		Suffix("ON CONFLICT (id) DO UPDATE SET hide = false").
		PlaceholderFormat(squirrel.Dollar)

	for i := range outlets {
		query = query.Values(
			outlets[i].GetOutletID().String(),
			outlets[i].GetStationID().String(),
			outlets[i].GetConnector(),
			outlets[i].GetKilowatts(),
			outlets[i].GetPower(),
			outlets[i].GetPrice(),
			outlets[i].GetPriceUnit(),
			outlets[i].GetCreatedAt(),
		)
	}

	err := o.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
