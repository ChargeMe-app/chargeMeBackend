package api

import (
	"github.com/ignishub/terr/transport/httperror"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"github.com/poorfrombabylon/chargeMeBackend/libhttp"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	"log"
	"net/http"
)

func (api *apiServer) Authenticate(w http.ResponseWriter, r *http.Request) {
	log.Println("api.auth.Authenticate")
	ctx := r.Context()
	var req schema.AuthenticateJSONRequestBody

	err := libhttp.ReceiveJSON(ctx, r, &req)
	if err != nil {
		httperror.ServeError(w, err)
		return
	}

	user := userDomain.NewUser(
		req.DisplayName,
		req.Email,
		req.UserIdentifier,
		req.PhotoUrl,
		req.SignType,
	)

	u, err := api.userService.GetOrCreateUser(ctx, user)
	if err != nil {
		httperror.ServeError(w, err)
		return
	}

	if u != nil {
		//userVehicles, err := api.userService.GetVehiclesByUserId(ctx, u.GetUserId())
		//if err != nil {
		//	httperror.ServeError(w, err)
		//	return
		//}
		//
		//userReviews, err := api.reviewService.GetReviewsListByUserID(ctx, u.GetUserId())
		//if err != nil {
		//	httperror.ServeError(w, err)
		//	return
		//}

		//userResponse := &schema.User{
		//	DisplayName:   *u.GetDisplayName(),
		//	SignInService: u.GetSignType(),
		//	VehicleType:   transformUserVehicles(userVehicles),
		//	PhotoUrl:      u.GetPhotoUrl(),
		//	TotalReviews:  transformReviewsNumber(userReviews),
		//}

		libhttp.SendJSON(ctx, w, schema.UserId{UserId: u.GetUserId().String()})
	} else {
		if req.GoogleCredentials != nil {
			credentials := convertGoogleCredentials(user.GetUserId(), req.GoogleCredentials)

			err = api.userService.CreateGoogleCredentials(ctx, credentials)
			if err != nil {
				httperror.ServeError(w, err)
				return
			}
		}

		if req.AppleCredentials != nil {
			credentials := convertAppleCredentials(user.GetUserId(), req.AppleCredentials)

			err = api.userService.CreateAppleCredentials(ctx, credentials)
			if err != nil {
				httperror.ServeError(w, err)
				return
			}
		}

		libhttp.SendJSON(ctx, w, schema.UserId{UserId: user.GetUserId().String()})
	}
}
