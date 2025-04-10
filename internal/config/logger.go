package config

import (
	"log/slog"
	"os"
)

// TODO: configure verbosity
// TODO: configure io.Writer
func NewLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stderr,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			},
		),
	)
}
