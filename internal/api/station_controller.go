package api

import (
	"chargeMe/specs/schema"
	"fmt"
	"log"
	"net/http"
)

// Проверка сервиса
// (GET /healthz)
func (api *apiServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	w.Write([]byte("hello healthCheck"))
	fmt.Println("hello healthCheck")
	log.Fatal("kek HealthCheck")
}

// Получение списка зарядных станций в пределах координат
// (GET /v1/stations)
func (api *apiServer) GetChargingStations(w http.ResponseWriter, r *http.Request, params schema.GetChargingStationsParams) {
	//ctx := r.Context()

	fmt.Println(int(*params.LongitudeMin))

	w.Write([]byte("hello GetChargingStations"))
}
