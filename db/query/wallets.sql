-- name: CreateWallet :one
INSERT INTO wallets (
    wallet_address,
    created_block,
    created_at
) VALUES ($1, $2, $3) RETURNING *;
