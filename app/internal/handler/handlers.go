// Package handler provides HTTP request handlers, routing, and middleware.
// It handles incoming HTTP requests, delegates to the service layer for
// business logic, and returns JSON responses with appropriate status codes.
package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"hotel.com/app/internal/service"
)

type Handler struct {
	s service.Service
	l *slog.Logger
}

func New(s service.Service, l *slog.Logger) *Handler {
	return &Handler{
		s: s,
		l: l,
	}
}

func (h *Handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	status := "healthy"
	httpStatus := http.StatusOK
	if err := h.s.Check(); err != nil {
		status = "down"
		httpStatus = http.StatusServiceUnavailable
	}
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(map[string]string{"status": status})
}
