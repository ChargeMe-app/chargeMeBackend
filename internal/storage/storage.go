package storage

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/outlet"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/place"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/station"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

type Storages struct {
	PlaceStorage   place.Storage
	StationStorage station.Storage
	OutletStorage  outlet.Storage
}

func NewStorageRegistry(db libdb.DB) *Storages {
	return &Storages{
		PlaceStorage:   place.NewPlaceStorage(db),
		StationStorage: station.NewStationStorage(db),
		OutletStorage:  outlet.NewOutletStorage(db),
	}
}
