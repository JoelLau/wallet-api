package config_test

import (
	"bank-app/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromEnv(t *testing.T) {
	t.Setenv("HTTP_ADDR", "aaaa")
	t.Setenv("MIG_DIR", "bbbb")
	t.Setenv("PG_DSN", "cccc")

	cfg, err := config.FromEnv()
	require.NoError(t, err)
	require.Equal(t, config.Config{
		HTTPAddr:    "aaaa",
		MirationDir: "bbbb",
		PostgresDSN: "cccc",
	}, cfg)
}
