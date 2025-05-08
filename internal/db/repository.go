package db

import (
	db "bank-app/internal/db/gen"
	"context"

	"github.com/jackc/pgx/v5"
)

// for replacing repository methods in tests.
// consider splitting this up if the domain gets too big.
type RepositoryService interface {
	CreateLedgeEntry(context.Context, db.CreateLedgeEntryParams) (db.LedgerEntry, error)
	CreateTransaction(context.Context, db.CreateTransactionParams) (db.Transaction, error)
	CreateWallet(context.Context) (db.Wallet, error)
	GetTransactionsByWalletID(context.Context, db.GetTransactionsByWalletIDParams) ([]db.Transaction, error)
}

// adds DB specific operations to the interface.
// TODO: weak abstraction, rethink this!
type Repository interface {
	RepositoryService

	WithTx(pgx.Tx) *db.Queries // for convenience
}

// sanity checks to make sure no methods were missed.
var x *db.Queries
var _ RepositoryService = x
