package api

import (
	"github.com/ignishub/terr/transport/httperror"
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	amenityDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/amenity"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	"github.com/poorfrombabylon/chargeMeBackend/libhttp"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	chargeMeV1 "github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	"log"
	"net/http"
)

// Получение списка локаций с зарядками в пределах координат
// (GET /v1/locations)
func (api *apiServer) GetLocations(w http.ResponseWriter, r *http.Request, params schema.GetLocationsParams) {
	log.Println("api.station.GetStations")
	ctx := r.Context()
	var addresses []schema.AddressStationsPreliminary

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

		addressSchema := schema.AddressStationsPreliminary{
			Access:    *place.GetPlaceAccess(),
			Address:   place.GetPlaceAddress(),
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
func (api *apiServer) GetChargingStationsByLocationID(
	w http.ResponseWriter,
	r *http.Request,
	params schema.GetChargingStationsByLocationIDParams,
) {
	log.Println("api.station.GetChargingStationsByLocationID")
	ctx := r.Context()
	placeID := placeDomain.PlaceID(params.LocationId)
	var reviewsResponse []schema.Review
	var stationsResponse []schema.StationFull
	var amenitiesResponse []schema.Amenity

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

	for _, review := range reviews {
		r := schema.Review{
			Id:            review.GetReviewID().String(),
			StationId:     review.GetStationID().String(),
			OutletId:      review.GetOutletID().String(),
			Comment:       review.GetComment(),
			Rating:        review.GetRating(),
			ConnectorType: review.GetConnectorType(),
			UserName:      review.GetUserName(),
			VehicleName:   review.GetVehicleName(),
			VehicleType:   nil,
			CreatedAt:     review.GetCreatedAt(),
		}

		reviewsResponse = append(reviewsResponse, r)
	}

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

		checkins, err := api.checkinService.GetValidCheckinForStation(ctx)
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}

		if checkins != nil && len(checkins) != 0 {
			c := transformCheckin(checkins[0])
			s.Checkin = &c
		}

		stationsResponse = append(stationsResponse, s)
	}

	amenities, err := api.amenityService.GetAmenitiesListByLocationID(ctx, placeID)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
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

	response := schema.AddressStationsFull{
		Access:                       place.GetPlaceAccess(),
		Address:                      place.GetPlaceAddress(),
		IconType:                     place.GetPlaceIconType(),
		Id:                           place.GetPlaceID().String(),
		Latitude:                     place.GetPlaceLatitude(),
		Longitude:                    place.GetPlaceLongitude(),
		Name:                         place.GetPlaceName(),
		Score:                        place.GetPlaceScore(),
		Description:                  place.GetDescription(),
		AccessRestriction:            place.GetAccessRestriction(),
		AccessRestrictionDescription: place.GetAccessRestrictionDescription(),
		Cost:                         place.GetCost(),
		CostDescription:              place.GetCostDescription(),
		Hours:                        place.GetHours(),
		Open247:                      place.GetOpen247(),
		ComingSoon:                   place.IsComingSoon(),
		PhoneNumber:                  place.GetPhoneNumber(),
		Stations:                     stationsResponse,
		Reviews:                      &reviewsResponse,
		Amenities:                    &amenitiesResponse,
	}

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

	location := placeDomain.NewPlace(
		req.Name,
		req.Score,
		req.Longitude,
		req.Latitude,
		req.Access,
		req.IconType,
		req.Address,
		req.Description,
		req.AccessRestriction,
		req.AccessRestrictionDescription,
		req.Cost,
		req.CostDescription,
		req.Hours,
		req.Open247,
		req.ComingSoon,
		req.PhoneNumber,
		domain.NewModel(),
	)

	log.Println()
	log.Println(location)
	log.Println()

	err = api.placeService.CreatePlace(ctx, location)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	for _, a := range *req.Amenities {
		amenity := amenityDomain.NewAmenity(
			location.GetPlaceID(),
			a.Form,
			domain.NewModel(),
		)

		err = api.amenityService.CreateAmenity(ctx, amenity)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err.Error())
			return
		}
	}

	for _, s := range req.Stations {
		station := stationDomain.NewStation(
			location.GetPlaceID(),
			s.Available,
			s.Cost,
			s.Name,
			s.Manufacturer,
			s.CostDescription,
			s.Hours,
			s.Kilowatts,
			domain.NewModel(),
		)

		err = api.stationService.CreateStation(ctx, station)
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err.Error())
			return
		}

		for _, o := range s.Outlets {
			outlet := outletDomain.NewOutlet(
				station.GetStationID(),
				o.Connector,
				o.Kilowatts,
				o.Power,
				domain.NewModel(),
			)

			err = api.outletService.CreateOutlet(ctx, outlet)
			if err != nil {
				w.Write([]byte(err.Error()))
				log.Println(err.Error())
				return
			}
		}
	}
}
