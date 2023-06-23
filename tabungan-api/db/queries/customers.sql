  -- name: CreateCustomer :one
INSERT INTO customers (
    nama,
    nik,
    no_hp
) VALUES (
    $1, $2, $3
) RETURNING *;