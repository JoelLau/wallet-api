package main

import (
	"bank-app/internal/app"
	"bank-app/internal/config"
	db "bank-app/internal/db/gen"
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

	cfg := config.FromEnv()

	db, err := NewDbConn(ctx, cfg.PG_DSN)
	if err != nil {
		logr.ErrorContext(ctx, "error connecting to database", slog.Any("error", err))
		panic(err)
	}

	a := app.NewApp(app.Config{
		Addr: cfg.HTTPAddr,
		DBTX: db,
	}, app.WithLogger(logr))

	logr.InfoContext(ctx, "Starting Wallet REST server...")

	if err := a.Run(ctx); err != nil {
		logr.ErrorContext(ctx, "Error with Wallet REST server", slog.Any("error", err))
	}

	logr.InfoContext(ctx, "Exiting...")
}

func NewDbConn(ctx context.Context, dsn string) (db.DBTX, error) {
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}
