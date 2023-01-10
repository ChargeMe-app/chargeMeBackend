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
		reqLocation.Cost,
		reqLocation.CostDescription,
		reqLocation.Hours,
		reqLocation.Open247,
		reqLocation.ComingSoon,
		reqLocation.PhoneNumber,
		domain.NewModel(),
	)
}

func convertOutlet(stationID stationDomain.StationID, outlet schema.Outlet) outletDomain.Outlet {
	return outletDomain.NewOutlet(
		stationID,
		outlet.Connector,
		outlet.Kilowatts,
		outlet.Power,
		domain.NewModel(),
	)
}

func convertOutletList(stationID stationDomain.StationID, outlets []schema.Outlet) []outletDomain.Outlet {
	result := make([]outletDomain.Outlet, 0, len(outlets))

	for i := range outlets {
		result = append(result, convertOutlet(stationID, outlets[i]))
	}

	return result
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

func transformOutlet(outlet outletDomain.Outlet) schema.Outlet {
	return schema.Outlet{
		Connector: outlet.GetConnector(),
		Kilowatts: outlet.GetKilowatts(),
		Id:        outlet.GetOutletID().String(),
		Power:     outlet.GetPower(),
		Price:     outlet.GetPrice(),
		PriceUnit: outlet.GetPriceUnit(),
	}
}

func transformOutletsList(outlets []outletDomain.Outlet) []schema.Outlet {
	result := make([]schema.Outlet, 0, len(outlets))

	for i := range outlets {
		result = append(result, transformOutlet(outlets[i]))
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
		IsAuto:      checkin.IsAutoCheckin(),
		FinishesAt:  checkin.GetFinishedAt(),
	}
}

func transformCheckinList(checkins []checkinDomain.Checkin) []schema.CheckinStation {
	result := make([]schema.CheckinStation, 0, len(checkins))

	for _, c := range checkins {
		result = append(result, transformCheckin(c))
	}

	return result
}

func transformLocationPreliminary(place placeDomain.Place, stations []schema.StationPreliminary) schema.LocationPreliminary {
	return schema.LocationPreliminary{
		Access:    place.GetPlaceAccess(),
		Address:   place.GetPlaceAddress(),
		Id:        place.GetPlaceID().String(),
		Latitude:  place.GetPlaceLatitude(),
		Longitude: place.GetPlaceLongitude(),
		Name:      place.GetPlaceName(),
		Score:     place.GetPlaceScore(),
		IconType:  place.GetPlaceIconType(),
		Stations:  stations,
	}
}

func transformReview(review reviewDomain.Review) schema.Review {
	return schema.Review{
		Id:            review.GetReviewID().String(),
		StationId:     review.GetStationID().String(),
		OutletId:      review.GetOutletID().String(),
		Comment:       review.GetComment(),
		Rating:        review.GetRating(),
		ConnectorType: review.GetConnectorType(),
		UserName:      review.GetUserName(),
		VehicleName:   review.GetVehicleName(),
		VehicleType:   review.GetVehicleType(),
		CreatedAt:     review.GetCreatedAt(),
	}
}

func transformReviewsList(reviewsList []reviewDomain.Review) []schema.Review {
	result := make([]schema.Review, 0, len(reviewsList))

	for i := range reviewsList {
		result = append(result, transformReview(reviewsList[i]))
	}

	return result
}

func transformAmenity(amenity amenityDomain.Amenity) schema.Amenity {
	return schema.Amenity{
		Id:         amenity.GetAmenityID().String(),
		LocationId: amenity.GetLocationID().String(),
		Form:       amenity.GetAmenityForm(),
	}
}

func transformAmenitiesList(amenitiesList []amenityDomain.Amenity) []schema.Amenity {
	result := make([]schema.Amenity, 0, len(amenitiesList))

	for i := range amenitiesList {
		result = append(result, transformAmenity(amenitiesList[i]))
	}

	return result
}

func transformFullStation(station stationDomain.Station, outletsResponse []schema.Outlet) schema.StationFull {
	return schema.StationFull{
		Id:              station.GetStationID().String(),
		Outlets:         outletsResponse,
		Available:       station.GetStationAvailability(),
		Cost:            station.GetStationCost(),
		Name:            station.GetStationName(),
		Manufacturer:    station.GetStationManufacturer(),
		CostDescription: station.GetStationCostDescription(),
		Hours:           station.GetStationWorkingHours(),
		Kilowatts:       station.GetStationKilowatts(),
	}
}

func transformPhoto(photo placeDomain.Photo, yandexLink string) schema.Photo {
	return schema.Photo{
		Id:        photo.GetPhotoID().String(),
		Caption:   photo.GetPhotoCaption(),
		Url:       yandexLink + "/" + photo.GetPhotoName(),
		UserId:    photo.GetUserID().String(),
		CreatedAt: photo.GetCreatedAt(),
	}
}

func transformPhotosList(photosList []placeDomain.Photo, yandexLink string) []schema.Photo {
	result := make([]schema.Photo, 0, len(photosList))

	for i := range photosList {
		result = append(result, transformPhoto(photosList[i], yandexLink))
	}

	return result
}

func transformLocationFull(
	place placeDomain.Place,
	photosList *[]schema.Photo,
	stationsList []schema.StationFull,
	reviewsList *[]schema.Review,
	amenitiesList *[]schema.Amenity,
) schema.LocationFull {
	return schema.LocationFull{
		Access:          place.GetPlaceAccess(),
		Address:         place.GetPlaceAddress(),
		IconType:        place.GetPlaceIconType(),
		Photos:          photosList,
		Id:              place.GetPlaceID().String(),
		Latitude:        place.GetPlaceLatitude(),
		Longitude:       place.GetPlaceLongitude(),
		Name:            place.GetPlaceName(),
		Score:           place.GetPlaceScore(),
		Description:     place.GetDescription(),
		Cost:            place.GetCost(),
		CostDescription: place.GetCostDescription(),
		Hours:           place.GetHours(),
		Open247:         place.GetOpen247(),
		ComingSoon:      place.IsComingSoon(),
		PhoneNumber:     place.GetPhoneNumber(),
		Stations:        stationsList,
		Reviews:         reviewsList,
		Amenities:       amenitiesList,
	}
}
