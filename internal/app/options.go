package app

import "log/slog"

type OptFunc func(*App)

func WithLogger(logr *slog.Logger) OptFunc {
	return func(a *App) {
		a.logr = logr
	}
}
