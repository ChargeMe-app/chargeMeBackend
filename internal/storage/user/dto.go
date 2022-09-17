package user

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"time"
)

type UserDTO struct {
	UserId         uuid.UUID `db:"id"`
	UserIdentifier string    `db:"user_identifier"`
	DisplayName    *string   `db:"display_name"`
	Email          *string   `db:"email"`
	PhotoUrl       *string   `db:"photo_url"`
	SignType       string    `db:"sign_type"`
	CreatedAt      time.Time `db:"created_at"`
}

type VehicleDTO struct {
	UserId      uuid.UUID `db:"user_id"`
	VehicleType int       `db:"vehicle_type"`
}

func NewUserFromDTO(dto UserDTO) userDomain.User {
	model := domain.NewModelFrom(dto.CreatedAt, nil)

	return userDomain.NewUserWithId(
		userDomain.UserID(dto.UserId),
		dto.DisplayName,
		dto.Email,
		dto.UserIdentifier,
		dto.PhotoUrl,
		dto.SignType,
		model,
	)
}

func NewVehicleFromDTO(dto VehicleDTO) userDomain.Vehicle {
	return userDomain.NewVehicle(
		userDomain.UserID(dto.UserId),
		dto.VehicleType,
	)
}

func NewVehiclesFromDTO(dto []VehicleDTO) []userDomain.Vehicle {
	var vehicles []userDomain.Vehicle

	for _, i := range dto {
		vehicles = append(vehicles, NewVehicleFromDTO(i))
	}

	return vehicles
}
