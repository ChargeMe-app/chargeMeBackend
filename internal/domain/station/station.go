package station

import (
	"github.com/google/uuid"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
)

type StationID string

func (stationID StationID) String() string {
	return stationID.String()
}

type Station struct {
	stationID StationID
	placeID   place.PlaceID
	outlets   []outletDomain.Outlet
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

func (s Station) GetOutlets() []outletDomain.Outlet {
	return s.outlets
}
