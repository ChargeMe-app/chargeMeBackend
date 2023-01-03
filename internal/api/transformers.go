package api

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/domain"
	amenityDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/amenity"
	checkinDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/checkin"
	outletDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/outlet"
	placeDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/place"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
	stationDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/station"
	userDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/user"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
)

func convertGoogleCredentials(
	userId userDomain.UserID,
	credentials *schema.GoogleAuthCredentials,
) userDomain.GoogleCredentials {
	return userDomain.NewGoogleCredentials(
		userId,
		credentials.IdToken,
		credentials.AccessToken,
	)
}

func convertAppleCredentials(
	userId userDomain.UserID,
	credentials *schema.AppleAuthCredentials,
) userDomain.AppleCredentials {
	return userDomain.NewAppleCredentials(
		userId,
		credentials.AuthorizationCode,
		credentials.IdentityToken,
	)
}

func convertStationsList(placeID placeDomain.PlaceID, stations []schema.StationFull) []stationDomain.Station {
	result := []stationDomain.Station{}

	for i := range stations {
		result = append(result, convertStation(placeID, stations[i]))
	}

	return result
}

func convertStation(placeID placeDomain.PlaceID, station schema.StationFull) stationDomain.Station {
	return stationDomain.NewStation(
		placeID,
		station.Available,
		station.Cost,
		station.Name,
		station.Manufacturer,
		station.CostDescription,
		station.Hours,
		station.Kilowatts,
		domain.NewModel(),
	)
}

func convertLocation(reqLocation schema.CreateFullLocationJSONBody) placeDomain.Place {
	return placeDomain.NewPlace(
		reqLocation.Name,
		reqLocation.Score,
		reqLocation.Longitude,
		reqLocation.Latitude,
		reqLocation.Access,
		reqLocation.IconType,
		reqLocation.Address,
		reqLocation.Description,
		reqLocation.AccessRestriction,
		reqLocation.AccessRestrictionDescription,
		reqLocation.Cost,
		reqLocation.CostDescription,
		reqLocation.Hours,
		reqLocation.Open247,
		reqLocation.ComingSoon,
		reqLocation.PhoneNumber,
		domain.NewModel(),
	)
}

func convertLocationWithID(reqLocation schema.CreateFullLocationJSONBody, locationID placeDomain.PlaceID) placeDomain.Place {
	return placeDomain.NewPlaceWithID(
		locationID,
		reqLocation.Name,
		reqLocation.Score,
		reqLocation.Longitude,
		reqLocation.Latitude,
		reqLocation.Access,
		reqLocation.IconType,
		reqLocation.Address,
		reqLocation.Description,
		reqLocation.AccessRestriction,
		reqLocation.AccessRestrictionDescription,
		reqLocation.Cost,
		reqLocation.CostDescription,
		reqLocation.Hours,
		reqLocation.Open247,
		reqLocation.ComingSoon,
		reqLocation.PhoneNumber,
		domain.NewModel(),
	)
}

func convertOutlet(stationID stationDomain.StationID, outlet schema.OutletPreliminary) outletDomain.Outlet {
	return outletDomain.NewOutlet(
		stationID,
		outlet.Connector,
		outlet.Kilowatts,
		outlet.Power,
		domain.NewModel(),
	)
}

func convertOutletList(stationID stationDomain.StationID, outlets []schema.OutletPreliminary) []outletDomain.Outlet {
	result := make([]outletDomain.Outlet, 0, len(outlets))

	for i := range outlets {
		result = append(result, convertOutlet(stationID, outlets[i]))
	}

	return result
}

func transformUserVehicles(vehicles []userDomain.Vehicle) []schema.Vehicle {
	if vehicles == nil {
		return nil
	}

	var response []schema.Vehicle

	for _, i := range vehicles {
		v := schema.Vehicle{
			VehicleType: i.GetVehicleType(),
		}

		response = append(response, v)
	}

	return response
}

func transformReviewsNumber(reviews []reviewDomain.Review) *int {
	num := len(reviews)
	return &num
}

func transformCheckin(checkin checkinDomain.Checkin) schema.CheckinStation {
	return schema.CheckinStation{
		UserId:      checkin.GetUserId().String(),
		StationId:   checkin.GetStationId().String(),
		OutletId:    checkin.GetOutletId().String(),
		VehicleType: checkin.GetVehicleType(),
		UserName:    checkin.GetUserName(),
		Comment:     checkin.GetComment(),
		Kilowatts:   checkin.GetKilowatts(),
		Rating:      checkin.GetRating(),
		FinishesAt:  checkin.GetFinishedAt(),
	}
}

func convertAmenity(placeID placeDomain.PlaceID, req schema.Amenity) amenityDomain.Amenity {
	return amenityDomain.NewAmenity(
		placeID,
		req.Form,
		domain.NewModel(),
	)
}

func convertAmenitiesList(placeID placeDomain.PlaceID, req []schema.Amenity) []amenityDomain.Amenity {
	result := []amenityDomain.Amenity{}

	for i := range req {
		result = append(result, convertAmenity(placeID, req[i]))
	}

	return result
}
