package outlet

import "github.com/poorfrombabylon/chargeMeBackend/libdb"

type Storage interface {
}

func NewOutletStorage(db libdb.DB) Storage {
	return &outletStorage{db: db}
}

type outletStorage struct {
	db libdb.DB
}
