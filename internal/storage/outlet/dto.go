package outlet

import (
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
)

type OutletDTO struct {
	OutletID  string   `db:"id"`
	StationID string   `db:"station_id"`
	Connector int      `db:"connector"`
	Kilowatts *float64 `db:"kilowatts"`
	Power     int      `db:"power"`
}

func NewOutletFromDTO(dto OutletDTO) outletDomain.Outlet {
	return outletDomain.NewOutletWithID(
		outletDomain.OutletID(dto.OutletID),
		stationDomain.StationID(dto.StationID),
		dto.Connector,
		dto.Kilowatts,
		dto.Power,
	)
}

func NewOutletListFromDTO(dto []OutletDTO) []outletDomain.Outlet {
	outlets := make([]outletDomain.Outlet, 0, len(dto))

	for i := range dto {
		outlets = append(outlets, NewOutletFromDTO(dto[i]))
	}

	return outlets
}
