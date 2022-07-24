package station

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
)

type StationID string

func (stationID StationID) String() string {
	return string(stationID)
}

type Station struct {
	stationID StationID
	placeID   place.PlaceID
}

func NewStation(placeID place.PlaceID) Station {
	return Station{
		stationID: StationID(uuid.New().String()),
		placeID:   placeID,
	}
}

func NewStationWithID(
	stationID StationID,
	placeID place.PlaceID,
) Station {
	return Station{
		stationID: stationID,
		placeID:   placeID,
	}
}

func (s Station) GetStationID() StationID {
	return s.stationID
}

func (s Station) GetPlaceID() place.PlaceID {
	return s.placeID
}
