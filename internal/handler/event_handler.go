package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/revpos/eTrak/internal/models"
	"github.com/revpos/eTrak/internal/service"
)

type EventHandler struct {
	Service *service.EventService
}

func NewEventHandler(s *service.EventService) *EventHandler {
	return &EventHandler{Service: s}
}

func (h *EventHandler) TrackEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// enforce timestamp if missing
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}

	if err := h.Service.TrackEvent(r.Context(), &event); err != nil {
		http.Error(w, "failed to store event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
