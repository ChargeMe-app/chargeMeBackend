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
	stationID       StationID
	placeID         place.PlaceID
	available       *int
	cost            *int
	name            *string
	manufacturer    *string
	costDescription *string
	hours           *string
	kilowatts       *float32
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
	available *int,
	cost *int,
	name *string,
	manufacturer *string,
	costDescription *string,
	hours *string,
	kilowatts *float32,
) Station {
	s := Station{
		stationID: stationID,
		placeID:   placeID,
	}

	s.SetStationAvailability(available)
	s.SetStationCost(cost)
	s.SetStationName(name)
	s.SetStationManufacturer(manufacturer)
	s.SetStationCostDescription(costDescription)
	s.SetStationWorkingHours(hours)
	s.SetStationKilowatts(kilowatts)

	return s
}

func (s *Station) GetStationID() StationID {
	return s.stationID
}

func (s *Station) GetPlaceID() place.PlaceID {
	return s.placeID
}

func (s *Station) SetStationAvailability(available *int) {
	s.available = available
}

func (s *Station) GetStationAvailability() *int {
	return s.available
}

func (s *Station) SetStationCost(cost *int) {
	s.cost = cost
}

func (s *Station) GetStationCost() *int {
	return s.cost
}

func (s *Station) SetStationName(name *string) {
	s.name = name
}

func (s *Station) GetStationName() *string {
	return s.name
}

func (s *Station) SetStationManufacturer(manufacturer *string) {
	s.manufacturer = manufacturer
}

func (s *Station) GetStationManufacturer() *string {
	return s.manufacturer
}

func (s *Station) SetStationCostDescription(costDescription *string) {
	s.costDescription = costDescription
}

func (s *Station) GetStationCostDescription() *string {
	return s.costDescription
}

func (s *Station) SetStationWorkingHours(hours *string) {
	s.hours = hours
}

func (s *Station) GetStationWorkingHours() *string {
	return s.hours
}

func (s *Station) SetStationKilowatts(kilowatts *float32) {
	s.kilowatts = kilowatts
}

func (s *Station) GetStationKilowatts() *float32 {
	return s.kilowatts
}
