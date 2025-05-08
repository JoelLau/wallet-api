package appcli

import (
	repo "bank-app/internal/db"
	db "bank-app/internal/db/gen"
	"log/slog"
)

type OptFunc func(*App)

func WithLogger(logr *slog.Logger) OptFunc {
	return func(a *App) {
		a.logr = logr
	}
}

func WithHTTPAddress(addr string) OptFunc {
	return func(a *App) {
		a.httpAddr = addr
	}
}

type Config struct {
	Addr string // e.g. ":8080"
	DBTX db.DBTX
	Repo repo.Repository
}
