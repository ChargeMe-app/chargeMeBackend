package review

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"time"
)

type ReviewDTO struct {
	ReviewID    string    `db:"id"`
	StationID   string    `db:"station_id"`
	OutletID    string    `db:"outlet_id"`
	Comment     *string   `db:"comment"`
	Rating      *int      `db:"rating"`
	VehicleName *string   `db:"vehicle_name"`
	VehicleType *string   `db:"vehicle_type"`
	CreatedAt   time.Time `db:"created_at"`
}

func NewReviewFromDTO(dto ReviewDTO) reviewDomain.Review {
	return reviewDomain.NewReviewWithID(
		reviewDomain.ReviewID(dto.ReviewID),
		stationDomain.StationID(dto.StationID),
		outletDomain.OutletID(dto.OutletID),
		dto.Comment,
		dto.Rating,
		nil,
		nil,
		domain.NewModelFrom(dto.CreatedAt, nil),
	)
}

func NewReviewsListFromDTO(dto []ReviewDTO) []reviewDomain.Review {
	reviews := make([]reviewDomain.Review, 0, len(dto))

	for i := range dto {
		reviews = append(reviews, NewReviewFromDTO(dto[i]))
	}

	return reviews
}
