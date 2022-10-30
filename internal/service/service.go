package service

import (
	amenityService "github.com/poorfrombabylon/chargeMeBackend/internal/service/amenity"
	checkinService "github.com/poorfrombabylon/chargeMeBackend/internal/service/checkin"
	outletService "github.com/poorfrombabylon/chargeMeBackend/internal/service/outlet"
	photoService "github.com/poorfrombabylon/chargeMeBackend/internal/service/photo"
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
	Photo   photoService.PhotoService
	Checkin checkinService.CheckinService
	Amenity amenityService.AmenityService
	User    userService.UserService
}

func NewServiceRegistry(storages *storage.Storages) *Services {
	place := placeService.NewPlaceService(storages.PlaceStorage, storages.ReviewStorage)
	station := stationService.NewStationService(storages.StationStorage)
	outlet := outletService.NewOutletService(storages.OutletStorage)
	review := reviewService.NewReviewService(storages.ReviewStorage, storages.PlaceStorage, storages.StationStorage)
	photo := photoService.NewPhotoService(storages.PhotoStorage)
	checkin := checkinService.NewCheckinService(
		storages.CheckinStorage,
		storages.ReviewStorage,
		storages.OutletStorage,
		storages.StationStorage,
		storages.PlaceStorage,
	)
	amenity := amenityService.NewAmenityService(storages.AmenityStorage)
	user := userService.NewUserService(storages.UserStorage)

	return &Services{
		Place:   place,
		Station: station,
		Outlet:  outlet,
		Review:  review,
		Photo:   photo,
		Checkin: checkin,
		Amenity: amenity,
		User:    user,
	}
}
