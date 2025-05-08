package db

import (
	db "bank-app/internal/db/gen"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

// test utility for convenience
type MockRepository struct {
	NewID   func() int64
	TimeNow func() time.Time
}

func (m *MockRepository) CreateLedgeEntry(context.Context, db.CreateLedgeEntryParams) (db.LedgerEntry, error) {
	return db.LedgerEntry{}, nil
}

func (m *MockRepository) CreateTransaction(context.Context, db.CreateTransactionParams) (db.Transaction, error) {
	return db.Transaction{}, nil
}

func (m *MockRepository) CreateWallet(context.Context) (db.Wallet, error) {
	return db.Wallet{
		ID:        m.NewID(),
		CreatedAt: m.TimeNow(),
	}, nil
}

func (m *MockRepository) GetTransactionsByWalletID(context.Context, db.GetTransactionsByWalletIDParams) ([]db.Transaction, error) {
	return []db.Transaction{}, nil
}

func (m *MockRepository) WithTx(tx pgx.Tx) (q *db.Queries) {
	return q
}

var _ Repository = &MockRepository{}
