package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/poorfrombabylon/chargeMeBackend/internal/service"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"

	"github.com/go-chi/chi/v5"
	"github.com/poorfrombabylon/chargeMeBackend/internal/api"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	"golang.org/x/sync/errgroup"
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

	drv := stdlib.GetDefaultDriver().(*stdlib.Driver)

	ctor, err := drv.OpenConnector("postgres://postgres:pass@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	db := sql.OpenDB(ctor)
	dbx := sqlx.NewDb(db, "pgx")
	libDBWrapper := libdb.NewSQLXDB(dbx)

	storageRegistry := storage.NewStorageRegistry(libDBWrapper)

	_ = service.NewServiceRegistry(storageRegistry)

	err = startHttpServer(ctx, apiServer)
	if err != nil {
		log.Fatal("failed to start httpServer:", err)
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
