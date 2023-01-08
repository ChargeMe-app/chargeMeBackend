package amenity

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
)

type AmenityID string

func (amenityID AmenityID) String() string {
	return string(amenityID)
}

type Amenity struct {
	id         AmenityID
	locationID placeDomain.PlaceID
	form       int
	domain.Model
}

func NewAmenity(
	locationID placeDomain.PlaceID,
	form int,
	model domain.Model,
) Amenity {
	return Amenity{
		id:         AmenityID(uuid.New().String()),
		locationID: locationID,
		form:       form,
		Model:      model,
	}
}

func NewAmenityWithID(
	id AmenityID,
	locationID placeDomain.PlaceID,
	form int,
	model domain.Model,
) Amenity {
	return Amenity{
		id:         id,
		locationID: locationID,
		form:       form,
		Model:      model,
	}
}

func (a *Amenity) GetAmenityID() AmenityID {
	return a.id
}

func (a *Amenity) GetLocationID() placeDomain.PlaceID {
	return a.locationID
}

func (a *Amenity) GetAmenityForm() int {
	return a.form
}
