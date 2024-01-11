package app

import (
	"bell/config"
	"bell/internal/controller/http"
	"bell/internal/repository"
	"bell/internal/service"
	"bell/pkg/database/postgres"
	"bell/pkg/server"
	"context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	ctx := context.Background()

	pg, err := postgres.New(ctx, cfg.PG.ConnURI)
	if err != nil {
		log.Println(err)
	}

	repositories := repository.NewRepositories(pg)

	deps := &service.Dependencies{
		Repositories: repositories,
		URLAPI:       cfg.UrlAPI,
	}

	services := service.NewServices(deps)
	err = services.B.Set(ctx)
	if err != nil {
		log.Println(err)
	}

	rout := mux.NewRouter()
	http.NewBellRoutes(rout, services.B)
	srv := server.NewServer(rout, server.Port(cfg.HTTP.Port), server.ReadTimeout(cfg.ReadTimeout),
		server.WriteTimeout(cfg.WriteTimeout), server.ShutdownTimeout(cfg.ShutdownTimeout))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("Run: " + s.String())
	case err := <-srv.Notify():
		log.Println(errors.Wrap(err, "Run: signal.Notify"))
	}

	err = srv.Shutdown()
	if err != nil {
		log.Println(errors.Wrap(err, "Run: server shutdown"))
	}
}
