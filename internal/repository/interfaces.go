package repo

import (
	db "bank-app/internal/db/gen"
	"context"

	"github.com/jackc/pgx/v5"
)

type RepositoryService interface {
	CreateLedgeEntry(context.Context, db.CreateLedgeEntryParams) (db.LedgerEntry, error)
	CreateTransaction(context.Context, db.CreateTransactionParams) (db.Transaction, error)
	CreateWallet(context.Context) (db.Wallet, error)
	GetTransactionsByWalletID(context.Context, db.GetTransactionsByWalletIDParams) ([]db.Transaction, error)
}

type Repository interface {
	RepositoryService

	WithTx(pgx.Tx) *db.Queries // for convenience
}

var x *db.Queries
var _ RepositoryService = x
