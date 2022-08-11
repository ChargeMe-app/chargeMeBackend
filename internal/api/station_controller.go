package api

import (
	"fmt"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	"github.com/poorfrombabylon/chargeMeBackend/libhttp"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	"net/http"
)

// Получение списка локаций с зарядками в пределах координат
// (GET /v1/locations)
func (api *apiServer) GetLocations(w http.ResponseWriter, r *http.Request, params schema.GetLocationsParams) {
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

		var stationResponse []schema.StationPreliminary

		for _, station := range stations {
			outlets, err := api.outletService.GetOutletsByStationID(ctx, station.GetStationID())
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			var outletResponse []schema.OutletPreliminary

			for _, outlet := range outlets {
				outletSchema := schema.OutletPreliminary{
					Connector: outlet.GetConnector(),
					Kilowatts: outlet.GetKilowatts(),
					Id:        outlet.GetOutletID().String(),
					Power:     outlet.GetPower(),
				}

				outletResponse = append(outletResponse, outletSchema)
			}

			stationSchema := schema.StationPreliminary{
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
			IconType:  place.GetPlaceIconType(),
			Stations:  stationResponse,
		}

		addresses = append(addresses, addressSchema)
	}

	response := schema.ResponseLocations{
		Locations: addresses,
	}

	libhttp.SendJSON(ctx, w, response)
}

// Получение станций в локации
// (GET /v1/locations/stations)
func (api *apiServer) GetChargingStations(w http.ResponseWriter, r *http.Request, params schema.GetChargingStationsParams) {
	ctx := r.Context()
	placeID := placeDomain.PlaceID(params.LocationId)
	var reviewsResponse []schema.Review
	var stationsResponse []schema.StationFull
	var amenitiesResponse []schema.Amenity

	reviews, err := api.reviewService.GetReviewsListByLocationID(ctx, placeID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, review := range reviews {
		r := schema.Review{
			Id:        review.GetReviewID().String(),
			StationId: review.GetStationID().String(),
			OutletId:  review.GetOutletID().String(),
			Comment:   review.GetComment(),
			Rating:    review.GetRating(),
		}

		reviewsResponse = append(reviewsResponse, r)
	}

	stations, err := api.stationService.GetStationsByPlaceID(ctx, placeID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, station := range stations {
		outlets, err := api.outletService.GetOutletsByStationID(ctx, station.GetStationID())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var outletResponse []schema.OutletPreliminary

		for _, outlet := range outlets {
			outletSchema := schema.OutletPreliminary{
				Connector: outlet.GetConnector(),
				Kilowatts: outlet.GetKilowatts(),
				Id:        outlet.GetOutletID().String(),
				Power:     outlet.GetPower(),
			}

			outletResponse = append(outletResponse, outletSchema)
		}

		s := schema.StationFull{
			Id:              station.GetStationID().String(),
			Outlets:         outletResponse,
			Available:       station.GetStationAvailability(),
			Cost:            station.GetStationCost(),
			Name:            station.GetStationName(),
			Manufacturer:    station.GetStationManufacturer(),
			CostDescription: station.GetStationCostDescription(),
			Hours:           station.GetStationWorkingHours(),
			Kilowatts:       station.GetStationKilowatts(),
		}

		stationsResponse = append(stationsResponse, s)
	}

	amenities, err := api.amenityService.GetAmenitiesListByLocationID(ctx, placeID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, amenity := range amenities {
		a := schema.Amenity{
			Id:         amenity.GetAmenityID().String(),
			LocationId: amenity.GetLocationID().String(),
			Form:       amenity.GetAmenityForm(),
		}

		amenitiesResponse = append(amenitiesResponse, a)
	}

	response := schema.ResponseStations{
		Stations:  stationsResponse,
		Reviews:   reviewsResponse,
		Amenities: amenitiesResponse,
	}

	libhttp.SendJSON(ctx, w, response)
}
