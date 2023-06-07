-- name: AddTxn :one
INSERT INTO transactions (
    wallet_address,
    chain,
    chain_id,
    txn_hash,
    from_add,
    to_add
) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetOneTxn :one
SELECT * FROM transactions
WHERE txn_hash = $1
LIMIT 1;

-- name: GetAllTxn :many
SELECT * FROM transactions
WHERE wallet_address = $1 AND chain = $2;