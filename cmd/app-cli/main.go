package main

import (
	appcli "bank-app/internal/app-cli"
	"bank-app/internal/config"
	"bank-app/internal/db"
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate mkdir -p ../../internal/api/gen
//go:generate oapi-codegen --config=../../internal/api/_config.yaml ../../internal/api/_openapi.yaml
//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate -f ../../internal/db/sqlc.yaml
func main() {
	ctx := context.Background()
	logr := config.NewLogger()

	cfg, err := config.FromEnv()
	if err != nil {
		logr.ErrorContext(ctx, "error fetching config", slog.Any("error", err))
		return
	}

	logr.InfoContext(ctx, "env", slog.Any("PG_DSN", cfg.PostgresDSN))

	dbtx, err := NewDBTX(ctx, cfg.PostgresDSN)
	if err != nil {
		logr.ErrorContext(ctx, "error connecting to postgres instance", slog.Any("error", err))
		return
	}
	repo := db.NewSQLCRepository(dbtx)

	app := appcli.NewApp(repo, appcli.WithLogger(logr))

	logr.InfoContext(ctx, "Starting Wallet REST server...")

	if err := app.Run(ctx); err != nil {
		logr.ErrorContext(ctx, "Error with Wallet REST server", slog.Any("error", err))
	}

	logr.InfoContext(ctx, "Exiting...")
}

func NewDBTX(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, err
}
