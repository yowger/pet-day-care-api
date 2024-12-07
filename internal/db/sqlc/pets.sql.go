// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: pets.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createPet = `-- name: CreatePet :one
INSERT INTO pets (name, species_id, breed_id)
VALUES ($1, $2, $3)
RETURNING id, age, name, species_id, breed_id, created_at, updated_at
`

type CreatePetParams struct {
	Name      string
	SpeciesID int32
	BreedID   int32
}

func (q *Queries) CreatePet(ctx context.Context, arg CreatePetParams) (Pet, error) {
	row := q.db.QueryRowContext(ctx, createPet, arg.Name, arg.SpeciesID, arg.BreedID)
	var i Pet
	err := row.Scan(
		&i.ID,
		&i.Age,
		&i.Name,
		&i.SpeciesID,
		&i.BreedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPetByID = `-- name: GetPetByID :one
SELECT p.id, p.age, p.name, p.species_id, p.breed_id, p.created_at, p.updated_at,
    b.name AS breed_name,
    s.name AS species_name
FROM pets p
    LEFT JOIN breeds b ON p.breed_id = b.id
    LEFT JOIN species s ON p.species_id = s.id
WHERE p.id = $1
`

type GetPetByIDRow struct {
	ID          int32
	Age         time.Time
	Name        string
	SpeciesID   int32
	BreedID     int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BreedName   sql.NullString
	SpeciesName sql.NullString
}

func (q *Queries) GetPetByID(ctx context.Context, id int32) (GetPetByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getPetByID, id)
	var i GetPetByIDRow
	err := row.Scan(
		&i.ID,
		&i.Age,
		&i.Name,
		&i.SpeciesID,
		&i.BreedID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BreedName,
		&i.SpeciesName,
	)
	return i, err
}

const getPetsWithOwnersPaginated = `-- name: GetPetsWithOwnersPaginated :many
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
LIMIT $1 OFFSET $2
`

type GetPetsWithOwnersPaginatedParams struct {
	Limit  int32
	Offset int32
}

type GetPetsWithOwnersPaginatedRow struct {
	PetID       int32
	PetName     string
	PetAge      time.Time
	SpeciesName sql.NullString
	BreedName   sql.NullString
	OwnerID     sql.NullInt32
	OwnerName   interface{}
	OwnerEmail  sql.NullString
	OwnerPhone  sql.NullString
}

func (q *Queries) GetPetsWithOwnersPaginated(ctx context.Context, arg GetPetsWithOwnersPaginatedParams) ([]GetPetsWithOwnersPaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, getPetsWithOwnersPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPetsWithOwnersPaginatedRow
	for rows.Next() {
		var i GetPetsWithOwnersPaginatedRow
		if err := rows.Scan(
			&i.PetID,
			&i.PetName,
			&i.PetAge,
			&i.SpeciesName,
			&i.BreedName,
			&i.OwnerID,
			&i.OwnerName,
			&i.OwnerEmail,
			&i.OwnerPhone,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePet = `-- name: UpdatePet :one
UPDATE pets
SET name = $1,
    species_id = $2,
    breed_id = $3,
    updated_at = NOW()
WHERE id = $4
RETURNING id, age, name, species_id, breed_id, created_at, updated_at
`

type UpdatePetParams struct {
	Name      string
	SpeciesID int32
	BreedID   int32
	ID        int32
}

func (q *Queries) UpdatePet(ctx context.Context, arg UpdatePetParams) (Pet, error) {
	row := q.db.QueryRowContext(ctx, updatePet,
		arg.Name,
		arg.SpeciesID,
		arg.BreedID,
		arg.ID,
	)
	var i Pet
	err := row.Scan(
		&i.ID,
		&i.Age,
		&i.Name,
		&i.SpeciesID,
		&i.BreedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
