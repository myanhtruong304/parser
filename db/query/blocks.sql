-- name: AddBlock :one
INSERT INTO blocks (
    block_number,
    processed,
    chain_id
) VALUES ($1, $2, $3) RETURNING block_number;

-- name: GetOneBlock :one
SELECT block_number FROM blocks
WHERE block_number = $1 AND chain_id = $2
ORDER BY block_number ASC
LIMIT 1 ;

-- name: GetNotProcessBlock :many
SELECT * FROM blocks
WHERE processed = $1 AND chain_id = $2
ORDER BY block_number ASC;

-- name: UpdateBlockProcess :one
UPDATE blocks
SET 
    processed = $3
WHERE block_number = $1 AND chain_id = $2 RETURNING *;

-- name: GetAllBlock :many
SELECT block_number FROM blocks
WHERE chain_id = $1
ORDER BY block_number ASC;