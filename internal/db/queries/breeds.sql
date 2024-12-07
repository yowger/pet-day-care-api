-- name: CreateBreed :one
INSERT INTO breeds (name, species_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetAllBreedsPaginated :many
SELECT *
FROM breeds
ORDER BY name
LIMIT $1 OFFSET $2;
-- name: GetBreedByID :one
SELECT *
FROM breeds
WHERE id = $1;
-- name: UpdateBreed :one
UPDATE breeds
SET name = $1,
    species_id = $2,
    updated_at = NOW()
WHERE id = $3
RETURNING *;