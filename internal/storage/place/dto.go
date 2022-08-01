package place

import placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"

type PlaceDTO struct {
	PlaceID   string   `db:"id"`
	Name      string   `db:"name"`
	Score     *float32 `db:"score"`
	Longitude float32  `db:"longitude"`
	Latitude  float32  `db:"latitude"`
	Address   *string  `db:"address"`
	Access    *int     `db:"access"`
	IconLink  *string  `db:"icon_link"`
}

func NewPlaceFromDTO(dto PlaceDTO) placeDomain.Place {
	return placeDomain.NewPlaceWithID(
		placeDomain.PlaceID(dto.PlaceID),
		dto.Name,
		dto.Score,
		dto.Longitude,
		dto.Latitude,
		dto.Access,
		dto.IconLink,
		dto.Address,
	)
}

func NewPlaceListDTO(dto []PlaceDTO) []placeDomain.Place {
	places := make([]placeDomain.Place, 0, len(dto))

	for i := range dto {
		places = append(places, NewPlaceFromDTO(dto[i]))
	}

	return places
}
