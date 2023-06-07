-- name: CreateWallet :one
INSERT INTO wallets (
    wallet_address,
    created_block
) VALUES ($1, $2) RETURNING *;

-- name: GetListWallet :many
SELECT wallet_address FROM wallets;
