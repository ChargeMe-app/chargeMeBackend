package place

import (
	"github.com/google/uuid"
)

type PlaceID string

func (placeID PlaceID) String() string {
	return string(placeID)
}

type Place struct {
	placeID   PlaceID
	name      string
	score     *float32
	longitude float32
	latitude  float32
	access    *int
	iconType  *string
	address   *string
}

func NewPlace(
	name string,
	score *float32,
	longitude float32,
	latitude float32,
	access *int,
	iconType *string,
	address *string,
) Place {
	p := Place{
		placeID:   PlaceID(uuid.New().String()),
		name:      name,
		longitude: longitude,
		latitude:  latitude,
	}

	p.SetPlaceScore(score)
	p.SetPlaceAccess(access)
	p.SetPlaceIconType(iconType)
	p.SetPlaceAddress(address)

	return p
}

func NewPlaceWithID(
	placeID PlaceID,
	name string,
	score *float32,
	longitude float32,
	latitude float32,
	access *int,
	iconType *string,
	address *string,
) Place {
	p := Place{
		placeID:   placeID,
		name:      name,
		longitude: longitude,
		latitude:  latitude,
	}

	p.SetPlaceScore(score)
	p.SetPlaceAccess(access)
	p.SetPlaceIconType(iconType)
	p.SetPlaceAddress(address)

	return p
}

func (p Place) GetPlaceID() PlaceID {
	return p.placeID
}

func (p Place) GetPlaceName() string {
	return p.name
}

func (p Place) GetPlaceScore() *float32 {
	return p.score
}

func (p *Place) SetPlaceScore(score *float32) {
	p.score = score
}

func (p Place) GetPlaceLongitude() float32 {
	return p.longitude
}

func (p Place) GetPlaceLatitude() float32 {
	return p.latitude
}

func (p Place) GetPlaceAccess() *int {
	return p.access
}

func (p *Place) SetPlaceAccess(access *int) {
	p.access = access
}

func (p Place) GetPlaceIconType() *string {
	return p.iconType
}

func (p *Place) SetPlaceIconType(iconType *string) {
	p.iconType = iconType
}

func (p Place) GetPlaceAddress() *string {
	return p.address
}

func (p *Place) SetPlaceAddress(address *string) {
	p.address = address
}
