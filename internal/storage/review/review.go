package review

import (
	"context"
	"github.com/Masterminds/squirrel"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	tableReviews = "reviews"
)

type Storage interface {
	CreateReview(context.Context, reviewDomain.Review) error
}

func NewReviewStorage(db libdb.DB) Storage {
	return &reviewStorage{db: db}
}

type reviewStorage struct {
	db libdb.DB
}

func (r *reviewStorage) CreateReview(ctx context.Context, review reviewDomain.Review) error {
	query := squirrel.Insert(tableReviews).
		Columns(
			"id",
			"comment",
			"station_id",
			"outlet_id",
			"rating",
			"vehicle_name",
			"vehicle_type",
			"created_at",
		).
		Values(
			review.GetReviewID().String(),
			review.GetComment(),
			review.GetStationID().String(),
			review.GetOutletID().String(),
			review.GetRating(),
			review.GetVehicleName(),
			review.GetVehicleType(),
			review.GetCreatedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := r.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
