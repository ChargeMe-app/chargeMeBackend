package service

import (
	outletService "github.com/poorfrombabylon/chargeMeBackend/internal/service/outlet"
	placeService "github.com/poorfrombabylon/chargeMeBackend/internal/service/place"
	stationService "github.com/poorfrombabylon/chargeMeBackend/internal/service/station"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
)

type Services struct {
	Place   placeService.PlaceService
	Station stationService.StationService
	Outlet  outletService.OutletService
}

func NewServiceRegistry(storages *storage.Storages) *Services {
	place := storages.PlaceStorage
	station := storages.StationStorage
	outlet := storages.OutletStorage

	return &Services{
		Place:   place,
		Station: station,
		Outlet:  outlet,
	}
}
