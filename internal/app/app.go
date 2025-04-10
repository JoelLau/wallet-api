package app

import (
	"bank-app/internal/api"
	genapi "bank-app/internal/api/gen"
	db "bank-app/internal/db/gen"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	cfg Config

	logr *slog.Logger
}

func NewApp(cfg Config, opts ...OptFunc) *App {
	app := &App{cfg: cfg, logr: slog.Default()}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func (app *App) Run(ctx context.Context) error {
	cctx := context.WithoutCancel(ctx)

	server := api.NewServer(db.New(app.cfg.DBTX))

	httpServer := &http.Server{
		Handler: NewHandler(server),
		Addr:    app.cfg.Addr,
	}

	if err := httpServer.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		app.logr.InfoContext(cctx, "server shutdown")
		return nil
	} else if err != nil {
		app.logr.ErrorContext(cctx, "unexpected server error", slog.Any("error", err))
		return err
	}

	return nil
}

func NewHandler(server genapi.ServerInterface) http.Handler {
	return genapi.HandlerFromMux(server, NewRouter())
}

func NewRouter() *chi.Mux {
	r := chi.NewMux()

	// recommended middleware stack.
	// see more: https://github.com/go-chi/chi/blob/d7034fdfdaefd10f1bc1a7b813bc979f2eda3a36/README.md?plain=1#L86-L95
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(2 * time.Second))

	return r
}
