package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/zuekveer/BWA/internal/entity"

	"github.com/go-chi/chi/v5"
)

const dateFormat = "2006-01-02"

func (h *handlers) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var event entity.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	_, err := h.service.Create(ctx, event)
	if err != nil {
		http.Error(w, "failed to update event", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *handlers) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	var changes entity.EventChanges
	if err := json.NewDecoder(r.Body).Decode(&changes); err != nil {
		http.Error(w, "failed to parse body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := h.service.Update(ctx, id, changes); err != nil {
		http.Error(w, "failed to update event", http.StatusBadRequest)
		return
	}
}

func (h *handlers) getEvents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dateFrom := r.URL.Query().Get("from")
	dateTo := r.URL.Query().Get("to")

	from, err := time.Parse(dateFormat, dateFrom)
	if err != nil {
		http.Error(w, "failed to parse date From. Expected YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	to, err := time.Parse(dateFormat, dateTo)
	if err != nil {
		http.Error(w, "failed to parse date To. Expected YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	evs, err := h.service.Get(ctx, from, to)
	if err != nil {
		http.Error(w, "failed to get Events", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(evs); err != nil {
		http.Error(w, "failed to encode Events", http.StatusBadRequest)
		return
	}
}
