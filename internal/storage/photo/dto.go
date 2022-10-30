package photo

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"time"
)

type PhotoDTO struct {
	PhotoID    string    `db:"id"`
	UserID     string    `db:"user_id"`
	Name       string    `db:"name"`
	LocationID string    `db:"location_id"`
	Caption    *string   `db:"caption"`
	CreatedAt  time.Time `db:"created_at"`
}

func NewPhotoFromDTO(dto PhotoDTO) placeDomain.Photo {
	model := domain.NewModelFrom(dto.CreatedAt, nil)

	return placeDomain.NewModelWithModel(
		placeDomain.PhotoID(dto.PhotoID),
		userDomain.UserID(uuid.MustParse(dto.UserID)),
		placeDomain.PlaceID(dto.LocationID),
		dto.Name,
		dto.Caption,
		model,
	)
}

func NewPhotoListFromDTO(dto []PhotoDTO) []placeDomain.Photo {
	result := make([]placeDomain.Photo, 0, len(dto))

	for _, p := range dto {
		result = append(result, NewPhotoFromDTO(p))
	}

	return result
}
