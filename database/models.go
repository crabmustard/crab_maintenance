// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

type Ptac struct {
	Room        int64  `json:"room"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	LastService int64  `json:"last_service"`
}
