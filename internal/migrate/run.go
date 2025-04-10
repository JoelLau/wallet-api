package migrate

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose"
)

func (m *Migrate) Run(ctx context.Context) error {
	db, err := sql.Open(PostgresDriver, m.cfg.PG_DSN)
	if err != nil {
		return fmt.Errorf("failed to open DB instance: %w", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	if err := goose.Run("up", db, m.cfg.MirationDir); err != nil {
		return fmt.Errorf("failed to run migration: %w", err)
	}

	return nil
}
