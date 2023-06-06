-- name: AddTxn :one
INSERT INTO wallets (
    wallet_address,
    chain,
    chain_id,
    txn_hash,
    from,
    to
    ) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;