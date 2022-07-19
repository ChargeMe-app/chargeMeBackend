package api

import (
	"fmt"
	"net/http"

	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
)

// Получение списка зарядных станций в пределах координат
// (GET /v1/stations)
func (api *apiServer) GetChargingStations(w http.ResponseWriter, r *http.Request, params schema.GetChargingStationsParams) {
	//ctx := r.Context()

	fmt.Println(int(*params.LongitudeMin))

	w.Write([]byte("hello GetChargingStations"))
}
