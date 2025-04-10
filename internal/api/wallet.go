package api

import (
	"context"
	"encoding/json"
	"net/http"
)

func (s *Server) CreateWallet(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	wallet, err := s.repo.CreateWallet(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := json.Marshal(wallet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (s *Server) CreateWalletDeposit(w http.ResponseWriter, r *http.Request, walletId string) {
}

func (s *Server) CreateWalletWithdrawal(w http.ResponseWriter, r *http.Request, walletId string) {
}

func (s *Server) GetWalletBalance(w http.ResponseWriter, r *http.Request, walletId string) {
}

func (s *Server) GetWalletTransactions(w http.ResponseWriter, r *http.Request, walletId string) {
}
