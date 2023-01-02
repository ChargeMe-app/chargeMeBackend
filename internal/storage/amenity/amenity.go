package amenity

import (
	"context"

	"github.com/Masterminds/squirrel"
	amenityDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/amenity"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	tableAmenities = "amenities"
)

type Storage interface {
	CreateAmenity(context.Context, amenityDomain.Amenity) error
	GetAmenitiesListByLocationID(context.Context, placeDomain.PlaceID) ([]amenityDomain.Amenity, error)
	DeleteAmenitiesByLocationID(context.Context, placeDomain.PlaceID) error
}

func NewAmenityStorage(db libdb.DB) Storage {
	return &amenityStorage{db: db}
}

type amenityStorage struct {
	db libdb.DB
}

func (a *amenityStorage) CreateAmenity(ctx context.Context, amenity amenityDomain.Amenity) error {
	query := squirrel.Insert(tableAmenities).
		Columns(
			"id",
			"location_id",
			"type",
			"created_at",
		).
		Values(
			amenity.GetAmenityID().String(),
			amenity.GetLocationID(),
			amenity.GetAmenityForm(),
			amenity.GetCreatedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := a.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (a *amenityStorage) GetAmenitiesListByLocationID(ctx context.Context, placeID placeDomain.PlaceID) ([]amenityDomain.Amenity, error) {
	query := squirrel.Select(
		"id",
		"location_id",
		"type",
		"created_at",
	).
		From(tableAmenities).
		Where(squirrel.Eq{"location_id": placeID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var result []AmenityDTO

	err := a.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewAmenitiesListFromDTO(result), nil
}

func (a *amenityStorage) DeleteAmenitiesByLocationID(ctx context.Context, placeID placeDomain.PlaceID) error {
	query := squirrel.Delete(tableAmenities).
		Where(squirrel.Eq{"location_id": placeID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	err := a.db.Delete(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
