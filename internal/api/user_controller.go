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
		userDomain.UserId(uuid.MustParse(req.UserId)),
		req.VehicleType,
	)

	err = api.userService.CreateVehicle(ctx, vehicle)
	if err != nil {
		httperror.ServeError(w, err)
		return
	}

}
