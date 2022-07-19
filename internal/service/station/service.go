package station

type StationService interface {
}

type StationStorage interface {
}

type service struct {
	stationStorage StationStorage
}

func NewStationService(stationStorage StationStorage) StationService {
	return &service{
		stationStorage: stationStorage,
	}
}
