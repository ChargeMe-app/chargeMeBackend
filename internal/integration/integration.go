package integration

import (
	"github.com/poorfrombabylon/chargeMeBackend/internal/config"
	"github.com/poorfrombabylon/chargeMeBackend/internal/integration/sitronics"
)

type Integration struct {
	SitronicsIntegration sitronics.Integration
}

func NewIntegrationRegistry(conf *config.Config) Integration {
	sitronics := sitronics.NewSitronicsIntegration(conf.Sitronics)

	return Integration{
		SitronicsIntegration: sitronics,
	}
}
