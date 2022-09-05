-- name: CreateAccount :one
INSERT INTO account (
    id, auth_id, nickname, email, password
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1
LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY email;

-- name: UpdateAccount :one
UPDATE account
SET auth_id = $2,
nickname = $3,
email = $4,
password = $5
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :one
DELETE FROM account
WHERE id = $1
RETURNING *;