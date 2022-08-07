package station

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"time"
)

type StationDTO struct {
	StationID       string    `db:"id"`
	PlaceID         string    `db:"location_id"`
	Available       *int      `db:"available"`
	Cost            *int      `db:"cost"`
	Name            *string   `db:"name"`
	Manufacturer    *string   `db:"manufacturer"`
	CostDescription *string   `db:"cost_description"`
	Hours           *string   `db:"hours"`
	Kilowatts       *float32  `db:"kilowatts"`
	CreatedAt       time.Time `db:"created_at"`
}

func NewStationFromDTO(dto StationDTO) stationDomain.Station {
	model := domain.NewModelFrom(dto.CreatedAt, nil)

	s := stationDomain.NewStationWithID(
		stationDomain.StationID(dto.StationID),
		placeDomain.PlaceID(dto.PlaceID),
		dto.Available,
		dto.Cost,
		dto.Name,
		dto.Manufacturer,
		dto.CostDescription,
		dto.Hours,
		dto.Kilowatts,
		model,
	)

	return s
}

func NewStationListFromDTO(dto []StationDTO) []stationDomain.Station {
	stations := make([]stationDomain.Station, 0, len(dto))

	for i := range dto {
		stations = append(stations, NewStationFromDTO(dto[i]))
	}

	return stations
}
