package http

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/zuekveer/BWA/internal/logger"
	"github.com/zuekveer/BWA/internal/service"
)

type handlers struct {
	service *service.EventService
	l       logger.Logger
}

// NewHandlers initializes the router and registers routes
func NewHandlers(events *service.EventService, log logger.Logger) *chi.Mux {
	h := handlers{
		service: events,
		l:       log,
	}
	r := chi.NewMux()
	h.build(r)
	return r
}

// build sets up routes for the handlers
func (h *handlers) build(r chi.Router) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Welcome endpoint at root
	r.Get("/", h.welcome)

	// Event-related endpoints
	r.Route("/api/v1/events", func(r chi.Router) {
		r.Post("/", h.create)
		r.Patch("/{id}", h.update)
		r.Get("/", h.getEvents)
	})
}

// welcome serves a welcome message
func (h *handlers) welcome(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("welcome"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
