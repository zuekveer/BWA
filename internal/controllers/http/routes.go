// http/http.go
package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zuekveer/BWA/internal/controllers/handlers"
)

func NewHandlers() *chi.Mux {
	r := chi.NewMux()
	r.Use(middleware.Recoverer)

	h := handlers.NewHandler()

	r.Get("/", h.Welcome)

	return r
}
