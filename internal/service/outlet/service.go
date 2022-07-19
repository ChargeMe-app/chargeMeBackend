package outlet

type OutletService interface {
}

type OutletStorage interface {
}

type service struct {
	outletStorage OutletStorage
}

func NewOutletService(outletStorage OutletStorage) OutletService {
	return &service{
		outletStorage: outletStorage,
	}
}
