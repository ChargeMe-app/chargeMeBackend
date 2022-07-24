package outlet

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
)

type OutletID string

func (o OutletID) String() string {
	return o.String()
}

type Outlet struct {
	outletID      OutletID
	stationID     station.StationID
	connectorType int
	kilowatts     *float64
	power         int
}

func NewOutlet(
	stationID station.StationID,
	connectorType int,
	kilowatts *float64,
	power int,
) Outlet {
	outlet := Outlet{
		outletID:      OutletID(uuid.New().String()),
		stationID:     stationID,
		connectorType: connectorType,
		power:         power,
	}

	outlet.SetKilowatts(kilowatts)

	return outlet
}

func NewOutletWithID(
	outletID OutletID,
	stationID station.StationID,
	connectorType int,
	kilowatts *float64,
	power int,
) Outlet {
	outlet := Outlet{
		outletID:      outletID,
		stationID:     stationID,
		connectorType: connectorType,
		power:         power,
	}

	outlet.SetKilowatts(kilowatts)

	return outlet
}

func (o Outlet) GetOutletID() OutletID {
	return o.outletID
}

func (o Outlet) GetStationID() station.StationID {
	return o.stationID
}

func (o Outlet) GetConnector() int {
	return o.connectorType
}

func (o Outlet) GetKilowatts() *float64 {
	return o.kilowatts
}

func (o *Outlet) SetKilowatts(kilowatts *float64) {
	o.kilowatts = kilowatts
}

func (o Outlet) GetPower() int {
	return o.power
}
