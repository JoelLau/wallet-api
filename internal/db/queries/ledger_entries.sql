-- name: CreateLedgeEntry :one
INSERT INTO ledger_entries (transaction_id, amount, direction)
VALUES(
        transaction_id = $1,
        amount = $2,
        direction = $3
    )
RETURNING *;
