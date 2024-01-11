package app

import (
	"fmt"
	"gateway/config"
	"gateway/internal/controller/http"
	"gateway/internal/service"
	"gateway/pkg/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	r := chi.NewRouter()

	deps := &service.DependenciesServices{
		PortBell:  cfg.ServicesPort.BellPort,
		PortAudio: cfg.ServicesPort.AudioPort,
		PortTrans: cfg.ServicesPort.TransPort,
	}
	settingsCors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			"GET", "HEAD", "OPTIONS", "POST", "PUT", "PATCH",
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	services := service.NewServices(deps)

	http.NewRoutes(r, services)
	srv := server.NewServer(settingsCors.Handler(r), server.Port(cfg.ServerCfg.Port), server.ReadTimeout(cfg.ReadTimeout),
		server.WriteTimeout(cfg.WriteTimeout), server.ShutdownTimeout(cfg.ShutdownTimeout))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app/run: signal: %s\n", s.String())
	case err := <-srv.Notify():
		log.Println(fmt.Errorf("app/run: server.Notify: %s", err.Error()))
	}

	err := srv.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
