package migrate

import (
	"bank-app/internal/config"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const PostgresDriver = "pgx"

type Migrate struct {
	cfg  config.Config
	logr *slog.Logger
}

func NewMigrate(cfg config.Config, opts ...OptFunc) *Migrate {
	m := &Migrate{cfg: cfg}

	for _, opt := range opts {
		opt(m)
	}

	return m
}
