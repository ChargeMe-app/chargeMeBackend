package api

import (
	"fmt"
	"github.com/poorfrombabylon/chargeMeBackend/internal/service"
	outletService "github.com/poorfrombabylon/chargeMeBackend/internal/service/outlet"
	placeService "github.com/poorfrombabylon/chargeMeBackend/internal/service/place"
	stationService "github.com/poorfrombabylon/chargeMeBackend/internal/service/station"
	"net/http"

	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
)

var _ schema.ServerInterface = &apiServer{}

type apiServer struct {
	placeService   placeService.PlaceService
	stationService stationService.StationService
	outletService  outletService.OutletService
}

func NewApiServer(serviceRegistry *service.Services) schema.ServerInterface {
	return &apiServer{
		serviceRegistry.Place,
		serviceRegistry.Station,
		serviceRegistry.Outlet,
	}
}

// Проверка сервиса
// (GET /healthz)
func (api *apiServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	w.Write([]byte("hello healthCheck"))
	fmt.Println("hello healthCheck")
}
