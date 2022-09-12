-- name: CreateAuthRole :one
INSERT INTO auth_role (
    id, description
) VALUES (
    @id, @description
)
RETURNING *;

-- name: GetAuthRole :one
SELECT * FROM auth_role
WHERE id = @id
LIMIT 1;

-- name: ListAuthRoles :many
SELECT * FROM auth_role
ORDER BY id;

-- name: UpdateAuthRole :one
UPDATE auth_role
SET 
    id = CASE WHEN @id_do_update::BOOLEAN
        THEN @new_id::TEXT ELSE id END,
    description = CASE WHEN @description_do_update::BOOLEAN
        THEN @description::TEXT ELSE description END,
WHERE id = @id
RETURNING *;

-- name: DeleteAuthRole :one
DELETE FROM auth_role
WHERE id = @id
RETURNING *;