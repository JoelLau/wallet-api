package api

import "net/http"

// CreateWallet implements genapi.ServerInterface.
func (s *Server) CreateWallet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

// CreateWalletDeposit implements genapi.ServerInterface.
func (s *Server) CreateWalletDeposit(w http.ResponseWriter, r *http.Request, walletId string) {
}

// CreateWalletWithdrawal implements genapi.ServerInterface.
func (s *Server) CreateWalletWithdrawal(w http.ResponseWriter, r *http.Request, walletId string) {
}

// GetWalletBalance implements genapi.ServerInterface.
func (s *Server) GetWalletBalance(w http.ResponseWriter, r *http.Request, walletId string) {
}

// GetWalletTransactions implements genapi.ServerInterface.
func (s *Server) GetWalletTransactions(w http.ResponseWriter, r *http.Request, walletId string) {
}
