package checkin

import (
	"context"
	checkinDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/checkin"
)

type CheckinService interface {
	CreateCheckin(context.Context, checkinDomain.Checkin) error
}

type CheckinStorage interface {
	CreateCheckin(context.Context, checkinDomain.Checkin) error
}

type service struct {
	checkinStorage CheckinService
}

func NewCheckinService(checkinStorage CheckinStorage) CheckinService {
	return &service{
		checkinStorage: checkinStorage,
	}
}

func (s *service) CreateCheckin(ctx context.Context, checkin checkinDomain.Checkin) error {
	return s.checkinStorage.CreateCheckin(ctx, checkin)
}
