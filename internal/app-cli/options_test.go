package appcli

import (
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithLogger(t *testing.T) {
	t.Parallel()

	before := slog.New(slog.NewTextHandler(io.Discard, nil))
	after := slog.New(slog.NewJSONHandler(io.Discard, nil))

	app := &App{logr: before}

	WithLogger(after)(app)

	require.Same(t, after, app.logr)

}

func TestWithHTTPAddress(t *testing.T) {
	t.Parallel()

	app := &App{httpAddr: "before"}
	WithHTTPAddress("after")(app)

	require.Equal(t, "after", app.httpAddr)
}
