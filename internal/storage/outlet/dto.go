package outlet

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"time"
)

type OutletDTO struct {
	OutletID  string    `db:"id"`
	StationID string    `db:"station_id"`
	Connector int       `db:"connector"`
	Kilowatts *float32  `db:"kilowatts"`
	Power     int       `db:"power"`
	CreatedAt time.Time `db:"created_at"`
}

func NewOutletFromDTO(dto OutletDTO) outletDomain.Outlet {
	model := domain.NewModelFrom(dto.CreatedAt, nil)

	o := outletDomain.NewOutletWithID(
		outletDomain.OutletID(dto.OutletID),
		stationDomain.StationID(dto.StationID),
		dto.Connector,
		dto.Kilowatts,
		dto.Power,
		model,
	)

	return o
}

func NewOutletListFromDTO(dto []OutletDTO) []outletDomain.Outlet {
	outlets := make([]outletDomain.Outlet, 0, len(dto))

	for i := range dto {
		outlets = append(outlets, NewOutletFromDTO(dto[i]))
	}

	return outlets
}
