package place

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
)

type PhotoID string

func (p PhotoID) String() string {
	return string(p)
}

type Photo struct {
	id      PhotoID
	userID  userDomain.UserID
	placeID PlaceID
	name    string
	caption *string
	domain.Model
}

func NewPhoto(
	id PhotoID,
	userID userDomain.UserID,
	placeID PlaceID,
	name string,
	caption *string,
) Photo {
	return Photo{
		id:      id,
		userID:  userID,
		placeID: placeID,
		name:    name,
		caption: caption,
		Model:   domain.NewModel(),
	}
}

func NewModelWithModel(
	id PhotoID,
	userID userDomain.UserID,
	placeID PlaceID,
	name string,
	caption *string,
	model domain.Model,
) Photo {
	return Photo{
		id:      id,
		userID:  userID,
		placeID: placeID,
		name:    name,
		caption: caption,
		Model:   model,
	}
}

func (p Photo) GetPhotoID() PhotoID {
	return p.id
}

func (p Photo) GetUserID() userDomain.UserID {
	return p.userID
}

func (p Photo) GetPlaceID() PlaceID {
	return p.placeID
}

func (p Photo) GetPhotoName() string {
	return p.name
}

func (p Photo) GetPhotoCaption() *string {
	return p.caption
}

func (p *Photo) SetPhotoCaption(caption *string) {
	p.caption = caption
}
