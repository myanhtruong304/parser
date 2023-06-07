-- name: AddBlock :one
INSERT INTO blocks (
    block_number,
    processed
) VALUES ($1, $2) RETURNING block_number;

-- name: GetOneBlock :one
SELECT block_number FROM blocks
WHERE block_number = $1
ORDER BY block_number ASC
LIMIT 1 ;

-- name: GetNotProcessBlock :one
SELECT * FROM blocks
WHERE processed = $1
ORDER BY block_number ASC
LIMIT 1;

-- name: UpdateBlockProcess :one
UPDATE blocks
SET 
    processed = $2
WHERE block_number = $1 RETURNING *;

-- name: GetAllBlock :many
SELECT block_number FROM blocks
ORDER BY block_number ASC;