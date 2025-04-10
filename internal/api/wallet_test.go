package api_test

// import (
// 	"bank-app/internal/api"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func TestCreateWallet(t *testing.T) {
// 	t.Parallel()

// 	walletID := "some-uuid"

// 	req := httptest.NewRequest(
// 		http.MethodGet,
// 		fmt.Sprintf("/api/v1/wallets/%s", walletID),
// 		nil,
// 	)
// 	w := httptest.NewRecorder()

// 	s := api.NewServer()
// 	s.CreateWallet(w, req)

// 	resp := w.Result()
// 	_, err := io.ReadAll(resp.Body)
// 	require.NoError(t, err)

// 	require.Equal(t, http.StatusCreated, resp.StatusCode)
// }

// func TestCreateWalletDeposit(t *testing.T) {
// 	t.Parallel()

// 	walletID := "some-uuid"

// 	req := httptest.NewRequest(
// 		http.MethodGet,
// 		fmt.Sprintf("/api/v1/wallets/%s", walletID),
// 		nil,
// 	)
// 	w := httptest.NewRecorder()

// 	s := api.NewServer()
// 	s.CreateWalletDeposit(w, req, walletID)

// 	resp := w.Result()
// 	_, err := io.ReadAll(resp.Body)
// 	require.NoError(t, err)

// 	require.Equal(t, http.StatusOK, resp.StatusCode)
// }

// func TestCreateWalletWithdrawal(t *testing.T) {
// 	t.Parallel()

// 	walletID := "some-uuid"

// 	req := httptest.NewRequest(
// 		http.MethodGet,
// 		fmt.Sprintf("/api/v1/wallets/%s", walletID),
// 		nil,
// 	)
// 	w := httptest.NewRecorder()

// 	s := api.NewServer()
// 	s.CreateWalletWithdrawal(w, req, walletID)

// 	resp := w.Result()
// 	_, err := io.ReadAll(resp.Body)
// 	require.NoError(t, err)

// 	require.Equal(t, http.StatusOK, resp.StatusCode)

// }

// func TestGetWalletBalance(t *testing.T) {
// 	t.Parallel()

// 	walletID := "some-uuid"

// 	req := httptest.NewRequest(
// 		http.MethodGet,
// 		fmt.Sprintf("/api/v1/wallets/%s", walletID),
// 		nil,
// 	)
// 	w := httptest.NewRecorder()

// 	s := api.NewServer()
// 	s.GetWalletBalance(w, req, walletID)

// 	resp := w.Result()
// 	_, err := io.ReadAll(resp.Body)
// 	require.NoError(t, err)

// 	require.Equal(t, http.StatusOK, resp.StatusCode)

// }

// func TestGetWalletTransactions(t *testing.T) {
// 	t.Parallel()

// 	walletID := "some-uuid"

// 	req := httptest.NewRequest(
// 		http.MethodGet,
// 		fmt.Sprintf("/api/v1/wallets/%s", walletID),
// 		nil,
// 	)
// 	w := httptest.NewRecorder()

// 	s := api.NewServer()
// 	s.GetWalletTransactions(w, req, walletID)

// 	resp := w.Result()
// 	_, err := io.ReadAll(resp.Body)
// 	require.NoError(t, err)

// 	require.Equal(t, http.StatusOK, resp.StatusCode)

// }
