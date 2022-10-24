package user

import (
	"context"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"log"
)

type UserService interface {
	GetOrCreateUser(context.Context, userDomain.User) (*userDomain.User, error)
	CreateVehicle(context.Context, userDomain.Vehicle) error
	GetVehiclesByUserId(context.Context, userDomain.UserID) ([]userDomain.Vehicle, error)
	GetUserByUserId(context.Context, userDomain.UserID) (userDomain.User, error)
	CreateAppleCredentials(context.Context, userDomain.AppleCredentials) error
	GetUserByIdentifier(context.Context, string) (userDomain.User, error)
	CreateGoogleCredentials(context.Context, userDomain.GoogleCredentials) error
}

type UserStorage interface {
	CreateUser(context.Context, userDomain.User) error
	IsUserExist(context.Context, userDomain.User) (*bool, error)
	GetUserByIdentifier(context.Context, string) (userDomain.User, error)
	GetUserByUserId(context.Context, userDomain.UserID) (userDomain.User, error)
	CreateVehicle(context.Context, userDomain.Vehicle) error
	GetVehiclesByUserId(context.Context, userDomain.UserID) ([]userDomain.Vehicle, error)
	CreateAppleCredentials(context.Context, userDomain.AppleCredentials) error
	CreateGoogleCredentials(context.Context, userDomain.GoogleCredentials) error
}

type service struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) UserService {
	return &service{
		userStorage: userStorage,
	}
}

func (s *service) GetOrCreateUser(ctx context.Context, user userDomain.User) (*userDomain.User, error) {
	log.Println("user.service.GetOrCreateUser")
	ok, err := s.userStorage.IsUserExist(ctx, user)
	if err != nil {
		return nil, err
	}

	if *ok {
		u, err := s.userStorage.GetUserByIdentifier(ctx, user.GetUserIdentifier())

		return &u, err
	}

	return nil, s.userStorage.CreateUser(ctx, user)
}

func (s *service) CreateVehicle(ctx context.Context, vehicle userDomain.Vehicle) error {
	return s.userStorage.CreateVehicle(ctx, vehicle)
}

func (s *service) CreateAppleCredentials(ctx context.Context, creds userDomain.AppleCredentials) error {
	return s.userStorage.CreateAppleCredentials(ctx, creds)
}

func (s *service) CreateGoogleCredentials(ctx context.Context, creds userDomain.GoogleCredentials) error {
	return s.userStorage.CreateGoogleCredentials(ctx, creds)
}

func (s *service) GetVehiclesByUserId(ctx context.Context, userId userDomain.UserID) ([]userDomain.Vehicle, error) {
	return s.userStorage.GetVehiclesByUserId(ctx, userId)
}

func (s *service) GetUserByIdentifier(ctx context.Context, identifier string) (userDomain.User, error) {
	return s.userStorage.GetUserByIdentifier(ctx, identifier)
}

func (s *service) GetUserByUserId(ctx context.Context, userId userDomain.UserID) (userDomain.User, error) {
	return s.userStorage.GetUserByUserId(ctx, userId)
}
