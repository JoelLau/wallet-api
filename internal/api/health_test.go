package api_test

import (
	"bank-app/internal/api"
	repo "bank-app/internal/repository"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetHealthCheck(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/health-check", nil)
	w := httptest.NewRecorder()

	s := api.NewServer(&repo.MockRepository{})
	s.GetHealthCheck(w, req)

	resp := w.Result()

	_, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode)
}
