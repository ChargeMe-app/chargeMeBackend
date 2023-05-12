package checkin

import (
	"context"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	checkinDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/checkin"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration/common_utils"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration/my_ecars"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration/sitronics"
	"log"
	"time"
)

type CheckinService interface {
	CreateCheckin(context.Context, checkinDomain.Checkin) error
	MoveFinishedCheckinsToReviews(context.Context) error
	GetValidCheckinForStation(context.Context, stationDomain.Station, *string) ([]checkinDomain.Checkin, error)
}

type CheckinStorage interface {
	CreateCheckin(context.Context, checkinDomain.Checkin) error
	GetFinishedCheckins(context.Context) ([]checkinDomain.Checkin, error)
	DeleteCheckinByID(context.Context, checkinDomain.CheckinID) error
	GetValidCheckinForStation(context.Context, stationDomain.StationID) ([]checkinDomain.Checkin, error)
	DeleteCheckinByOutletID(context.Context, outletDomain.OutletID) error
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

type MyECarsIntegration interface {
	GetStationByID(context.Context, string) (my_ecars.MyECarsStationsResponse, error)
}

type SitronicsIntegration interface {
	GetPlaceByName(context.Context, string) (sitronics.SitronicsStationInfo, error)
}

type UserStorage interface {
	GetUserByIdentifier(context.Context, string) (userDomain.User, error)
}

type service struct {
	checkinStorage       CheckinStorage
	reviewStorage        ReviewStorage
	outletStorage        OutletStorage
	stationStorage       StationStorage
	placeStorage         PlaceStorage
	userStorage          UserStorage
	myECarsIntegration   MyECarsIntegration
	sitronicsIntegration SitronicsIntegration
}

func NewCheckinService(
	checkinStorage CheckinStorage,
	reviewStorage ReviewStorage,
	outletStorage OutletStorage,
	stationStorage StationStorage,
	placeStorage PlaceStorage,
	userStorage UserStorage,
	myECarsIntegration MyECarsIntegration,
	sitronicsIntegration SitronicsIntegration,
) CheckinService {
	return &service{
		checkinStorage:       checkinStorage,
		reviewStorage:        reviewStorage,
		outletStorage:        outletStorage,
		stationStorage:       stationStorage,
		placeStorage:         placeStorage,
		userStorage:          userStorage,
		myECarsIntegration:   myECarsIntegration,
		sitronicsIntegration: sitronicsIntegration,
	}
}

func (s *service) CreateCheckin(ctx context.Context, checkin checkinDomain.Checkin) error {
	return s.checkinStorage.CreateCheckin(ctx, checkin)
}

func (s *service) MoveFinishedCheckinsToReviews(ctx context.Context) error {
	log.Println("checkin.service.MoveFinishedCheckinsToReviews")
	finishedCheckinList, err := s.checkinStorage.GetFinishedCheckins(ctx)
	if err != nil {
		return err
	}

	for _, checkin := range finishedCheckinList {
		err = s.checkinStorage.DeleteCheckinByID(ctx, checkin.GetCheckinId())
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

		if len(allRatings) >= 5 {
			newScore := (float32(len(positiveRatings)) / float32(len(allRatings))) * 10

			return s.placeStorage.UpdatePlaceScoreByID(ctx, placeId, newScore)
		}
	}

	return nil
}

func (s *service) GetValidCheckinForStation(
	ctx context.Context,
	station stationDomain.Station,
	companyName *string,
) ([]checkinDomain.Checkin, error) {
	baseCheckins, err := s.checkinStorage.GetValidCheckinForStation(ctx, station.GetStationID())
	if err != nil {
		return nil, err
	}

	if companyName != nil {
		stationsInfo, err := s.getStationAvailability(ctx, station, *companyName)
		if err != nil {
			return nil, err
		}

		for _, c := range stationsInfo.Connectors {
			isExist := s.isCheckinExist(baseCheckins, stationsInfo.StationID, c.OutletID)

			if c.Available == 1 {
				if isExist {
					err = s.checkinStorage.DeleteCheckinByOutletID(ctx, c.OutletID)
					if err != nil {
						return nil, err
					}
				}
			} else {
				if !isExist {
					user, err := s.userStorage.GetUserByIdentifier(ctx, *companyName)
					if err != nil {
						return nil, err
					}

					integrationCheckin := checkinDomain.NewCheckin(
						user.GetUserId(),
						station.GetStationID(),
						c.OutletID,
						*user.GetDisplayName(),
						0,
						nil,
						nil,
						nil,
						0,
						true,
						time.Now(),
					)

					baseCheckins = append(baseCheckins, integrationCheckin)
				}
			}
		}
	}

	return baseCheckins, nil
}

func (s *service) getStationAvailability(
	ctx context.Context,
	station stationDomain.Station,
	companyName string,
) (common_utils.Station, error) {
	switch companyName {
	case placeDomain.Sitronics:
		return s.getSitronicsStationAvailability(ctx, station)
	case placeDomain.MyECars:
		return s.getMyECarsStationAvailability(ctx, station)
	}

	return common_utils.Station{}, nil
}

func (s *service) getSitronicsStationAvailability(
	ctx context.Context,
	station stationDomain.Station,
) (common_utils.Station, error) {
	sitronicsStation, err := s.sitronicsIntegration.GetPlaceByName(ctx, *station.GetStationName())
	if err != nil {
		return common_utils.Station{}, err
	}

	result := common_utils.ConvertSitronicsStationToCommonStation(sitronicsStation.CPCard)

	return result, nil
}

func (s *service) getMyECarsStationAvailability(
	ctx context.Context,
	station stationDomain.Station,
) (common_utils.Station, error) {
	myECarsStation, err := s.myECarsIntegration.GetStationByID(ctx, station.GetStationID().String())
	if err != nil {
		return common_utils.Station{}, err
	}

	result, err := common_utils.ConvertMyECarsStationToCommonStation(myECarsStation.Evse[0])
	if err != nil {
		return common_utils.Station{}, err
	}

	return result, nil
}

func (s *service) isCheckinExist(
	checkins []checkinDomain.Checkin,
	stationID stationDomain.StationID,
	outletID outletDomain.OutletID,
) bool {
	result := false

	for _, c := range checkins {
		if c.GetStationId() == stationID && c.GetOutletId() == outletID {
			result = true
		}
	}

	return result
}
