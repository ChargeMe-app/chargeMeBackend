package checkin

import (
	"context"
	checkinDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/checkin"
)

type CheckinService interface {
	CreateCheckin(context.Context, checkinDomain.Checkin) error
	MoveFinishedCheckinsToReviews(context.Context) error
}

type CheckinStorage interface {
	CreateCheckin(context.Context, checkinDomain.Checkin) error
	GetFinishedCheckins(context.Context) ([]checkinDomain.Checkin, error)
	DeleteCheckinByCheckinID(context.Context, checkinDomain.CheckinID) error
}

type service struct {
	checkinStorage CheckinStorage
}

func NewCheckinService(checkinStorage CheckinStorage) CheckinService {
	return &service{
		checkinStorage: checkinStorage,
	}
}

func (s *service) CreateCheckin(ctx context.Context, checkin checkinDomain.Checkin) error {
	return s.checkinStorage.CreateCheckin(ctx, checkin)
}

func (s *service) MoveFinishedCheckinsToReviews(ctx context.Context) error {
	finishedCheckinList, err := s.checkinStorage.GetFinishedCheckins(ctx)
	if err != nil {
		return err
	}

	for _, checkin := range finishedCheckinList {
		err = s.checkinStorage.DeleteCheckinByCheckinID(ctx, checkin.GetCheckinId())
		if err != nil {
			return err
		}
	}

	return nil
}
