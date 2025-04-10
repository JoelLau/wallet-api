package app

import "log/slog"

type Config struct {
	Addr string `ENV:"http_addr"`
}

type OptFunc func(*App)

func WithLogger(logr *slog.Logger) OptFunc {
	return func(a *App) {
		a.logr = logr
	}
}
