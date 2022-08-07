-- name: GetAccount :one
SELECT * FROM account
WHERE id = ? LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY email;

-- name: CreateAccount :execresult
INSERT INTO account (
  id, email, nickname
) VALUES (
  ?, ?, ?
);

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = ?;