package station

import "github.com/poorfrombabylon/chargeMeBackend/libdb"

type Storage interface {
}

func NewStationStorage(db libdb.DB) Storage {
	return &stationStorage{db: db}
}

type stationStorage struct {
	db libdb.DB
}
