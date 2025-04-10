package main

import (
	"bank-app/internal/app"
	"context"
	"log/slog"
)

// scaffolds application
//
//go:generate mkdir -p ../../internal/api/gen
//go:generate oapi-codegen --config=../../internal/api/_config.yaml ../../internal/api/_openapi.yaml
func main() {
	ctx := context.Background()
	logr := NewLogger()

	a := app.NewApp(app.Config{}, app.WithLogger(logr))

	logr.InfoContext(ctx, "Starting Wallet REST API...")
	err := a.Run(ctx)
	if err != nil {
		logr.ErrorContext(ctx, "app stopped running error", slog.Any("error", err))
	}

	logr.InfoContext(ctx, "exiting...")
}

func NewLogger() *slog.Logger {
	return slog.Default()
}
