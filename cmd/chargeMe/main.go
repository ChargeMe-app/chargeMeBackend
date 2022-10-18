package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/poorfrombabylon/chargeMeBackend/internal/service"
	"github.com/poorfrombabylon/chargeMeBackend/internal/storage"
	"github.com/poorfrombabylon/chargeMeBackend/libdb"

	"github.com/poorfrombabylon/chargeMeBackend/internal/api"
	"github.com/poorfrombabylon/chargeMeBackend/specs/schema"
	"golang.org/x/sync/errgroup"
)

const (
	host     = "158.160.15.192"
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
		log.Println("connected to db")
	}

	dbx := sqlx.NewDb(db, "pgx")
	libDBWrapper := libdb.NewSQLXDB(dbx)

	storageRegistry := storage.NewStorageRegistry(libDBWrapper)

	serviceRegistry := service.NewServiceRegistry(storageRegistry)

	apiServer := api.NewApiServer(serviceRegistry)

	err = startHttpServer(ctx, apiServer, serviceRegistry)
	if err != nil {
		log.Fatal("failed to start httpServer:", err)
	}
}

func startHttpServer(
	ctx context.Context,
	apiServer schema.ServerInterface,
	serviceRegistry *service.Services,
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

	group.Go(func() error {
		for {
			err := serviceRegistry.Checkin.MoveFinishedCheckinsToReviews(ctx)
			if err != nil {
				log.Println("error while moving chekins to reviews:", err.Error())
			}

			time.Sleep(time.Minute)
		}
	})

	return group.Wait()
}
