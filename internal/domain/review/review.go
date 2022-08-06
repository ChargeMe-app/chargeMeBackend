package review

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
)

type ReviewID string

func (reviewID ReviewID) String() string {
	return string(reviewID)
}

type Review struct {
	id          ReviewID
	stationID   stationDomain.StationID
	outletID    outletDomain.OutletID
	comment     *string
	rating      *int
	vehicleName *string
	vehicleType *string
	domain.Model
}

func NewReview(
	stationID stationDomain.StationID,
	outletID outletDomain.OutletID,
	comment *string,
	rating *int,
	vehicleName *string,
	vehicleType *string,
	model domain.Model,
) Review {
	return Review{
		id:          ReviewID(uuid.New().String()),
		stationID:   stationID,
		outletID:    outletID,
		comment:     comment,
		rating:      rating,
		vehicleName: vehicleName,
		vehicleType: vehicleType,
		Model:       model,
	}
}

func NewReviewWithID(
	reviewID ReviewID,
	stationID stationDomain.StationID,
	outletID outletDomain.OutletID,
	comment *string,
	rating *int,
	vehicleName *string,
	vehicleType *string,
	model domain.Model,
) Review {
	return Review{
		id:          reviewID,
		stationID:   stationID,
		outletID:    outletID,
		comment:     comment,
		rating:      rating,
		vehicleName: vehicleName,
		vehicleType: vehicleType,
		Model:       model,
	}
}

func (r *Review) GetReviewID() ReviewID {
	return r.id
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

func (r *Review) SetVehicleType(vehicleType *string) {
	r.vehicleType = vehicleType
}

func (r *Review) GetVehicleType() *string {
	return r.vehicleType
}
