-- name: ListRoles :many
SELECT *
FROM roles
ORDER BY name;
-- name: GetRoleByID :one
SELECT *
FROM roles
WHERE id = $1;