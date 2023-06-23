-- name: CreateAccount :one
INSERT INTO accounts (
    customer_id,
    no_rekening,
    saldo
) VALUES (
    $1, $2, 0
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE no_rekening = $1;

-- name: UpdateSaldo :one
UPDATE accounts
SET saldo = $2
WHERE no_rekening = $1
RETURNING *;