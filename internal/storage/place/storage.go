package place

import (
	"context"

	"github.com/Masterminds/squirrel"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	tablePlaces = "places"
)

type Storage interface {
	CreatePlace(context.Context, placeDomain.Place) error
}

func NewPlaceStorage(db libdb.DB) Storage {
	return &placeStorage{db: db}
}

type placeStorage struct {
	db libdb.DB
}

func (s *placeStorage) CreatePlace(ctx context.Context, place placeDomain.Place) error {
	query := squirrel.Insert(tablePlaces).
		Columns(
			"id",
			"name",
			"score",
			"longitude",
			"latitude",
			"address",
			"access",
			"icon_link",
		).
		Values(
			place.GetPlaceID().String(),
			place.GetPlaceName(),
			place.GetPlaceName(),
			place.GetPlaceLongitude(),
			place.GetPlaceLatitude(),
			place.GetPlaceAddress(),
			place.GetPlaceAccess(),
			place.GetPlaceIconLink(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
