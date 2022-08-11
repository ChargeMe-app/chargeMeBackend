package place

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	"time"
)

type PlaceDTO struct {
	PlaceID                      string    `db:"id"`
	Name                         string    `db:"name"`
	Score                        *float32  `db:"score"`
	Longitude                    float32   `db:"longitude"`
	Latitude                     float32   `db:"latitude"`
	Address                      *string   `db:"address"`
	Access                       *int      `db:"access"`
	IconType                     *string   `db:"icon_type"`
	Description                  *string   `db:"description"`
	AccessRestriction            *string   `db:"access_restriction"`
	AccessRestrictionDescription *string   `db:"access_restriction_description"`
	Cost                         *bool     `db:"cost"`
	CostDescription              *string   `db:"cost_description"`
	Hours                        *string   `db:"hours"`
	Open247                      *bool     `db:"open247"`
	IsOpenOrActive               *bool     `db:"is_open_or_active"`
	PhoneNumber                  *string   `db:"phone_number"`
	CreatedAt                    time.Time `db:"created_at"`
}

func NewPlaceFromDTO(dto PlaceDTO) placeDomain.Place {
	model := domain.NewModelFrom(dto.CreatedAt, nil)

	p := placeDomain.NewPlaceWithID(
		placeDomain.PlaceID(dto.PlaceID),
		dto.Name,
		dto.Score,
		dto.Longitude,
		dto.Latitude,
		dto.Access,
		dto.IconType,
		dto.Address,
		dto.Description,
		dto.AccessRestriction,
		dto.AccessRestrictionDescription,
		dto.Cost,
		dto.CostDescription,
		dto.Hours,
		dto.Open247,
		dto.IsOpenOrActive,
		dto.PhoneNumber,
		model,
	)

	return p
}

func NewPlaceListDTO(dto []PlaceDTO) []placeDomain.Place {
	places := make([]placeDomain.Place, 0, len(dto))

	for i := range dto {
		places = append(places, NewPlaceFromDTO(dto[i]))
	}

	return places
}
