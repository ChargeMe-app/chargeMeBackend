package photo

import (
	"context"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
)

type PhotoService interface {
	CreatePhoto(context.Context, placeDomain.Photo) error
	GetPhotoListByPlaceID(context.Context, placeDomain.PlaceID) ([]placeDomain.Photo, error)
}

type PhotoStorage interface {
	CreatePhoto(context.Context, placeDomain.Photo) error
	GetPhotoListByPlaceID(context.Context, placeDomain.PlaceID) ([]placeDomain.Photo, error)
}

type service struct {
	photoStorage PhotoStorage
}

func NewPhotoService(photoStorage PhotoStorage) PhotoService {
	return &service{
		photoStorage: photoStorage,
	}
}

func (s *service) GetPhotoListByPlaceID(ctx context.Context, placeID placeDomain.PlaceID) ([]placeDomain.Photo, error) {
	return s.photoStorage.GetPhotoListByPlaceID(ctx, placeID)
}

func (s *service) CreatePhoto(ctx context.Context, photo placeDomain.Photo) error {
	return s.photoStorage.CreatePhoto(ctx, photo)
}
