package app

import (
	db "bank-app/internal/db/gen"
	"log/slog"
)

type OptFunc func(*App)

func WithLogger(logr *slog.Logger) OptFunc {
	return func(a *App) {
		a.logr = logr
	}
}

type Config struct {
	Addr string // e.g. ":8080"
	DBTX db.DBTX
}
