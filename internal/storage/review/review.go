package review

import (
	"context"
	"github.com/Masterminds/squirrel"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	station "github.com/poorfrombabylon/chargeMeBackend/internal/storage/station"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	tableReviews = "reviews"
)

type Storage interface {
	CreateReview(context.Context, reviewDomain.Review) error
	GetReviewsListByStationID(context.Context, stationDomain.StationID) ([]reviewDomain.Review, error)
	GetReviewsListByUserID(context.Context, userDomain.UserID) ([]reviewDomain.Review, error)
	GetReviewsListByLocationID(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
	GetReviewWithPositiveRating(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
	GetReviewsWithNotNullRating(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
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
			"user_id",
			"comment",
			"station_id",
			"outlet_id",
			"rating",
			"connector_type",
			"user_name",
			"kilowatts",
			"vehicle_type",
			"created_at",
		).
		Values(
			review.GetReviewID().String(),
			review.GetUserID().String(),
			review.GetComment(),
			review.GetStationID().String(),
			review.GetOutletID().String(),
			review.GetRating(),
			review.GetConnectorType(),
			review.GetUserName(),
			review.GetKilowatts(),
			review.GetVehicleType(),
			review.GetCreatedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	//if review.GetUserID() != nil {
	//	query = query.Columns(
	//		"user_id",
	//	).
	//		Values(
	//			review.GetUserID().String(),
	//		)
	//}

	err := r.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (r *reviewStorage) GetReviewsListByStationID(
	ctx context.Context,
	stationID stationDomain.StationID,
) ([]reviewDomain.Review, error) {
	query := squirrel.Select(
		"id",
		"comment",
		"station_id",
		"outlet_id",
		"rating",
		"connector_type",
		"user_name",
		"vehicle_name",
		"vehicle_type",
		"created_at",
	).
		From(tableReviews).
		Where(squirrel.Eq{"station_id": stationID.String()}).
		OrderBy("created_at DESC").
		PlaceholderFormat(squirrel.Dollar)

	var result []ReviewDTO

	err := r.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewReviewsListFromDTO(result), nil
}

func (r *reviewStorage) GetReviewsListByLocationID(
	ctx context.Context,
	placeID placeDomain.PlaceID,
) ([]reviewDomain.Review, error) {
	query := squirrel.Select(
		"r.id",
		"r.comment",
		"r.station_id",
		"r.outlet_id",
		"r.rating",
		"r.connector_type",
		"r.user_name",
		"r.vehicle_name",
		"r.vehicle_type",
		"r.created_at",
	).
		From(tableReviews + " as r").
		Join(station.TableStations + " as s ON s.id = r.station_id").
		Where(squirrel.Eq{"s.location_id": placeID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var result []ReviewDTO

	err := r.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewReviewsListFromDTO(result), nil
}

func (r *reviewStorage) GetReviewsListByUserID(
	ctx context.Context,
	userId userDomain.UserID,
) ([]reviewDomain.Review, error) {
	query := squirrel.Select(
		"id",
		"comment",
		"station_id",
		"outlet_id",
		"user_id",
		"rating",
		"connector_type",
		"user_name",
		"vehicle_name",
		"vehicle_type",
		"created_at",
	).
		From(tableReviews).
		Where(squirrel.Eq{"user_id": userId.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var result []ReviewDTO

	err := r.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewReviewsListFromDTO(result), nil
}

func (r *reviewStorage) GetReviewWithPositiveRating(ctx context.Context, placeID placeDomain.PlaceID) ([]reviewDomain.Review, error) {
	query := squirrel.Select(
		"r.id",
		"r.comment",
		"r.station_id",
		"r.outlet_id",
		"r.rating",
		"r.connector_type",
		"r.user_name",
		"r.vehicle_name",
		"r.vehicle_type",
		"r.created_at",
	).
		From(tableReviews + " as r").
		Join(station.TableStations + " as s ON s.id = r.station_id").
		Where(squirrel.Eq{"s.location_id": placeID.String()}).
		Where(squirrel.Eq{"r.rating": 1}).
		PlaceholderFormat(squirrel.Dollar)

	var result []ReviewDTO

	err := r.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewReviewsListFromDTO(result), nil
}

func (r *reviewStorage) GetReviewsWithNotNullRating(ctx context.Context, placeID placeDomain.PlaceID) ([]reviewDomain.Review, error) {
	query := squirrel.Select(
		"r.id",
		"r.comment",
		"r.station_id",
		"r.outlet_id",
		"r.rating",
		"r.connector_type",
		"r.user_name",
		"r.vehicle_name",
		"r.vehicle_type",
		"r.created_at",
	).
		From(tableReviews + " as r").
		Join(station.TableStations + " as s ON s.id = r.station_id").
		Where(squirrel.Eq{"s.location_id": placeID.String()}).
		Where(squirrel.NotEq{"r.rating": nil}).
		Where(squirrel.NotEq{"r.rating": 0}).
		PlaceholderFormat(squirrel.Dollar)

	var result []ReviewDTO

	err := r.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewReviewsListFromDTO(result), nil
}
