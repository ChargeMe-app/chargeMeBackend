package checkin

import (
	"context"
	"github.com/Masterminds/squirrel"
	checkinDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/checkin"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
	"time"
)

const (
	tableCheckins = "checkins"
)

type Storage interface {
	CreateCheckin(context.Context, checkinDomain.Checkin) error
	GetFinishedCheckins(context.Context) ([]checkinDomain.Checkin, error)
	DeleteCheckinByID(context.Context, checkinDomain.CheckinID) error
	DeleteCheckinByOutletID(context.Context, outletDomain.OutletID) error
	GetValidCheckinForStation(context.Context, stationDomain.StationID) ([]checkinDomain.Checkin, error)
}

func NewCheckinStorage(db libdb.DB) Storage {
	return &checkinStorage{db: db}
}

type checkinStorage struct {
	db libdb.DB
}

func (c *checkinStorage) GetFinishedCheckins(ctx context.Context) ([]checkinDomain.Checkin, error) {
	query := squirrel.Select(
		"id",
		"user_id",
		"station_id",
		"outlet_id",
		"user_name",
		"duration",
		"vehicle_type",
		"comment",
		"kilowatts",
		"rating",
		"is_auto",
		"started_at",
	).
		From(tableCheckins).
		Where(squirrel.Expr("finished_at < current_timestamp")).
		Where(squirrel.Eq{"deleted_at": nil}).
		PlaceholderFormat(squirrel.Dollar)

	var result []CheckinDTO

	err := c.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewCheckinListFromDTO(result), nil
}

func (c *checkinStorage) GetValidCheckinForStation(ctx context.Context, stationID stationDomain.StationID) ([]checkinDomain.Checkin, error) {
	query := squirrel.Select(
		"id",
		"user_id",
		"station_id",
		"outlet_id",
		"user_name",
		"duration",
		"vehicle_type",
		"comment",
		"kilowatts",
		"rating",
		"is_auto",
		"started_at",
		"finished_at",
	).
		From(tableCheckins).
		Where(squirrel.Expr("finished_at > current_timestamp")).
		Where(squirrel.Eq{"station_id": stationID.String()}).
		Where(squirrel.Eq{"deleted_at": nil}).
		PlaceholderFormat(squirrel.Dollar)

	var result []CheckinDTO

	err := c.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewCheckinListFromDTO(result), nil
}

func (c *checkinStorage) CreateCheckin(ctx context.Context, checkin checkinDomain.Checkin) error {
	query := squirrel.Insert(tableCheckins).
		Columns(
			"id",
			"user_id",
			"station_id",
			"outlet_id",
			"vehicle_type",
			"user_name",
			"duration",
			"comment",
			"kilowatts",
			"rating",
			"is_auto",
			"started_at",
			"finished_at",
		).
		Values(
			checkin.GetCheckinId().String(),
			checkin.GetUserId().String(),
			checkin.GetStationId().String(),
			checkin.GetOutletId().String(),
			checkin.GetVehicleType(),
			checkin.GetUserName(),
			checkin.GetDuration(),
			checkin.GetComment(),
			checkin.GetKilowatts(),
			checkin.GetRating(),
			checkin.IsAutoCheckin(),
			checkin.GetCreatedAt(),
			checkin.GetFinishedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := c.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (c *checkinStorage) DeleteCheckinByID(ctx context.Context, checkinID checkinDomain.CheckinID) error {
	query := squirrel.Update(tableCheckins).
		Set("deleted_at", time.Now().In(time.UTC)).
		Where(squirrel.Eq{"id": checkinID.String()}).
		Where(squirrel.Eq{"deleted_at": nil}).
		PlaceholderFormat(squirrel.Dollar)

	err := c.db.Update(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (c *checkinStorage) DeleteCheckinByOutletID(ctx context.Context, outletID outletDomain.OutletID) error {
	query := squirrel.Update(tableCheckins).
		Set("deleted_at", time.Now().In(time.UTC)).
		Where(squirrel.Eq{"outlet_id": outletID.String()}).
		Where(squirrel.Eq{"deleted_at": nil}).
		PlaceholderFormat(squirrel.Dollar)

	err := c.db.Update(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
