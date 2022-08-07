package service

import (
	amenityService "github.com/poorfrombabylon/chargeMeBackend/internal/service/amenity"
	outletService "github.com/poorfrombabylon/chargeMeBackend/internal/service/outlet"
	placeService "github.com/poorfrombabylon/chargeMeBackend/internal/service/place"
	reviewService "github.com/poorfrombabylon/chargeMeBackend/internal/service/review"
	stationService "github.com/poorfrombabylon/chargeMeBackend/internal/service/station"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
)

type Services struct {
	Place   placeService.PlaceService
	Station stationService.StationService
	Outlet  outletService.OutletService
	Review  reviewService.ReviewService
	Amenity amenityService.AmenityService
}

func NewServiceRegistry(storages *storage.Storages) *Services {
	place := placeService.NewPlaceService(storages.PlaceStorage)
	station := stationService.NewStationService(storages.StationStorage)
	outlet := outletService.NewOutletService(storages.OutletStorage)
	review := reviewService.NewReviewService(storages.ReviewStorage)
	amenity := amenityService.NewAmenityService(storages.AmenityStorage)

	return &Services{
		Place:   place,
		Station: station,
		Outlet:  outlet,
		Review:  review,
		Amenity: amenity,
	}
}
