package api

import (
	"fmt"
	"github.com/poorfrombabylon/chargeMeBackend/libhttp"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	"net/http"
)

// Получение списка зарядных станций в пределах координат
// (GET /v1/stations)
func (api *apiServer) GetChargingStations(w http.ResponseWriter, r *http.Request, params schema.GetChargingStationsParams) {
	ctx := r.Context()
	var addresses []schema.AddressStations

	minLongitude := *params.LongitudeMin
	maxLongitude := *params.LongitudeMax

	minLatitude := *params.LatitudeMin
	maxLatitude := *params.LatitudeMax

	places, err := api.placeService.GetPlaces(ctx, minLongitude, maxLongitude, minLatitude, maxLatitude)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, place := range places {
		stations, err := api.stationService.GetStationsByPlaceID(ctx, place.GetPlaceID())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var stationResponse []schema.Station

		for _, station := range stations {
			outlets, err := api.outletService.GetOutletsByStationID(ctx, station.GetStationID())
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			var outletResponse []schema.Outlet

			for _, outlet := range outlets {
				outletSchema := schema.Outlet{
					Connector: outlet.GetConnector(),
					Kilowatts: outlet.GetKilowatts(),
					Id:        outlet.GetOutletID().String(),
					Power:     outlet.GetPower(),
				}

				outletResponse = append(outletResponse, outletSchema)
			}

			stationSchema := schema.Station{
				Id:      station.GetStationID().String(),
				Outlets: outletResponse,
			}

			stationResponse = append(stationResponse, stationSchema)
		}

		addressSchema := schema.AddressStations{
			Access:    *place.GetPlaceAccess(),
			Address:   *place.GetPlaceAddress(),
			Id:        place.GetPlaceID().String(),
			Latitude:  place.GetPlaceLatitude(),
			Longitude: place.GetPlaceLongitude(),
			Name:      place.GetPlaceName(),
			Score:     place.GetPlaceScore(),
			Stations:  stationResponse,
		}

		addresses = append(addresses, addressSchema)
	}

	response := schema.ResponseLocations{
		Data: addresses,
	}

	libhttp.SendJSON(ctx, w, response)
}
