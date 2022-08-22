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
	DisplayName    string    `db:"display_name"`
	Email          string    `db:"email"`
	PhotoUrl       *string   `db:"photo_url"`
	SignType       string    `db:"sign_type"`
	CreatedAt      time.Time `db:"created_at"`
}

func NewUserFromDTO(dto UserDTO) userDomain.User {
	model := domain.NewModelFrom(dto.CreatedAt, nil)

	return userDomain.NewUserWithId(
		userDomain.UserId(dto.UserId),
		dto.DisplayName,
		dto.Email,
		dto.UserIdentifier,
		dto.PhotoUrl,
		dto.SignType,
		model,
	)
}
