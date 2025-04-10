package migrate

import (
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type OptFunc func(*Migrate)

func WithLogger(logr *slog.Logger) OptFunc {
	return func(a *Migrate) {
		a.logr = logr
	}
}
