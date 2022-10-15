package api

import (
	checkinDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/checkin"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
)

func convertGoogleCredentials(
	userId userDomain.UserID,
	credentials *schema.GoogleAuthCredentials,
) userDomain.GoogleCredentials {
	return userDomain.NewGoogleCredentials(
		userId,
		credentials.IdToken,
		credentials.AccessToken,
	)
}

func convertAppleCredentials(
	userId userDomain.UserID,
	credentials *schema.AppleAuthCredentials,
) userDomain.AppleCredentials {
	return userDomain.NewAppleCredentials(
		userId,
		credentials.AuthorizationCode,
		credentials.IdentityToken,
	)
}

func transformUserVehicles(vehicles []userDomain.Vehicle) []schema.Vehicle {
	if vehicles == nil {
		return nil
	}

	var response []schema.Vehicle

	for _, i := range vehicles {
		v := schema.Vehicle{
			VehicleType: i.GetVehicleType(),
		}

		response = append(response, v)
	}

	return response
}

func transformReviewsNumber(reviews []reviewDomain.Review) *int {
	num := len(reviews)
	return &num
}

func transformCheckin(checkin checkinDomain.Checkin) schema.CheckinStation {
	return schema.CheckinStation{
		UserId:      checkin.GetUserId().String(),
		StationId:   checkin.GetStationId().String(),
		OutletId:    checkin.GetOutletId().String(),
		VehicleType: checkin.GetVehicleType(),
		UserName:    checkin.GetUserName(),
		Comment:     checkin.GetComment(),
		Kilowatts:   checkin.GetKilowatts(),
		Rating:      checkin.GetRating(),
		FinishesAt:  checkin.GetFinishedAt(),
	}
}
