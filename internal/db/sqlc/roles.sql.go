// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: roles.sql

package db

import (
	"context"
)

const getRoleByID = `-- name: GetRoleByID :one
SELECT id, name, created_at, updated_at
FROM roles
WHERE id = $1
`

func (q *Queries) GetRoleByID(ctx context.Context, id int32) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRoleByID, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listRoles = `-- name: ListRoles :many
SELECT id, name, created_at, updated_at
FROM roles
ORDER BY name
`

func (q *Queries) ListRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.QueryContext(ctx, listRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
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
