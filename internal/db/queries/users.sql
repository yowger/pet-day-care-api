-- name: CreateUser :one
INSERT INTO users (
        first_name,
        last_name,
        email,
        phone_number,
        password,
        role_id
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetUsersWithPetsPaginated :many
SELECT u.id AS user_id,
    u.first_name,
    u.last_name,
    u.email,
    u.phone_number,
    p.id AS pet_id,
    p.name AS pet_name,
    p.birth_date AS pet_birth_date,
    s.name AS species_name,
    b.name AS breed_name
FROM users u
    LEFT JOIN pets p ON p.owner_id = u.id
    LEFT JOIN species s ON p.species_id = s.id
    LEFT JOIN breeds b ON p.breed_id = b.id
ORDER BY u.created_at DESC
LIMIT $1 OFFSET $2;
-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;
-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;
-- name: UpdateUserByID :one
update users
SET first_name = $1,
    last_name = $2,
    email = $3,
    phone_number = $4,
    role_id = $5
WHERE id = $6
RETURNING *;
