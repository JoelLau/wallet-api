package app

import (
	"bank-app/internal/api"
	genapi "bank-app/internal/api/gen"
	"bank-app/internal/config"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	cfg  config.Config
	logr *slog.Logger
}

func NewApp(cfg config.Config, opts ...OptFunc) *App {
	app := &App{cfg: cfg}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func (a *App) Run(ctx context.Context) error {
	cctx := context.WithoutCancel(ctx)

	server := api.NewServer()

	r := chi.NewMux()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(2 * time.Second))

	h := genapi.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    a.cfg.HTTPAddr,
	}

	err := s.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		a.logr.InfoContext(cctx, "server shutdown")
		return nil
	} else if err != nil {
		a.logr.ErrorContext(cctx, "unexpected server error", slog.Any("error", err))
		return err
	}

	return nil
}
