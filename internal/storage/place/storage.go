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
	GetPlaces(context.Context, float32, float32, float32, float32) ([]placeDomain.Place, error)
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
			"icon_type",
		).
		Values(
			place.GetPlaceID().String(),
			place.GetPlaceName(),
			place.GetPlaceScore(),
			place.GetPlaceLongitude(),
			place.GetPlaceLatitude(),
			place.GetPlaceAddress(),
			place.GetPlaceAccess(),
			place.GetPlaceIconType(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (s *placeStorage) GetPlaces(
	ctx context.Context,
	minLongitude float32,
	maxLongitude float32,
	minLatitude float32,
	maxLatitude float32,
) ([]placeDomain.Place, error) {
	query := squirrel.Select(
		"id",
		"name",
		"score",
		"longitude",
		"latitude",
		"address",
		"access",
		"icon_type",
	).
		From(tablePlaces).
		Where(squirrel.Expr("longitude > ? AND longitude < ?", minLongitude, maxLongitude)).
		Where(squirrel.Expr("latitude > ? AND latitude < ?", minLatitude, maxLatitude)).
		PlaceholderFormat(squirrel.Dollar)

	var result []PlaceDTO

	err := s.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewPlaceListDTO(result), nil
}
