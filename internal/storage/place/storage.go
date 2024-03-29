package place

import (
	"context"

	"github.com/Masterminds/squirrel"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	TablePlaces = "places"
)

type Storage interface {
	CreatePlace(context.Context, placeDomain.Place) error
	GetPlacesByCoordinates(context.Context, float32, float32, float32, float32) ([]placeDomain.Place, error)
	GetFullPlaceByID(context.Context, placeDomain.PlaceID) (placeDomain.Place, error)
	UpdatePlaceScoreByID(context.Context, placeDomain.PlaceID, float32) error
	DeletePlaceByID(ctx context.Context, placeID placeDomain.PlaceID) error
	UpdatePlace(context.Context, placeDomain.Place) error
	GetAllPlaces(ctx context.Context) ([]placeDomain.Place, error)
	HideCompanyPlaces(context.Context, string) error
	CreateOrUnhidePlace(context.Context, placeDomain.Place) error
}

func NewPlaceStorage(db libdb.DB) Storage {
	return &placeStorage{db: db}
}

type placeStorage struct {
	db libdb.DB
}

func (s *placeStorage) GetAllPlaces(ctx context.Context) ([]placeDomain.Place, error) {
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
		From(TablePlaces).
		PlaceholderFormat(squirrel.Dollar)

	var result []PlaceDTO

	err := s.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewPlaceListDTO(result), nil
}

func (s *placeStorage) CreatePlace(ctx context.Context, place placeDomain.Place) error {
	query := squirrel.Insert(TablePlaces).
		Columns(
			"id",
			"name",
			"score",
			"longitude",
			"latitude",
			"address",
			"access",
			"icon_type",
			"description",
			"cost",
			"cost_description",
			"hours",
			"open247",
			"is_open_or_active",
			"phone_number",
			"company_name",
			"created_at",
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
			place.GetDescription(),
			place.GetCost(),
			place.GetCostDescription(),
			place.GetHours(),
			place.GetOpen247(),
			place.IsComingSoon(),
			place.GetPhoneNumber(),
			place.GetCompanyName(),
			place.GetCreatedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (s *placeStorage) UpdatePlace(
	ctx context.Context,
	place placeDomain.Place,
) error {
	query := squirrel.Update(TablePlaces).
		Set("name", place.GetPlaceName()).
		Set("description", place.GetDescription()).
		Set("phone_number", place.GetPhoneNumber()).
		Set("longitude", place.GetPlaceLongitude()).
		Set("latitude", place.GetPlaceLatitude()).
		Set("address", place.GetPlaceAddress()).
		Set("access", place.GetPlaceAccess()).
		Set("cost", place.GetCost()).
		Set("cost_description", place.GetCostDescription()).
		Set("hours", place.GetHours()).
		Set("is_open_or_active", place.IsComingSoon()).
		Set("open247", place.GetOpen247()).
		Where(squirrel.Eq{"id": place.GetPlaceID().String()}).
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Update(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (s *placeStorage) GetPlacesByCoordinates(
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
		From(TablePlaces).
		Where(squirrel.Expr("longitude > ? AND longitude < ?", minLongitude, maxLongitude)).
		Where(squirrel.Expr("latitude > ? AND latitude < ?", minLatitude, maxLatitude)).
		Where(squirrel.Eq{"hide": false}).
		PlaceholderFormat(squirrel.Dollar)

	var result []PlaceDTO

	err := s.db.Select(ctx, query, &result)
	if err != nil {
		return nil, err
	}

	return NewPlaceListDTO(result), nil
}

func (s *placeStorage) GetFullPlaceByID(ctx context.Context, placeID placeDomain.PlaceID) (placeDomain.Place, error) {
	query := squirrel.Select(
		"id",
		"name",
		"score",
		"longitude",
		"latitude",
		"address",
		"access",
		"icon_type",
		"description",
		"cost",
		"cost_description",
		"hours",
		"open247",
		"is_open_or_active",
		"phone_number",
		"company_name",
		"created_at",
	).
		From(TablePlaces).
		Where(squirrel.Eq{"id": placeID.String()}).
		Where(squirrel.Eq{"hide": false}).
		PlaceholderFormat(squirrel.Dollar)

	var result PlaceDTO

	err := s.db.Get(ctx, query, &result)
	if err != nil {
		return placeDomain.Place{}, err
	}

	return NewPlaceFromDTO(result), nil
}

func (s *placeStorage) UpdatePlaceScoreByID(ctx context.Context, placeID placeDomain.PlaceID, rating float32) error {
	query := squirrel.Update(TablePlaces).
		Set("score", rating).
		Where(squirrel.Eq{"id": placeID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Update(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (s *placeStorage) DeletePlaceByID(ctx context.Context, placeID placeDomain.PlaceID) error {
	query := squirrel.Delete(TablePlaces).
		Where(squirrel.Eq{"id": placeID.String()}).
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Delete(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (s *placeStorage) HideCompanyPlaces(ctx context.Context, companyName string) error {
	query := squirrel.Update(TablePlaces).
		Set("hide", true).
		Where(squirrel.Eq{"company_name": companyName}).
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Update(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (s *placeStorage) CreateOrUnhidePlace(ctx context.Context, place placeDomain.Place) error {
	query := squirrel.Insert(TablePlaces).
		Columns(
			"id",
			"name",
			"score",
			"longitude",
			"latitude",
			"address",
			"access",
			"icon_type",
			"description",
			"cost",
			"cost_description",
			"hours",
			"open247",
			"is_open_or_active",
			"phone_number",
			"company_name",
			"created_at",
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
			place.GetDescription(),
			place.GetCost(),
			place.GetCostDescription(),
			place.GetHours(),
			place.GetOpen247(),
			place.IsComingSoon(),
			place.GetPhoneNumber(),
			place.GetCompanyName(),
			place.GetCreatedAt(),
		).
		Suffix("ON CONFLICT ON CONSTRAINT places_pkey DO UPDATE SET hide = false").
		PlaceholderFormat(squirrel.Dollar)

	err := s.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
