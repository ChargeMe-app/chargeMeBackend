package amenity

import (
	"context"

	amenityDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/amenity"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
)

type AmenityService interface {
	CreateAmenity(context.Context, amenityDomain.Amenity) error
	GetAmenitiesListByLocationID(context.Context, placeDomain.PlaceID) ([]amenityDomain.Amenity, error)
	DeleteAmenitiesByLocationID(context.Context, placeDomain.PlaceID) error
	CreateAmenitiesList(context.Context, []amenityDomain.Amenity) error
}
type AmenityStorage interface {
	CreateAmenity(context.Context, amenityDomain.Amenity) error
	GetAmenitiesListByLocationID(context.Context, placeDomain.PlaceID) ([]amenityDomain.Amenity, error)
	DeleteAmenitiesByLocationID(context.Context, placeDomain.PlaceID) error
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
	return s.amenityStorage.CreateAmenity(ctx, amenity)
}

func (s *service) CreateAmenitiesList(ctx context.Context, amenitiesList []amenityDomain.Amenity) error {
	for _, a := range amenitiesList {
		err := s.amenityStorage.CreateAmenity(ctx, a)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) GetAmenitiesListByLocationID(
	ctx context.Context,
	placeID placeDomain.PlaceID,
) ([]amenityDomain.Amenity, error) {
	return s.amenityStorage.GetAmenitiesListByLocationID(ctx, placeID)
}

func (s *service) DeleteAmenitiesByLocationID(ctx context.Context, placeID placeDomain.PlaceID) error {
	return s.amenityStorage.DeleteAmenitiesByLocationID(ctx, placeID)
}
