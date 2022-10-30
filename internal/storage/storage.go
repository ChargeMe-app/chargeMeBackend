package storage

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/amenity"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/checkin"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/outlet"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/photo"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/place"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/review"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/station"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage/user"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"
)

type Storages struct {
	PlaceStorage   place.Storage
	StationStorage station.Storage
	OutletStorage  outlet.Storage
	ReviewStorage  review.Storage
	PhotoStorage   photo.Storage
	CheckinStorage checkin.Storage
	AmenityStorage amenity.Storage
	UserStorage    user.Storage
}

func NewStorageRegistry(db libdb.DB) *Storages {
	return &Storages{
		PlaceStorage:   place.NewPlaceStorage(db),
		StationStorage: station.NewStationStorage(db),
		OutletStorage:  outlet.NewOutletStorage(db),
		ReviewStorage:  review.NewReviewStorage(db),
		PhotoStorage:   photo.NewPhotoStorage(db),
		CheckinStorage: checkin.NewCheckinStorage(db),
		AmenityStorage: amenity.NewAmenityStorage(db),
		UserStorage:    user.NewUserStorage(db),
	}
}
