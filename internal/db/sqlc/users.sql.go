// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
        first_name,
        last_name,
        email,
        phone_number,
        role_id
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING id, first_name, last_name, email, phone_number, password, role_id, created_at, updated_at
`

type CreateUserParams struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	RoleID      int32
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.PhoneNumber,
		arg.RoleID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.Password,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, first_name, last_name, email, phone_number, password, role_id, created_at, updated_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.Password,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsersWithPetsPaginated = `-- name: GetUsersWithPetsPaginated :many
SELECT u.id AS user_id,
    u.first_name,
    u.last_name,
    u.email,
    u.phone_number,
    p.id AS pet_id,
    p.name AS pet_name,
    p.age AS pet_age,
    s.name AS species_name,
    b.name AS breed_name
FROM users u
    LEFT JOIN pets p ON p.owner_id = u.id
    LEFT JOIN species s ON p.species_id = s.id
    LEFT JOIN breeds b ON p.breed_id = b.id
ORDER BY u.created_at DESC
LIMIT $1 OFFSET $2
`

type GetUsersWithPetsPaginatedParams struct {
	Limit  int32
	Offset int32
}

type GetUsersWithPetsPaginatedRow struct {
	UserID      int32
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	PetID       sql.NullInt32
	PetName     sql.NullString
	PetAge      sql.NullTime
	SpeciesName sql.NullString
	BreedName   sql.NullString
}

func (q *Queries) GetUsersWithPetsPaginated(ctx context.Context, arg GetUsersWithPetsPaginatedParams) ([]GetUsersWithPetsPaginatedRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsersWithPetsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersWithPetsPaginatedRow
	for rows.Next() {
		var i GetUsersWithPetsPaginatedRow
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.PhoneNumber,
			&i.PetID,
			&i.PetName,
			&i.PetAge,
			&i.SpeciesName,
			&i.BreedName,
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

const updateUserByID = `-- name: UpdateUserByID :one
update users
SET first_name = $1,
    last_name = $2,
    email = $3,
    phone_number = $4,
    role_id = $5
WHERE id = $6
RETURNING id, first_name, last_name, email, phone_number, password, role_id, created_at, updated_at
`

type UpdateUserByIDParams struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	RoleID      int32
	ID          int32
}

func (q *Queries) UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserByID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.PhoneNumber,
		arg.RoleID,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.Password,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
