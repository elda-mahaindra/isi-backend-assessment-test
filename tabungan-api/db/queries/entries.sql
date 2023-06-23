-- name: CreateEntry :one
INSERT INTO entries (
    code,
    nominal,
    no_rekening
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetEntries :many
SELECT * FROM entries 
WHERE no_rekening = $1
ORDER BY created_at;