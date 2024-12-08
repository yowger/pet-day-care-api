// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Breed struct {
	ID        int32            `db:"id" json:"id"`
	Name      string           `db:"name" json:"name"`
	SpeciesID int32            `db:"species_id" json:"species_id"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type Pet struct {
	ID        int32            `db:"id" json:"id"`
	Age       pgtype.Timestamp `db:"age" json:"age"`
	Name      string           `db:"name" json:"name"`
	SpeciesID int32            `db:"species_id" json:"species_id"`
	BreedID   int32            `db:"breed_id" json:"breed_id"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type Role struct {
	ID        int32            `db:"id" json:"id"`
	Name      string           `db:"name" json:"name"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type Species struct {
	ID        int32            `db:"id" json:"id"`
	Name      string           `db:"name" json:"name"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type User struct {
	ID          int32            `db:"id" json:"id"`
	FirstName   string           `db:"first_name" json:"first_name"`
	LastName    string           `db:"last_name" json:"last_name"`
	Email       string           `db:"email" json:"email"`
	PhoneNumber string           `db:"phone_number" json:"phone_number"`
	Password    string           `db:"password" json:"password"`
	RoleID      int32            `db:"role_id" json:"role_id"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
