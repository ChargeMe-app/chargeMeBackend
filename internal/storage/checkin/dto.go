package checkin

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	checkinDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/checkin"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"time"
)

type CheckinDTO struct {
	CheckinID   uuid.UUID `db:"id"`
	UserID      uuid.UUID `db:"user_id"`
	StationID   string    `db:"station_id"`
	OutletID    string    `db:"outlet_id"`
	UserName    string    `db:"user_name"`
	Duration    int       `db:"duration"`
	VehicleType *int      `db:"vehicle_type"`
	Comment     *string   `db:"comment"`
	Kilowatts   *float32  `db:"kilowatts"`
	Rating      int       `db:"rating"`
	StartedAt   time.Time `db:"started_at"`
	FinishedAt  time.Time `db:"finished_at"`
}

func NewCheckinFromDTO(dto CheckinDTO) checkinDomain.Checkin {
	return checkinDomain.NewCheckinWithID(
		checkinDomain.CheckinID(dto.CheckinID),
		userDomain.UserID(dto.UserID),
		stationDomain.StationID(dto.StationID),
		outletDomain.OutletID(dto.OutletID),
		dto.UserName,
		dto.Duration,
		dto.VehicleType,
		dto.Comment,
		dto.Kilowatts,
		dto.Rating,
		dto.FinishedAt,
		domain.NewModelFrom(dto.StartedAt, nil),
	)
}

func NewCheckinListFromDTO(dto []CheckinDTO) []checkinDomain.Checkin {
	var result []checkinDomain.Checkin

	for _, i := range dto {
		checkin := NewCheckinFromDTO(i)
		result = append(result, checkin)
	}

	return result
}
