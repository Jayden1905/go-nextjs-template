// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (first_name, last_name, email, password)
VALUES (?, ?, ?, ?)
`

type CreateUserParams struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
	)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, password, createdat, updatedat FROM users WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, first_name, last_name, email, password, createdat, updatedat FROM users WHERE id = ?
`

func (q *Queries) GetUserByID(ctx context.Context, id uint32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}
