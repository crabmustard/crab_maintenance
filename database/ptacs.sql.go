// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: ptacs.sql

package database

import (
	"context"
)

const clearPtacList = `-- name: ClearPtacList :exec
DELETE FROM ptacs
`

func (q *Queries) ClearPtacList(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, clearPtacList)
	return err
}

const createPtac = `-- name: CreatePtac :one
INSERT INTO ptacs (room, brand, model, last_service)
VALUES (
    ?,
    ?,
    ?,
    ?
)
RETURNING room, brand, model, last_service
`

type CreatePtacParams struct {
	Room        int64  `json:"room"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	LastService string `json:"last_service"`
}

func (q *Queries) CreatePtac(ctx context.Context, arg CreatePtacParams) (Ptac, error) {
	row := q.db.QueryRowContext(ctx, createPtac,
		arg.Room,
		arg.Brand,
		arg.Model,
		arg.LastService,
	)
	var i Ptac
	err := row.Scan(
		&i.Room,
		&i.Brand,
		&i.Model,
		&i.LastService,
	)
	return i, err
}

const getAllPtac = `-- name: GetAllPtac :many
SELECT room, brand, model, last_service FROM ptacs
`

func (q *Queries) GetAllPtac(ctx context.Context) ([]Ptac, error) {
	rows, err := q.db.QueryContext(ctx, getAllPtac)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ptac
	for rows.Next() {
		var i Ptac
		if err := rows.Scan(
			&i.Room,
			&i.Brand,
			&i.Model,
			&i.LastService,
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

const getPtac = `-- name: GetPtac :one
SELECT room, brand, model, last_service FROM ptacs WHERE room = ?
`

func (q *Queries) GetPtac(ctx context.Context, room int64) (Ptac, error) {
	row := q.db.QueryRowContext(ctx, getPtac, room)
	var i Ptac
	err := row.Scan(
		&i.Room,
		&i.Brand,
		&i.Model,
		&i.LastService,
	)
	return i, err
}

const getPtacCount = `-- name: GetPtacCount :one
SELECT COUNT(room) FROM ptacs
`

func (q *Queries) GetPtacCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getPtacCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}
