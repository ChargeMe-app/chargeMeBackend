package main

import (
	"chargeMe/internal/api"
	"chargeMe/specs/schema"
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
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

	apiServer := api.NewApiServer()

	err := startHttpServer(ctx, apiServer)
	if err != nil {
		log.Fatal("failed to start httpServer", err)
	}

}

func startHttpServer(
	ctx context.Context,
	apiServer schema.ServerInterface,
	middlewares ...schema.MiddlewareFunc,
) error {
	handler := schema.HandlerWithOptions(apiServer, schema.ChiServerOptions{
		BaseURL:     "",
		Middlewares: middlewares,
	})

	router := chi.NewRouter()
	router.Handle("/*", handler)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		return httpServer.Shutdown(ctx)
	})

	return group.Wait()
}
