package place

type PlaceService interface {
}

type PlaceStorage interface {
}

type service struct {
	placeStorage PlaceStorage
}

func NewPlaceService(placeStorage PlaceStorage) PlaceService {
	return &service{
		placeStorage: placeStorage,
	}
}
