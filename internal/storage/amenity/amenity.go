package amenity

import (
	"context"
	"github.com/Masterminds/squirrel"
	amenityDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/amenity"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	tableAmenities = "amenities"
)

type Storage interface {
	CreateAmenity(context.Context, amenityDomain.Amenity) error
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
