package api

import (
	"github.com/google/uuid"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	checkinDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/checkin"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"github.com/poorfrombabylon/chargeMeBackend/libhttp"
	chargeMeV1 "github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	"log"
	"net/http"
	"time"
)

func (api *apiServer) CreateReview(w http.ResponseWriter, r *http.Request) {
	log.Println("api.review.CreateReview")
	ctx := r.Context()

	var req chargeMeV1.CreateReviewJSONRequestBody

	err := libhttp.ReceiveJSON(ctx, r, &req)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	userID := userDomain.UserID(uuid.MustParse(req.UserId))

	review := reviewDomain.NewReview(
		stationDomain.StationID(req.StationId),
		outletDomain.OutletID(req.OutletId),
		&userID,
		req.Comment,
		req.Rating,
		req.ConnectorType,
		req.Kilowatts,
		&req.UserName,
		req.VehicleName,
		req.VehicleType,
		domain.NewModel(),
	)

	err = api.reviewService.CreateReview(ctx, review)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
}

func (api *apiServer) CreateCheckin(w http.ResponseWriter, r *http.Request) {
	log.Println("api.review.CreateCheckin")
	ctx := r.Context()

	var req chargeMeV1.CreateCheckinJSONRequestBody

	err := libhttp.ReceiveJSON(ctx, r, &req)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	checkin := checkinDomain.NewCheckin(
		userDomain.UserID(uuid.MustParse(req.UserId)),
		stationDomain.StationID(req.StationId),
		outletDomain.OutletID(req.OutletId),
		req.UserName,
		req.Duration,
		req.VehicleType,
		req.Comment,
		req.Kilowatts,
		req.Rating,
		time.Now().Add(time.Duration(req.Duration)*time.Minute),
	)

	err = api.checkinService.CreateCheckin(ctx, checkin)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
}
