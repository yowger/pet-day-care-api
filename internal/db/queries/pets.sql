-- name: CreatePet :one
INSERT INTO pets (name, species_id, breed_id)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetPetsWithOwnersPaginated :many
SELECT p.id AS pet_id,
    p.name AS pet_name,
    p.age AS pet_age,
    s.name AS species_name,
    b.name AS breed_name,
    u.id AS owner_id,
    CONCAT(u.first_name, ' ', u.last_name) AS owner_name,
    u.email AS owner_email,
    u.phone_number AS owner_phone
FROM pets p
    LEFT JOIN species s ON p.species_id = s.id
    LEFT JOIN breeds b ON p.breed_id = b.id
    LEFT JOIN users u ON p.owner_id = u.id
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2;
;
-- name: GetPetByID :one
SELECT p.*,
    b.name AS breed_name,
    s.name AS species_name
FROM pets p
    LEFT JOIN breeds b ON p.breed_id = b.id
    LEFT JOIN species s ON p.species_id = s.id
WHERE p.id = $1;
-- name: UpdatePet :one
UPDATE pets
SET name = $1,
    species_id = $2,
    breed_id = $3,
    updated_at = NOW()
WHERE id = $4
RETURNING *;