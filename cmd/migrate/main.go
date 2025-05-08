package main

import (
	"bank-app/internal/config"
	"bank-app/internal/migrate"
	"context"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ctx := context.Background()
	logr := config.NewLogger()

	cfg, err := config.FromEnv()
	if err != nil {
		panic(err)
	}

	m := migrate.NewMigrate(cfg, migrate.WithLogger(logr))

	logr.InfoContext(ctx, "starting goose migration...")

	if err := m.Run(ctx); err != nil {
		logr.ErrorContext(ctx, "error while running goose migration", slog.Any("error", err))
	}

	logr.InfoContext(ctx, "exiting...")
}
