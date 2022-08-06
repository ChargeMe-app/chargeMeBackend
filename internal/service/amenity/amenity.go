package amenity

import (
	"context"
	amenityDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/amenity"
)

type AmenityService interface {
	CreateAmenity(context.Context, amenityDomain.Amenity) error
}
type AmenityStorage interface {
	CreateAmenity(context.Context, amenityDomain.Amenity) error
}

type service struct {
	amenityStorage AmenityStorage
}

func NewAmenityService(amenityStorage AmenityStorage) AmenityService {
	return &service{
		amenityStorage: amenityStorage,
	}
}

func (s *service) CreateAmenity(ctx context.Context, amenity amenityDomain.Amenity) error {
	return s.CreateAmenity(ctx, amenity)
}
