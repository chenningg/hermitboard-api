-- name: CreateAccount :one
INSERT INTO account (
    id, auth_id, nickname, email, password
) VALUES (
    @id, @auth_id, @nickname, @email, @password
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = @id
LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY email;

-- name: UpdateAccount :one
UPDATE account
SET 
    auth_id = CASE WHEN @auth_id_do_update::boolean
        THEN @auth_id::TEXT ELSE auth_id END,
    nickname = CASE WHEN @nickname_do_update::boolean
        THEN @nickname::TEXT ELSE nickname END,
    email = CASE WHEN @email_do_update::boolean
        THEN @email::TEXT ELSE email END,
    password = CASE WHEN @password_do_update::boolean
        THEN @password::TEXT ELSE password END
WHERE id = @id
RETURNING *;

-- name: DeleteAccount :one
DELETE FROM account
WHERE id = @id
RETURNING *;