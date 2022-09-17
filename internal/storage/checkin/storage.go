package checkin

import (
	"context"
	"github.com/Masterminds/squirrel"
	checkinDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/checkin"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
	"time"
)

const (
	tableCheckins = "checkins"
)

type Storage interface {
	CreateCheckin(context.Context, checkinDomain.Checkin) error
}

func NewCheckinStorage(db libdb.DB) Storage {
	return &checkinStorage{db: db}
}

type checkinStorage struct {
	db libdb.DB
}

func (c *checkinStorage) CreateCheckin(ctx context.Context, checkin checkinDomain.Checkin) error {
	finishedAt := checkin.GetCreatedAt().Add(time.Duration(checkin.GetDuration()) * time.Minute)

	query := squirrel.Insert(tableCheckins).
		Columns(
			"id",
			"user_id",
			"station_id",
			"outlet_id",
			"vehicle_type",
			"comment",
			"kilowatts",
			"rating",
			"started_at",
			"finished_at",
		).
		Values(
			checkin.GetCheckinId().String(),
			checkin.GetUserId().String(),
			checkin.GetStationId().String(),
			checkin.GetOutletId().String(),
			checkin.GetVehicleType(),
			checkin.GetComment(),
			checkin.GetKilowatts(),
			checkin.GetRating(),
			checkin.GetCreatedAt(),
			finishedAt,
		).
		PlaceholderFormat(squirrel.Dollar)

	err := c.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
