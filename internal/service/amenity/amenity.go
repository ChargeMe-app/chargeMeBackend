package amenity

import (
	"context"
	amenityDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/amenity"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	"log"
)

type AmenityService interface {
	CreateAmenity(context.Context, amenityDomain.Amenity) error
	GetAmenitiesListByLocationID(context.Context, placeDomain.PlaceID) ([]amenityDomain.Amenity, error)
}
type AmenityStorage interface {
	CreateAmenity(context.Context, amenityDomain.Amenity) error
	GetAmenitiesListByLocationID(context.Context, placeDomain.PlaceID) ([]amenityDomain.Amenity, error)
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
	log.Println("service.amenity.CreateAmenity")

	return s.CreateAmenity(ctx, amenity)
}

func (s *service) GetAmenitiesListByLocationID(
	ctx context.Context,
	placeID placeDomain.PlaceID,
) ([]amenityDomain.Amenity, error) {
	log.Println("service.amenity.GetAmenitiesListByLocationID")

	return s.amenityStorage.GetAmenitiesListByLocationID(ctx, placeID)
}
