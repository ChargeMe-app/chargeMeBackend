package place

import (
	"context"

	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
)

type PlaceService interface {
	CreatePlace(ctx context.Context, place placeDomain.Place) error
}

type PlaceStorage interface {
	CreatePlace(ctx context.Context, place placeDomain.Place) error
}

type service struct {
	placeStorage PlaceStorage
}

func NewPlaceService(placeStorage PlaceStorage) PlaceService {
	return &service{
		placeStorage: placeStorage,
	}
}

func (s *service) CreatePlace(ctx context.Context, place placeDomain.Place) error {
	return s.placeStorage.CreatePlace(ctx, place)
}
