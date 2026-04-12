package handler

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v3"
)

func (h *Handler) NewServerMux() *chi.Mux {
	r := chi.NewRouter()

	// Global middleware
	r.Use(httplog.RequestLogger(h.l, &httplog.Options{
		Level:         slog.LevelDebug,
		Schema:        httplog.SchemaOTEL,
		RecoverPanics: true,
	}))

	r.Use(SecureHeaders)

	// Routes
	r.Get("/health", h.healthCheck)
	return r
}
