package service

import (
	amenityService "github.com/poorfrombabylon/chargeMeBackend/internal/service/amenity"
	checkinService "github.com/poorfrombabylon/chargeMeBackend/internal/service/checkin"
	outletService "github.com/poorfrombabylon/chargeMeBackend/internal/service/outlet"
	placeService "github.com/poorfrombabylon/chargeMeBackend/internal/service/place"
	reviewService "github.com/poorfrombabylon/chargeMeBackend/internal/service/review"
	stationService "github.com/poorfrombabylon/chargeMeBackend/internal/service/station"
	userService "github.com/poorfrombabylon/chargeMeBackend/internal/service/user"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
)

type Services struct {
	Place   placeService.PlaceService
	Station stationService.StationService
	Outlet  outletService.OutletService
	Review  reviewService.ReviewService
	Checkin checkinService.CheckinService
	Amenity amenityService.AmenityService
	User    userService.UserService
}

func NewServiceRegistry(storages *storage.Storages) *Services {
	place := placeService.NewPlaceService(storages.PlaceStorage)
	station := stationService.NewStationService(storages.StationStorage)
	outlet := outletService.NewOutletService(storages.OutletStorage)
	review := reviewService.NewReviewService(storages.ReviewStorage)
	checkin := checkinService.NewCheckinService(storages.CheckinStorage)
	amenity := amenityService.NewAmenityService(storages.AmenityStorage)
	user := userService.NewUserService(storages.UserStorage)

	return &Services{
		Place:   place,
		Station: station,
		Outlet:  outlet,
		Review:  review,
		Checkin: checkin,
		Amenity: amenity,
		User:    user,
	}
}
