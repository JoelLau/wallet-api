package main

import (
	"bank-app/internal/app"
	"bank-app/internal/config"
	"context"
	"log/slog"
)

//go:generate mkdir -p ../../internal/api/gen
//go:generate oapi-codegen --config=../../internal/api/_config.yaml ../../internal/api/_openapi.yaml
func main() {
	ctx := context.Background()
	logr := config.NewLogger()

	cfg := config.FromEnv()
	a := app.NewApp(cfg, app.WithLogger(logr))

	logr.InfoContext(ctx, "Starting Wallet REST server...")
	err := a.Run(ctx)
	if err != nil {
		logr.ErrorContext(ctx, "Error with Wallet REST server", slog.Any("error", err))
	}

	logr.InfoContext(ctx, "Exiting...")
}
