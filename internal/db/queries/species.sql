-- name: CreateSpecies :one
INSERT INTO species (name)
VALUES ($1)
RETURNING *;
-- name: GetSpeciesPaginated :many
SELECT *
FROM species
ORDER BY name
LIMIT $1 OFFSET $2;
-- name: GetSpecieByID :one
SELECT *
FROM species
WHERE id = $1;
-- name: UpdateSpeciesByID :one
UPDATE species
SET name = $1,
    updated_at = NOW()
WHERE id = $2
RETURNING *;