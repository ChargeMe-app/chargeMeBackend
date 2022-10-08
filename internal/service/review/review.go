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
	GetReviewWithPositiveRating(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
	GetReviewsWithNotNullRating(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
}

type ReviewStorage interface {
	CreateReview(context.Context, reviewDomain.Review) error
	GetReviewsListByStationID(context.Context, stationDomain.StationID) ([]reviewDomain.Review, error)
	GetReviewsListByUserID(context.Context, userDomain.UserID) ([]reviewDomain.Review, error)
	GetReviewsListByLocationID(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
	GetReviewWithPositiveRating(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
	GetReviewsWithNotNullRating(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
}

type PlaceStorage interface {
	UpdatePlaceScoreByID(context.Context, placeDomain.PlaceID, float32) error
}

type StationStorage interface {
	GetPlaceIdByStationID(context.Context, stationDomain.StationID) (placeDomain.PlaceID, error)
}

type service struct {
	reviewStorage  ReviewStorage
	placeStorage   PlaceStorage
	stationStorage StationStorage
}

func NewReviewService(reviewStorage ReviewStorage, placeStorage PlaceStorage, stationStorage StationStorage) ReviewService {
	return &service{
		reviewStorage:  reviewStorage,
		placeStorage:   placeStorage,
		stationStorage: stationStorage,
	}
}

func (s *service) CreateReview(ctx context.Context, review reviewDomain.Review) error {
	err := s.reviewStorage.CreateReview(ctx, review)
	if err != nil {
		return err
	}

	placeId, err := s.stationStorage.GetPlaceIdByStationID(ctx, review.GetStationID())
	if err != nil {
		return err
	}

	positiveRatings, err := s.reviewStorage.GetReviewWithPositiveRating(ctx, placeId)
	if err != nil {
		return err
	}

	allRatings, err := s.reviewStorage.GetReviewsWithNotNullRating(ctx, placeId)
	if err != nil {
		return err
	}

	newScore := (float32(len(positiveRatings)) / float32(len(allRatings))) * 10

	return s.placeStorage.UpdatePlaceScoreByID(ctx, placeId, newScore)
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

func (s *service) GetReviewWithPositiveRating(ctx context.Context, placeID placeDomain.PlaceID) ([]reviewDomain.Review, error) {
	return s.reviewStorage.GetReviewWithPositiveRating(ctx, placeID)
}

func (s *service) GetReviewsWithNotNullRating(ctx context.Context, placeID placeDomain.PlaceID) ([]reviewDomain.Review, error) {
	return s.reviewStorage.GetReviewsWithNotNullRating(ctx, placeID)
}
