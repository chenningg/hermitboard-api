-- name: CreateAccountAuthRole :one
INSERT INTO account_auth_role (
    id, account_id, auth_role_id
) VALUES (
    @id, @account_id, @auth_role_id
)
RETURNING *;

-- name: GetAccountAuthRole :one
SELECT * FROM account_auth_role
WHERE id = @id
LIMIT 1;

-- name: ListAccountAuthRoles :many
SELECT * FROM account_auth_role
ORDER BY account_id, auth_role_id;

-- name: UpdateAccountAuthRoles :one
UPDATE account_auth_role
SET 
    id = CASE WHEN @id_do_update::BOOLEAN
        THEN @new_id::TEXT ELSE id END,
    account_id = CASE WHEN @account_id_do_update::BOOLEAN
        THEN @account_id::uuid ELSE account_id END,
    auth_role_id = CASE WHEN @auth_role_id_do_update::BOOLEAN
        THEN @auth_role_id::TEXT ELSE account_id END,
WHERE id = @id
RETURNING *;

-- name: DeleteAuthRole :one
DELETE FROM auth_role
WHERE id = @id
RETURNING *;