package user

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
)

type UserId uuid.UUID

func (u UserId) String() string {
	return uuid.UUID(u).String()
}

type User struct {
	id             UserId
	displayName    string
	email          string
	userIdentifier string
	photoUrl       *string
	signType       string
	domain.Model
}

func NewUser(
	displayName string,
	email string,
	userIdentifier string,
	photoUrl *string,
	signType string,
) User {
	return User{
		UserId(uuid.New()),
		displayName,
		email,
		userIdentifier,
		photoUrl,
		signType,
		domain.NewModel(),
	}
}

func NewUserWithId(
	id UserId,
	displayName string,
	email string,
	userIdentifier string,
	photoUrl *string,
	signType string,
	model domain.Model,
) User {
	return User{
		id,
		displayName,
		email,
		userIdentifier,
		photoUrl,
		signType,
		model,
	}
}

func (u *User) GetUserId() UserId {
	return u.id
}

func (u *User) GetDisplayName() string {
	return u.displayName
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetUserIdentifier() string {
	return u.userIdentifier
}

func (u *User) GetPhotoUrl() *string {
	return u.photoUrl
}

func (u *User) SetPhotoUrl(photoUrl *string) {
	u.photoUrl = photoUrl
}

func (u *User) GetSignType() string {
	return u.signType
}
