package user

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/ignishub/terr"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

const (
	TableUsers             = "users"
	TableVehicles          = "vehicles"
	TableAppleCredentials  = "apple_users"
	TableGoogleCredentials = "google_users"
)

type Storage interface {
	CreateUser(context.Context, userDomain.User) error
	IsUserExist(context.Context, userDomain.User) (*bool, error)
	GetUserByIdentifier(context.Context, string) (userDomain.User, error)
	CreateVehicle(context.Context, userDomain.Vehicle) error
	GetVehiclesByUserId(context.Context, userDomain.UserId) ([]userDomain.Vehicle, error)
	CreateAppleCredentials(context.Context, userDomain.AppleCredentials) error
	CreateGoogleCredentials(context.Context, userDomain.GoogleCredentials) error
}

func NewUserStorage(db libdb.DB) Storage {
	return &userStorage{db: db}
}

type userStorage struct {
	db libdb.DB
}

func (u *userStorage) CreateUser(ctx context.Context, user userDomain.User) error {
	query := squirrel.Insert(TableUsers).
		Columns(
			"id",
			"user_identifier",
			"display_name",
			"email",
			"photo_url",
			"sign_type",
			"created_at",
		).
		Values(
			user.GetUserId().String(),
			user.GetUserIdentifier(),
			user.GetDisplayName(),
			user.GetEmail(),
			user.GetPhotoUrl(),
			user.GetSignType(),
			user.GetCreatedAt(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := u.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (u *userStorage) GetUserByIdentifier(ctx context.Context, userIdentifier string) (userDomain.User, error) {
	query := squirrel.Select(
		"id",
		"user_identifier",
		"display_name",
		"email",
		"photo_url",
		"sign_type",
		"created_at",
	).
		From(TableUsers).
		Where(squirrel.Eq{"user_identifier": userIdentifier}).
		PlaceholderFormat(squirrel.Dollar)

	var result UserDTO

	err := u.db.Get(ctx, query, &result)
	if err != nil {
		return userDomain.User{}, err
	}

	return NewUserFromDTO(result), nil
}

func (u *userStorage) IsUserExist(ctx context.Context, user userDomain.User) (*bool, error) {
	query := squirrel.Select("1").
		Prefix("SELECT EXISTS (").
		From(TableUsers).
		Where(squirrel.Eq{"user_identifier": user.GetUserIdentifier()}).
		Suffix(")").
		PlaceholderFormat(squirrel.Dollar)

	var isExist bool

	err := u.db.Get(ctx, query, &isExist)
	if err != nil {
		return nil, err
	}

	return &isExist, nil
}

func (u *userStorage) CreateVehicle(ctx context.Context, vehicle userDomain.Vehicle) error {
	query := squirrel.Insert(TableVehicles).
		Columns(
			"user_id",
			"vehicle_type",
		).
		Values(
			vehicle.GetUserId().String(),
			vehicle.GetVehicleType(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := u.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (u *userStorage) GetVehiclesByUserId(ctx context.Context, userId userDomain.UserId) ([]userDomain.Vehicle, error) {
	queru := squirrel.Select(
		"user_id",
		"vehicle_type",
	).
		From(TableVehicles).
		Where(squirrel.Eq{"user_id": userId.String()}).
		PlaceholderFormat(squirrel.Dollar)

	var result []VehicleDTO

	err := u.db.Select(ctx, queru, &result)
	if err != nil && err == terr.NotFound() {
		return nil, nil
	} else if err != nil && err != terr.NotFound() {
		return nil, err
	}

	return NewVehiclesFromDTO(result), nil
}

func (u *userStorage) CreateAppleCredentials(ctx context.Context, creds userDomain.AppleCredentials) error {
	query := squirrel.Insert(TableAppleCredentials).
		Columns(
			"user_id",
			"authorization_code",
			"identity_token",
		).
		Values(
			creds.GetUserId().String(),
			creds.GetAuthorizationCode(),
			creds.GetIdentityToken(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := u.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (u *userStorage) CreateGoogleCredentials(ctx context.Context, creds userDomain.GoogleCredentials) error {
	query := squirrel.Insert(TableGoogleCredentials).
		Columns(
			"user_id",
			"id_token",
			"access_token",
		).
		Values(
			creds.GetUserId().String(),
			creds.GetIdToken(),
			creds.GetAccessToken(),
		).
		PlaceholderFormat(squirrel.Dollar)

	err := u.db.Insert(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
