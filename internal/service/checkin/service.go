package checkin

import (
	"context"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	checkinDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/checkin"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
)

type CheckinService interface {
	CreateCheckin(context.Context, checkinDomain.Checkin) error
	MoveFinishedCheckinsToReviews(context.Context) error
	GetValidCheckinForStation(context.Context) ([]checkinDomain.Checkin, error)
}

type CheckinStorage interface {
	CreateCheckin(context.Context, checkinDomain.Checkin) error
	GetFinishedCheckins(context.Context) ([]checkinDomain.Checkin, error)
	DeleteCheckinByCheckinID(context.Context, checkinDomain.CheckinID) error
	GetValidCheckinForStation(context.Context) ([]checkinDomain.Checkin, error)
}

type ReviewStorage interface {
	CreateReview(context.Context, reviewDomain.Review) error
	GetReviewWithPositiveRating(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
	GetReviewsWithNotNullRating(context.Context, placeDomain.PlaceID) ([]reviewDomain.Review, error)
}

type OutletStorage interface {
	GetOutletByID(context.Context, outletDomain.OutletID) (outletDomain.Outlet, error)
	GetOutletsByStationID(context.Context, stationDomain.StationID) ([]outletDomain.Outlet, error)
}

type StationStorage interface {
	GetPlaceIdByStationID(context.Context, stationDomain.StationID) (placeDomain.PlaceID, error)
}

type PlaceStorage interface {
	UpdatePlaceScoreByID(context.Context, placeDomain.PlaceID, float32) error
}

type service struct {
	checkinStorage CheckinStorage
	reviewStorage  ReviewStorage
	outletStorage  OutletStorage
	stationStorage StationStorage
	placeStorage   PlaceStorage
}

func NewCheckinService(
	checkinStorage CheckinStorage,
	reviewStorage ReviewStorage,
	outletStorage OutletStorage,
	stationStorage StationStorage,
	placeStorage PlaceStorage,
) CheckinService {
	return &service{
		checkinStorage: checkinStorage,
		reviewStorage:  reviewStorage,
		outletStorage:  outletStorage,
		stationStorage: stationStorage,
		placeStorage:   placeStorage,
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

		outlet, err := s.outletStorage.GetOutletByID(ctx, checkin.GetOutletId())
		if err != nil {
			return err
		}

		userId := checkin.GetUserId()
		rating := checkin.GetRating()
		connectorType := outlet.GetConnector()
		//kilowatts := checkin.GetKilowatts()
		userName := checkin.GetUserName()

		review := reviewDomain.NewReview(
			checkin.GetStationId(),
			checkin.GetOutletId(),
			&userId,
			checkin.GetComment(),
			&rating,
			&connectorType,
			checkin.GetKilowatts(),
			&userName,
			nil,
			checkin.GetVehicleType(),
			domain.NewModelFrom(checkin.GetCreatedAt(), nil),
		)

		err = s.reviewStorage.CreateReview(ctx, review)
		if err != nil {
			return err
		}

		placeId, err := s.stationStorage.GetPlaceIdByStationID(ctx, review.GetStationID())
		if err != nil {
			return err
		}

		positiveRatings, err := s.reviewStorage.GetReviewWithPositiveRating(ctx, placeId)
		if err != nil {
			return err
		}

		allRatings, err := s.reviewStorage.GetReviewsWithNotNullRating(ctx, placeId)
		if err != nil {
			return err
		}

		newScore := (float32(len(positiveRatings)) / float32(len(allRatings))) * 10

		return s.placeStorage.UpdatePlaceScoreByID(ctx, placeId, newScore)
	}

	return nil
}

func (s *service) GetValidCheckinForStation(ctx context.Context) ([]checkinDomain.Checkin, error) {
	return s.checkinStorage.GetValidCheckinForStation(ctx)
}
