// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
)

type Room struct {
	Room     int64          `json:"room"`
	Layout   string         `json:"layout"`
	Guest    sql.NullString `json:"guest"`
	Occupied int64          `json:"occupied"`
}

type Workorder struct {
	ID       interface{} `json:"id"`
	Room     int64       `json:"room"`
	IsClosed int64       `json:"is_closed"`
	Code     int64       `json:"code"`
	Problem  string      `json:"problem"`
}
