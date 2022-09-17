package checkin

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
)

type CheckinID uuid.UUID

func (c CheckinID) String() string {
	return uuid.UUID(c).String()
}

type Checkin struct {
	id          CheckinID
	userId      userDomain.UserID
	stationId   stationDomain.StationID
	outletId    outletDomain.OutletID
	duration    int
	vehicleType *int
	comment     *string
	kilowatts   *float32
	rating      int
	domain.Model
}

func NewCheckin(
	userId userDomain.UserID,
	stationId stationDomain.StationID,
	outletId outletDomain.OutletID,
	duration int,
	vehicleType *int,
	comment *string,
	kilowatts *float32,
	rating int,
) Checkin {
	return Checkin{
		id:          CheckinID(uuid.New()),
		userId:      userId,
		stationId:   stationId,
		outletId:    outletId,
		duration:    duration,
		vehicleType: vehicleType,
		comment:     comment,
		kilowatts:   kilowatts,
		rating:      rating,
		Model:       domain.NewModel(),
	}
}

func NewCheckinWithID(
	checkinId CheckinID,
	userId userDomain.UserID,
	stationId stationDomain.StationID,
	outletId outletDomain.OutletID,
	duration int,
	vehicleType *int,
	comment *string,
	kilowatts *float32,
	rating int,
	model domain.Model,
) Checkin {
	return Checkin{
		id:          checkinId,
		userId:      userId,
		stationId:   stationId,
		outletId:    outletId,
		duration:    duration,
		vehicleType: vehicleType,
		comment:     comment,
		kilowatts:   kilowatts,
		rating:      rating,
		Model:       model,
	}
}

func (c Checkin) GetCheckinId() CheckinID {
	return c.id
}

func (c Checkin) GetUserId() userDomain.UserID {
	return c.userId
}

func (c Checkin) GetStationId() stationDomain.StationID {
	return c.stationId
}

func (c Checkin) GetOutletId() outletDomain.OutletID {
	return c.outletId
}

func (c Checkin) GetDuration() int {
	return c.duration
}

func (c Checkin) GetVehicleType() *string {
	return c.comment
}

func (c Checkin) SetVehicleType(vehicleType *int) {
	c.vehicleType = vehicleType
}

func (c Checkin) GetComment() *string {
	return c.comment
}

func (c Checkin) SetComment(comment *string) {
	c.comment = comment
}

func (c Checkin) GetKilowatts() *float32 {
	return c.kilowatts
}

func (c Checkin) SetKilowatts(kilowatts *float32) {
	c.kilowatts = kilowatts
}

func (c Checkin) GetRating() int {
	return c.rating
}
