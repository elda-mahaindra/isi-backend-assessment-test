-- name: CreateEntry :one
INSERT INTO entries (
    no_rekening,
    nominal,
    type_id
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetEntries :many
SELECT * FROM entries 
WHERE no_rekening = $1
ORDER BY created_at;