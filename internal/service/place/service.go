package place

import (
	"context"
	"log"

	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"

	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
)

type ReviewStorage interface {
	GetReviewWithPositiveRating(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
	GetReviewsWithNotNullRating(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
}

type PlaceService interface {
	CreatePlace(ctx context.Context, place placeDomain.Place) error
	GetPlacesByCoordinates(context.Context, float32, float32, float32, float32) ([]placeDomain.Place, error)
	GetFullPlaceByID(context.Context, placeDomain.PlaceID) (placeDomain.Place, error)
	UpdatePlace(context.Context, placeDomain.Place) error
}

type PlaceStorage interface {
	CreatePlace(ctx context.Context, place placeDomain.Place) error
	GetPlacesByCoordinates(context.Context, float32, float32, float32, float32) ([]placeDomain.Place, error)
	GetFullPlaceByID(context.Context, placeDomain.PlaceID) (placeDomain.Place, error)
	UpdatePlaceScoreByID(context.Context, placeDomain.PlaceID, float32) error
	UpdatePlace(context.Context, placeDomain.Place) error
}

type service struct {
	placeStorage  PlaceStorage
	reviewStorage ReviewStorage
}

func NewPlaceService(placeStorage PlaceStorage, reviewStorage ReviewStorage) PlaceService {
	return &service{
		placeStorage:  placeStorage,
		reviewStorage: reviewStorage,
	}
}

func (s *service) CreatePlace(ctx context.Context, place placeDomain.Place) error {
	log.Println("place.service.CreatePlace")

	return s.placeStorage.CreatePlace(ctx, place)
}

func (s *service) GetPlacesByCoordinates(
	ctx context.Context,
	minLongitude float32,
	maxLongitude float32,
	minLatitude float32,
	maxLatitude float32,
) ([]placeDomain.Place, error) {
	return s.placeStorage.GetPlacesByCoordinates(ctx, minLongitude, maxLongitude, minLatitude, maxLatitude)
}

func (s *service) GetFullPlaceByID(ctx context.Context, place placeDomain.PlaceID) (placeDomain.Place, error) {
	return s.placeStorage.GetFullPlaceByID(ctx, place)
}

func (s *service) UpdatePlace(ctx context.Context, place placeDomain.Place) error {
	return s.placeStorage.UpdatePlace(ctx, place)
}
