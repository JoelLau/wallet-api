package api_test

import (
	"bank-app/internal/api"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/transactions",
		nil,
	)
	w := httptest.NewRecorder()

	s := api.NewServer()
	s.CreateTransfer(w, req)

	resp := w.Result()
	_, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode)
}
