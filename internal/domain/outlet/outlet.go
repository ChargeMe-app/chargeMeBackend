package outlet

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
)

type OutletID string

func (o OutletID) String() string {
	return string(o)
}

type Outlet struct {
	outletID      OutletID
	stationID     station.StationID
	connectorType int
	kilowatts     *float32
	power         int
	price         *float32
	priceUnit     *string
	domain.Model
}

func NewOutlet(
	stationID station.StationID,
	connectorType int,
	kilowatts *float32,
	power int,
	model domain.Model,
) Outlet {
	outlet := Outlet{
		outletID:      OutletID(uuid.New().String()),
		stationID:     stationID,
		connectorType: connectorType,
		kilowatts:     kilowatts,
		power:         power,
		Model:         model,
	}

	return outlet
}

func NewOutletWithID(
	outletID OutletID,
	stationID station.StationID,
	connectorType int,
	kilowatts *float32,
	power int,
	model domain.Model,
) Outlet {
	outlet := Outlet{
		outletID:      outletID,
		stationID:     stationID,
		connectorType: connectorType,
		kilowatts:     kilowatts,
		power:         power,
		Model:         model,
	}

	return outlet
}

func (o *Outlet) GetOutletID() OutletID {
	return o.outletID
}

func (o *Outlet) GetStationID() station.StationID {
	return o.stationID
}

func (o *Outlet) GetConnector() int {
	return o.connectorType
}

func (o *Outlet) GetKilowatts() *float32 {
	return o.kilowatts
}

func (o *Outlet) SetKilowatts(kilowatts *float32) {
	o.kilowatts = kilowatts
}

func (o *Outlet) GetPower() int {
	return o.power
}

func (o *Outlet) GetPrice() *float32 {
	return o.price
}

func (o *Outlet) SetPrice(price *float32) {
	o.price = price
}

func (o *Outlet) GetPriceUnit() *string {
	return o.priceUnit
}

func (o *Outlet) SetPriceUnit(unit *string) {
	o.priceUnit = unit
}
