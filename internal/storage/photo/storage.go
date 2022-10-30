package photo

import (
	"context"
	"github.com/Masterminds/squirrel"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	photoTable = "location_photos"
)

type Storage interface {
	CreatePhoto(context.Context, placeDomain.Photo) error
	GetPhotoListByPlaceID(context.Context, placeDomain.PlaceID) ([]placeDomain.Photo, error)
}

func NewPhotoStorage(db libdb.DB) Storage {
	return &photoStorage{db: db}
}

type photoStorage struct {
	db libdb.DB
}

func (p *photoStorage) GetPhotoListByPlaceID(ctx context.Context, placeID placeDomain.PlaceID) ([]placeDomain.Photo, error) {
	query := squirrel.Select(
		"id",
		"user_id",
		"name",
		"location_id",
		"caption",
		"created_at",
	).
		From(photoTable).
		Where(squirrel.Eq{"location_id": placeID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var result []PhotoDTO

	err := p.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewPhotoListFromDTO(result), nil
}

func (p *photoStorage) CreatePhoto(ctx context.Context, photo placeDomain.Photo) error {
	query := squirrel.Insert(photoTable).
		Columns(
			"id",
			"user_id",
			"name",
			"location_id",
			"caption",
			"created_at",
		).
		Values(
			photo.GetPhotoID().String(),
			photo.GetUserID().String(),
			photo.GetPhotoName(),
			photo.GetPlaceID().String(),
			photo.GetPhotoCaption(),
			photo.GetCreatedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := p.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
