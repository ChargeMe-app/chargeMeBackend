package integration

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/config"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration/my_ecars"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration/sitronics"
)

type Integration struct {
	SitronicsIntegration sitronics.Integration
	MyECarsIntegration   my_ecars.Integration
}

func NewIntegrationRegistry(conf *config.Config) Integration {
	sitronics := sitronics.NewSitronicsIntegration(conf.Sitronics)
	myECars := my_ecars.NewMyECarsIntegration(conf.MyECars)

	return Integration{
		SitronicsIntegration: sitronics,
		MyECarsIntegration:   myECars,
	}
}
