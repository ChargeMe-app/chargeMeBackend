package place

import (
	"github.com/google/uuid"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
)

type PlaceID string

func (placeID PlaceID) String() string {
	return placeID.String()
}

type Place struct {
	placeID   PlaceID
	name      string
	score     *float64
	longitude float64
	latitude  float64
	stations  []stationDomain.Station
}

func NewPlace(
	name string,
	score *float64,
	longitude float64,
	latitude float64,
	stations []stationDomain.Station,
) Place {
	p := Place{
		placeID:   PlaceID(uuid.New().String()),
		name:      name,
		longitude: longitude,
		latitude:  latitude,
		stations:  stations,
	}

	p.SetPlaceScore(score)

	return p
}

func NewPlaceWithID(
	placeID PlaceID,
	name string,
	score *float64,
	longitude float64,
	latitude float64,
	stations []stationDomain.Station,
) Place {
	p := Place{
		placeID:   placeID,
		name:      name,
		longitude: longitude,
		latitude:  latitude,
		stations:  stations,
	}

	p.SetPlaceScore(score)

	return p
}

func (p Place) GetPlaceID() PlaceID {
	return p.placeID
}

func (p Place) GetPlaceName() string {
	return p.name
}

func (p Place) GetPlaceScore() *float64 {
	return p.score
}

func (p *Place) SetPlaceScore(score *float64) {
	p.score = score
}

func (p Place) GetPlaceLongitude() float64 {
	return p.longitude
}

func (p Place) GetPlaceLatitude() float64 {
	return p.latitude
}
