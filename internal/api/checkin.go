package api

import (
	"fmt"
	"github.com/google/uuid"
	checkinDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/checkin"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"github.com/poorfrombabylon/chargeMeBackend/libhttp"
	chargeMeV1 "github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	"net/http"
)

func (api *apiServer) CreateCheckin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api.checkin.CreateCheckin")
	ctx := r.Context()

	var req chargeMeV1.CreateCheckinJSONRequestBody

	err := libhttp.ReceiveJSON(ctx, r, &req)
	if err != nil {
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return
	}

	checkin := checkinDomain.NewCheckin(
		userDomain.UserID(uuid.MustParse(req.UserId)),
		stationDomain.StationID(req.StationId),
		outletDomain.OutletID(req.OutletId),
		req.Duration,
		req.VehicleType,
		req.Comment,
		req.Kilowatts,
		req.Rating,
	)

	err = api.checkinService.CreateCheckin(ctx, checkin)
	if err != nil {
		w.Write([]byte(err.Error()))
		fmt.Println(err.Error())
		return
	}
}
