-- name: GetAccount :one
SELECT * FROM authors
WHERE id = ? LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAccount :execresult
INSERT INTO authors (
  name, bio
) VALUES (
  ?, ?
);

-- name: DeleteAccount :exec
DELETE FROM authors
WHERE id = ?;