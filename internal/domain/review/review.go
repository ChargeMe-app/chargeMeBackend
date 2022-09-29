package review

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
)

type ReviewID string

func (reviewID ReviewID) String() string {
	return string(reviewID)
}

type Review struct {
	id            ReviewID
	stationID     stationDomain.StationID
	outletID      outletDomain.OutletID
	userID        *userDomain.UserID
	comment       *string
	rating        *int
	connectorType *int
	kilowatts     *float32
	userName      *string
	vehicleName   *string
	vehicleType   *int
	domain.Model
}

func NewReview(
	stationID stationDomain.StationID,
	outletID outletDomain.OutletID,
	userID *userDomain.UserID,
	comment *string,
	rating *int,
	connectorType *int,
	kilowatts *float32,
	userName *string,
	vehicleName *string,
	vehicleType *int,
	model domain.Model,
) Review {
	return Review{
		id:            ReviewID(uuid.New().String()),
		stationID:     stationID,
		outletID:      outletID,
		userID:        userID,
		comment:       comment,
		rating:        rating,
		connectorType: connectorType,
		kilowatts:     kilowatts,
		userName:      userName,
		vehicleName:   vehicleName,
		vehicleType:   vehicleType,
		Model:         model,
	}
}

func NewReviewWithID(
	reviewID ReviewID,
	stationID stationDomain.StationID,
	outletID outletDomain.OutletID,
	userID *userDomain.UserID,
	comment *string,
	rating *int,
	connectorType *int,
	kilowatts *float32,
	userName *string,
	vehicleName *string,
	vehicleType *int,
	model domain.Model,
) Review {
	return Review{
		id:            reviewID,
		stationID:     stationID,
		outletID:      outletID,
		userID:        userID,
		comment:       comment,
		rating:        rating,
		connectorType: connectorType,
		kilowatts:     kilowatts,
		userName:      userName,
		vehicleName:   vehicleName,
		vehicleType:   vehicleType,
		Model:         model,
	}
}

func (r *Review) GetReviewID() ReviewID {
	return r.id
}

func (r *Review) GetUserID() *userDomain.UserID {
	return r.userID
}

func (r *Review) GetStationID() stationDomain.StationID {
	return r.stationID
}

func (r *Review) GetOutletID() outletDomain.OutletID {
	return r.outletID
}

func (r *Review) SetComment(comment *string) {
	r.comment = comment
}

func (r *Review) GetComment() *string {
	return r.comment
}

func (r *Review) SetRating(rating *int) {
	r.rating = rating
}

func (r *Review) GetRating() *int {
	return r.rating
}

func (r *Review) SetVehicleName(vehicleName *string) {
	r.vehicleName = vehicleName
}

func (r *Review) GetVehicleName() *string {
	return r.vehicleName
}

func (r *Review) SetVehicleType(vehicleType *int) {
	r.vehicleType = vehicleType
}

func (r *Review) GetVehicleType() *int {
	return r.vehicleType
}

func (r *Review) GetConnectorType() *int {
	return r.connectorType
}

func (r *Review) SetConnectorType(connectorType *int) {
	r.connectorType = connectorType
}

func (r *Review) GetUserName() *string {
	return r.userName
}

func (r *Review) SetUserName(userName *string) {
	r.userName = userName
}

func (r *Review) SetKilowatts(kilowatts *float32) {
	r.kilowatts = kilowatts
}

func (r *Review) GetKilowatts() *float32 {
	return r.kilowatts
}
