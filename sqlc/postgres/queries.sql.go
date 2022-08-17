// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: queries.sql

package postgres

import (
	"context"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (user_id, task, done)
VALUES ($1, $2, $3) RETURNING id, user_id, task, done
`

type CreateTodoParams struct {
	UserID int32  `json:"user_id"`
	Task   string `json:"task"`
	Done   bool   `json:"done"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, arg.UserID, arg.Task, arg.Done)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Task,
		&i.Done,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (firstname, lastname)
VALUES ($1, $2) RETURNING id, firstname, lastname
`

type CreateUserParams struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Firstname, arg.Lastname)
	var i User
	err := row.Scan(&i.ID, &i.Firstname, &i.Lastname)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, firstname, lastname FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Firstname, &i.Lastname)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, firstname, lastname FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Firstname, &i.Lastname); err != nil {
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
