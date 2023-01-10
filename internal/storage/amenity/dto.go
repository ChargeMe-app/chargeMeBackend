package amenity

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	amenityDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/amenity"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	"time"
)

type AmenityDTO struct {
	AmenityID  string    `db:"id"`
	LocationID string    `db:"location_id"`
	Form       int       `db:"type"`
	CreatedAt  time.Time `db:"created_at"`
}

func NewAmenityFromDTO(dto AmenityDTO) amenityDomain.Amenity {
	return amenityDomain.NewAmenityWithID(
		amenityDomain.AmenityID(dto.AmenityID),
		placeDomain.PlaceID(dto.LocationID),
		dto.Form,
		domain.NewModelFrom(dto.CreatedAt, nil),
	)
}

func NewAmenitiesListFromDTO(dto []AmenityDTO) []amenityDomain.Amenity {
	amenities := make([]amenityDomain.Amenity, 0, len(dto))

	for i := range dto {
		amenities = append(amenities, NewAmenityFromDTO(dto[i]))
	}

	return amenities
}
