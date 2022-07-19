package place

import "github.com/poorfrombabylon/chargeMeBackend/libdb"

type Storage interface {
}

func NewPlaceStorage(db libdb.DB) Storage {
	return &placeStorage{db: db}
}

type placeStorage struct {
	db libdb.DB
}
