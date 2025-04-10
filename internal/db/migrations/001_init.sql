-- +goose Up
CREATE TABLE wallet(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);
CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    from_wallet BIGSERIAL REFERENCES wallet(id),
    to_wallet BIGSERIAL REFERENCES wallet(id),
    amount NUMERIC(20, 18),
    transaction_type TEXT NOT NULL CHECK(
        transaction_type in ('deposit', 'withdraw', 'transfer')
    ),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);
CREATE TABLE ledger_entries (
    id BIGSERIAL PRIMARY KEY,
    transaction_id BIGSERIAL NOT NULL REFERENCES transactions(id),
    amount NUMERIC(20, 18),
    direction TEXT NOT NULL CHECK(direction in ('debit', 'credit')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);