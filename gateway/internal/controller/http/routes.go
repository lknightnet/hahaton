package http

import (
	"gateway/internal/service"
	"github.com/go-chi/chi/v5"
)

func NewRoutes(c *chi.Mux, services *service.Services) {
	NewAudioRoutes(c, services.MA)
	NewBellRoutes(c, services.MB)
	NewTransRoutes(c, services.MT)
}
