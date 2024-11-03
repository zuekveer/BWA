// handlers/handlers.go
package handlers

import (
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("welcome"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
