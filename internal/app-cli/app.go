package appcli

import (
	"bank-app/internal/api"
	genapi "bank-app/internal/api/gen"
	"bank-app/internal/db"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	repo db.Repository

	httpAddr string
	logr     *slog.Logger
}

// validate config here. fail early!
func NewApp(repo db.Repository, opts ...OptFunc) *App {
	app := &App{
		repo:     repo,
		logr:     slog.Default(),
		httpAddr: ":8080",
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

// should be as fast as possible. should already be configured.
func (app *App) Run(ctx context.Context) error {
	cctx := context.WithoutCancel(ctx)

	server := api.NewServer(app.repo)
	handler := NewHandler(server)

	httpServer := &http.Server{
		Handler: handler,
		Addr:    app.httpAddr,
	}

	app.logr.InfoContext(ctx, fmt.Sprintf("listening on %s", httpServer.Addr))

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
	r.Use(middleware.Logger) // TODO: replace logger
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(2 * time.Second))

	return r
}
