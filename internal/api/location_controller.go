package api

import (
	"log"
	"net/http"

	"github.com/ignishub/terr/transport/httperror"
	"github.com/poorfrombabylon/chargeMeBackend/internal/config"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	"github.com/poorfrombabylon/chargeMeBackend/libhttp"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	chargeMeV1 "github.com/poorfrombabylon/chargeMeBackend/specs/schema"
)

// Получение списка локаций с зарядками в пределах координат
// (GET /v1/locations)
func (api *apiServer) GetLocations(w http.ResponseWriter, r *http.Request, params schema.GetLocationsParams) {
	log.Println("api.station.GetStations")
	ctx := r.Context()
	var addresses []schema.LocationPreliminary

	minLongitude := *params.LongitudeMin
	maxLongitude := *params.LongitudeMax

	minLatitude := *params.LatitudeMin
	maxLatitude := *params.LatitudeMax

	places, err := api.placeService.GetPlacesByCoordinates(ctx, minLongitude, maxLongitude, minLatitude, maxLatitude)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	for _, place := range places {
		stations, err := api.stationService.GetStationsByPlaceID(ctx, place.GetPlaceID())
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err.Error())
			return
		}

		var stationResponse []schema.StationPreliminary

		for _, station := range stations {
			outlets, err := api.outletService.GetOutletsByStationID(ctx, station.GetStationID())
			if err != nil {
				w.Write([]byte(err.Error()))
				log.Println(err.Error())
				return
			}

			outletResponse := transformOutletsList(outlets)

			stationSchema := schema.StationPreliminary{
				Id:      station.GetStationID().String(),
				Outlets: outletResponse,
			}

			stationResponse = append(stationResponse, stationSchema)
		}

		addressSchema := transformLocationPreliminary(place, stationResponse)

		addresses = append(addresses, addressSchema)
	}

	response := schema.LocationsOnMap{
		Locations: addresses,
	}

	libhttp.SendJSON(ctx, w, response)
}

// Получение станций в локации
// (GET /v1/locations/stations)
func (api *apiServer) GetChargingStationsByLocationID(
	w http.ResponseWriter,
	r *http.Request,
	params schema.GetChargingStationsByLocationIDParams,
) {
	log.Println("api.station.GetChargingStationsByLocationID")
	ctx := r.Context()
	placeID := placeDomain.PlaceID(params.LocationId)
	var stationsResponse []schema.StationFull

	place, err := api.placeService.GetFullPlaceByID(ctx, placeID)
	if err != nil {
		httperror.ServeError(w, err)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	reviews, err := api.reviewService.GetReviewsListByLocationID(ctx, placeID)
	if err != nil {
		httperror.ServeError(w, err)
		return
	}

	reviewsResponse := transformReviewsList(reviews)

	stations, err := api.stationService.GetStationsByPlaceID(ctx, placeID)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	for _, station := range stations {
		outlets, err := api.outletService.GetOutletsByStationID(ctx, station.GetStationID())
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err.Error())
			return
		}

		outletResponse := transformOutletsList(outlets)

		stationResponse := transformFullStation(station, outletResponse)

		checkins, err := api.checkinService.GetValidCheckinForStation(ctx, station, place.GetCompanyName())
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}

		if checkins != nil && len(checkins) != 0 {
			c := transformCheckinList(checkins)
			stationResponse.Checkins = &c
		}

		stationsResponse = append(stationsResponse, stationResponse)
	}

	photos, err := api.photoService.GetPhotoListByPlaceID(ctx, placeID)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	yandexLink := config.GetConfig().YandexStorage.BaseUrl

	photosResponse := transformPhotosList(photos, yandexLink)

	amenities, err := api.amenityService.GetAmenitiesListByLocationID(ctx, placeID)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	amenitiesResponse := transformAmenitiesList(amenities)

	response := transformLocationFull(place, &photosResponse, stationsResponse, &reviewsResponse, &amenitiesResponse)

	log.Println(response)

	libhttp.SendJSON(ctx, w, response)
}

// Создание локации со станциями
// (POST /v1/locations)
func (api *apiServer) CreateFullLocation(w http.ResponseWriter, r *http.Request) {
	log.Println("api.CreateFullLocation")
	ctx := r.Context()
	var req chargeMeV1.CreateFullLocationJSONBody

	err := libhttp.ReceiveJSON(ctx, r, &req)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	log.Println(r)

	log.Println(req)

	location := convertLocation(req)

	log.Println()
	log.Println(location)
	log.Println()

	err = api.placeService.CreatePlace(ctx, location)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	if req.Amenities != nil {
		amenities := convertAmenitiesList(location.GetPlaceID(), *req.Amenities)

		err = api.amenityService.CreateAmenitiesList(ctx, amenities)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err.Error())
			return
		}
	}

	for _, s := range req.Stations {
		station := convertStation(location.GetPlaceID(), s)

		err = api.stationService.CreateStation(ctx, station)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err.Error())
			return
		}

		outlets := convertOutletList(station.GetStationID(), s.Outlets)

		err = api.outletService.CreateOutletsList(ctx, outlets)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err.Error())
			return
		}
	}
}

// обновление локации со станциями
// (PUT /v1/locations)
func (api *apiServer) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	log.Println("api.UpdateLocation")
	ctx := r.Context()
	var req chargeMeV1.CreateFullLocationJSONBody

	err := libhttp.ReceiveJSON(ctx, r, &req)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	locationId := placeDomain.PlaceID(req.Id)
	location := convertLocationWithID(req, locationId)

	err = api.placeService.UpdatePlace(ctx, location)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	err = api.stationService.DeleteStationsByPlaceID(ctx, locationId)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	err = api.amenityService.DeleteAmenitiesByLocationID(ctx, locationId)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	for _, s := range req.Stations {
		station := convertStation(locationId, s)

		err = api.stationService.CreateStation(ctx, station)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err.Error())
			return
		}

		outlets := convertOutletList(station.GetStationID(), s.Outlets)

		err = api.outletService.CreateOutletsList(ctx, outlets)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err.Error())
			return
		}
	}

	if req.Amenities != nil {
		amenities := convertAmenitiesList(location.GetPlaceID(), *req.Amenities)

		err = api.amenityService.CreateAmenitiesList(ctx, amenities)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err.Error())
			return
		}
	}
}
