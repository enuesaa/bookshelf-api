// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: dbquery.sql

package dbq

import (
	"context"
)

const createTea = `-- name: CreateTea :one
INSERT INTO teas (
  teapod, teabox, teaid, value
) VALUES (
  ?, ?, ?, ?
)
RETURNING id, teapod, teabox, teaid, value, created, updated
`

type CreateTeaParams struct {
	Teapod string
	Teabox string
	Teaid  string
	Value  interface{}
}

func (q *Queries) CreateTea(ctx context.Context, arg CreateTeaParams) (Tea, error) {
	row := q.db.QueryRowContext(ctx, createTea,
		arg.Teapod,
		arg.Teabox,
		arg.Teaid,
		arg.Value,
	)
	var i Tea
	err := row.Scan(
		&i.ID,
		&i.Teapod,
		&i.Teabox,
		&i.Teaid,
		&i.Value,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const deleteTea = `-- name: DeleteTea :exec
DELETE FROM teas
WHERE teapod = ? and teaid = ?
`

type DeleteTeaParams struct {
	Teapod string
	Teaid  string
}

func (q *Queries) DeleteTea(ctx context.Context, arg DeleteTeaParams) error {
	_, err := q.db.ExecContext(ctx, deleteTea, arg.Teapod, arg.Teaid)
	return err
}

const getTea = `-- name: GetTea :one
SELECT id, teapod, teabox, teaid, value, created, updated FROM teas
WHERE teapod = ? and teaid = ? LIMIT 1
`

type GetTeaParams struct {
	Teapod string
	Teaid  string
}

func (q *Queries) GetTea(ctx context.Context, arg GetTeaParams) (Tea, error) {
	row := q.db.QueryRowContext(ctx, getTea, arg.Teapod, arg.Teaid)
	var i Tea
	err := row.Scan(
		&i.ID,
		&i.Teapod,
		&i.Teabox,
		&i.Teaid,
		&i.Value,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listTeas = `-- name: ListTeas :many
SELECT id, teapod, teabox, teaid, value, created, updated FROM teas
WHERE teapod = ?
`

func (q *Queries) ListTeas(ctx context.Context, teapod string) ([]Tea, error) {
	rows, err := q.db.QueryContext(ctx, listTeas, teapod)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tea
	for rows.Next() {
		var i Tea
		if err := rows.Scan(
			&i.ID,
			&i.Teapod,
			&i.Teabox,
			&i.Teaid,
			&i.Value,
			&i.Created,
			&i.Updated,
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

const listTeasByTeaboxName = `-- name: ListTeasByTeaboxName :many
SELECT id, teapod, teabox, teaid, value, created, updated FROM teas
WHERE teapod = ? and teabox = ?
`

type ListTeasByTeaboxNameParams struct {
	Teapod string
	Teabox string
}

func (q *Queries) ListTeasByTeaboxName(ctx context.Context, arg ListTeasByTeaboxNameParams) ([]Tea, error) {
	rows, err := q.db.QueryContext(ctx, listTeasByTeaboxName, arg.Teapod, arg.Teabox)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tea
	for rows.Next() {
		var i Tea
		if err := rows.Scan(
			&i.ID,
			&i.Teapod,
			&i.Teabox,
			&i.Teaid,
			&i.Value,
			&i.Created,
			&i.Updated,
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