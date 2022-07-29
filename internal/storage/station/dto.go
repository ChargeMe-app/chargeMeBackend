package station

import (
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
)

type StationDTO struct {
	StationID string `db:"id"`
	PlaceID   string `db:"location_id"`
}

func NewStationFromDTO(dto StationDTO) stationDomain.Station {
	return stationDomain.NewStationWithID(
		stationDomain.StationID(dto.StationID),
		placeDomain.PlaceID(dto.PlaceID),
	)
}

func NewStationListFromDTO(dto []StationDTO) []stationDomain.Station {
	stations := make([]stationDomain.Station, 0, len(dto))

	for i := range dto {
		stations = append(stations, NewStationFromDTO(dto[i]))
	}

	return stations
}
