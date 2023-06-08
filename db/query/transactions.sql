-- name: AddTxn :one
INSERT INTO transactions (
    wallet_address,
    chain,
    chain_id,
    txn_hash,
    from_address,
    to_address,
    block_created_at,
    block,
    status, 
    created_at,
    sequence,
    type,
    fee,
    metadata
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING txn_hash;

-- name: GetOneTxn :one
SELECT * FROM transactions
WHERE txn_hash = $1 AND chain_id = $2
LIMIT 1;

-- name: GetAllTxn :many
SELECT * FROM transactions
WHERE wallet_address = $1 AND chain_id = $2;