package review

import (
	"context"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
)

type ReviewService interface {
	CreateReview(context.Context, reviewDomain.Review) error
	GetReviewsListByStationID(context.Context, stationDomain.StationID) ([]reviewDomain.Review, error)
	GetReviewsListByUserID(context.Context, userDomain.UserID) ([]reviewDomain.Review, error)
	GetReviewsListByLocationID(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
}

type ReviewStorage interface {
	CreateReview(context.Context, reviewDomain.Review) error
	GetReviewsListByStationID(context.Context, stationDomain.StationID) ([]reviewDomain.Review, error)
	GetReviewsListByUserID(context.Context, userDomain.UserID) ([]reviewDomain.Review, error)
	GetReviewsListByLocationID(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
}

type service struct {
	reviewStorage ReviewStorage
}

func NewReviewService(reviewStorage ReviewStorage) ReviewService {
	return &service{
		reviewStorage: reviewStorage,
	}
}

func (s *service) CreateReview(ctx context.Context, review reviewDomain.Review) error {
	return s.CreateReview(ctx, review)
}

func (s *service) GetReviewsListByStationID(
	ctx context.Context,
	stationID stationDomain.StationID,
) ([]reviewDomain.Review, error) {
	return s.reviewStorage.GetReviewsListByStationID(ctx, stationID)
}

func (s *service) GetReviewsListByLocationID(
	ctx context.Context,
	placeID placeDomain.PlaceID,
) ([]reviewDomain.Review, error) {
	return s.reviewStorage.GetReviewsListByLocationID(ctx, placeID)
}

func (s *service) GetReviewsListByUserID(ctx context.Context, userId userDomain.UserID) ([]reviewDomain.Review, error) {
	return s.reviewStorage.GetReviewsListByUserID(ctx, userId)
}
