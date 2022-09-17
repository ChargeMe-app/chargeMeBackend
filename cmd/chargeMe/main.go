package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/poorfrombabylon/chargeMeBackend/internal/service"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"

	"github.com/go-chi/chi/v5"
	"github.com/poorfrombabylon/chargeMeBackend/internal/api"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	"golang.org/x/sync/errgroup"
)

const (
	host     = "176.119.158.240"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "postgres"
)

//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "pass"
//	dbname   = "postgres"
//)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect to database:", err.Error())
	} else {
		fmt.Println("connected to db")
	}

	dbx := sqlx.NewDb(db, "pgx")
	libDBWrapper := libdb.NewSQLXDB(dbx)

	storageRegistry := storage.NewStorageRegistry(libDBWrapper)

	serviceRegistry := service.NewServiceRegistry(storageRegistry)

	apiServer := api.NewApiServer(serviceRegistry)

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
		Addr:    ":8081",
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
