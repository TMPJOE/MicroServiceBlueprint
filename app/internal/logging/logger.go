package logging

import (
	"log/slog"
	"os"

	"github.com/go-chi/httplog/v3"
)

func New() *slog.Logger {
	logFormat := httplog.SchemaECS.Concise(true)
	// handler options
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: logFormat.ReplaceAttr,
	})).With(
		slog.String("service", "example-service"),
		slog.String("version", "v0.0.1"),
		slog.String("env", "develpment"),
	)

	return logger
}
