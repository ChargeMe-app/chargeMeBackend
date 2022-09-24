package review

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"time"
)

type ReviewDTO struct {
	ReviewID      string    `db:"id"`
	StationID     string    `db:"station_id"`
	OutletID      string    `db:"outlet_id"`
	UserID        *string   `db:"user_id"`
	Comment       *string   `db:"comment"`
	Rating        *int      `db:"rating"`
	ConnectorType *int      `db:"connector_type"`
	UserName      *string   `db:"user_name"`
	VehicleName   *string   `db:"vehicle_name"`
	VehicleType   *int      `db:"vehicle_type"`
	CreatedAt     time.Time `db:"created_at"`
}

func NewReviewFromDTO(dto ReviewDTO) reviewDomain.Review {
	var userId userDomain.UserID

	if dto.UserID != nil {
		userId = userDomain.UserID(uuid.MustParse(*dto.UserID))
	}

	return reviewDomain.NewReviewWithID(
		reviewDomain.ReviewID(dto.ReviewID),
		stationDomain.StationID(dto.StationID),
		outletDomain.OutletID(dto.OutletID),
		&userId,
		dto.Comment,
		dto.Rating,
		dto.ConnectorType,
		dto.Rating,
		dto.UserName,
		dto.VehicleName,
		dto.VehicleType,
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
