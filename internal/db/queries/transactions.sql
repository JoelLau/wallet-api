-- name: CreateTransaction :one
INSERT INTO transactions (
        from_wallet,
        to_wallet,
        amount,
        transaction_type
    )
VALUES(
        from_wallet = $1,
        to_wallet = $2,
        amount = $3,
        transaction_type = $4
    )
RETURNING *;

-- name: GetTransactionsByWalletID :many
SELECT *
FROM transactions
WHERE from_wallet = $1
    OR to_wallet = $2;