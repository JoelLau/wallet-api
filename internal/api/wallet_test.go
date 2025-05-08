package api_test

import (
	"bank-app/internal/config"
	"bank-app/internal/db"
	"bank-app/internal/migrate"
	testutils "bank-app/internal/test-utils"
	"io"
	"log/slog"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestCreateWallet(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	dsn := testutils.NewTestDB(t)

	conn, err := pgx.Connect(ctx, dsn)
	require.NoError(t, err)

	repo := db.NewSQLCRepository(conn)

	m := migrate.NewMigrate(config.Config{
		PostgresDSN: dsn,
		MirationDir: "../../internal/db/migrations/",
	}, migrate.WithLogger(slog.New(slog.NewTextHandler(io.Discard, nil))))

	err = m.Run(ctx)
	require.NoError(t, err)

	var originalCount int
	conn.QueryRow(ctx, "SELECT COUNT(*) FROM wallets;").
		Scan(&originalCount)
	require.Equal(t, 0, originalCount)

	w, err := repo.CreateWallet(ctx)
	require.NoError(t, err)
	require.Equal(t, int64(1), w.ID)
	require.False(t, w.CreatedAt.IsZero())

	var newCount int
	conn.QueryRow(ctx, "SELECT COUNT(*) FROM wallets;").
		Scan(&newCount)
	require.Equal(t, 1, newCount)
}
