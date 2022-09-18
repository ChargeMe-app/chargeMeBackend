package api

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ignishub/terr/transport/httperror"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"github.com/poorfrombabylon/chargeMeBackend/libhttp"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	"net/http"
)

func (api *apiServer) AddVehicle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api.user.AddVehicle")
	ctx := r.Context()

	var req schema.AddVehicleJSONRequestBody

	err := libhttp.ReceiveJSON(ctx, r, &req)
	if err != nil {
		httperror.ServeError(w, err)
		return
	}

	vehicle := userDomain.NewVehicle(
		userDomain.UserID(uuid.MustParse(req.UserId)),
		req.VehicleType,
	)

	err = api.userService.CreateVehicle(ctx, vehicle)
	if err != nil {
		httperror.ServeError(w, err)
		return
	}
}

func (api *apiServer) GetUserByUserId(w http.ResponseWriter, r *http.Request, userId string) {
	fmt.Println("api.user.GetUserByID")
	ctx := r.Context()

	user, err := api.userService.GetUserByUserId(ctx, userDomain.UserID(uuid.MustParse(userId)))
	if err != nil {
		httperror.ServeError(w, err)
		return
	}

	userVehicles, err := api.userService.GetVehiclesByUserId(ctx, user.GetUserId())
	if err != nil {
		httperror.ServeError(w, err)
		return
	}

	userReviews, err := api.reviewService.GetReviewsListByUserID(ctx, user.GetUserId())
	if err != nil {
		httperror.ServeError(w, err)
		return
	}

	userIdentifierResponse := user.GetUserId().String()

	userResponse := &schema.User{
		Id:            &userIdentifierResponse,
		DisplayName:   *user.GetDisplayName(),
		SignInService: user.GetSignType(),
		VehicleType:   transformUserVehicles(userVehicles),
		PhotoUrl:      user.GetPhotoUrl(),
		TotalReviews:  transformReviewsNumber(userReviews),
	}

	libhttp.SendJSON(ctx, w, userResponse)
}
