-- name: CreateWallet :one
INSERT INTO wallets DEFAULT
VALUES
RETURNING *;