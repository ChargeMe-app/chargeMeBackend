package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/poorfrombabylon/chargeMeBackend/internal/integration"

	"github.com/poorfrombabylon/chargeMeBackend/internal/config"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("error while init config")
	}

	integr := integration.NewIntegrationRegistry(cfg)

	kek := integr.SitronicsIntegration

	mem, _ := kek.GetStationByName(ctx, "10001")

	fmt.Println(mem.CPCard.Connectors[0].StatusName)
}
