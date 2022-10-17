package api

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/service"
	amenityService "github.com/poorfrombabylon/chargeMeBackend/internal/service/amenity"
	checkinService "github.com/poorfrombabylon/chargeMeBackend/internal/service/checkin"
	outletService "github.com/poorfrombabylon/chargeMeBackend/internal/service/outlet"
	placeService "github.com/poorfrombabylon/chargeMeBackend/internal/service/place"
	reviewService "github.com/poorfrombabylon/chargeMeBackend/internal/service/review"
	stationService "github.com/poorfrombabylon/chargeMeBackend/internal/service/station"
	userService "github.com/poorfrombabylon/chargeMeBackend/internal/service/user"
	"log"
	"net/http"

	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
)

var _ schema.ServerInterface = &apiServer{}

type apiServer struct {
	placeService   placeService.PlaceService
	stationService stationService.StationService
	outletService  outletService.OutletService
	reviewService  reviewService.ReviewService
	checkinService checkinService.CheckinService
	amenityService amenityService.AmenityService
	userService    userService.UserService
}

func NewApiServer(serviceRegistry *service.Services) schema.ServerInterface {
	return &apiServer{
		serviceRegistry.Place,
		serviceRegistry.Station,
		serviceRegistry.Outlet,
		serviceRegistry.Review,
		serviceRegistry.Checkin,
		serviceRegistry.Amenity,
		serviceRegistry.User,
	}
}

// Проверка сервиса
// (GET /healthz)
func (api *apiServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	w.Write([]byte("hello healthCheck"))
	log.Println("hello healthCheck")
}
