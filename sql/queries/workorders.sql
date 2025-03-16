-- name: CreateWorkOrder :one
INSERT INTO workorders (
  id, room, is_closed, code, problem
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: CloseWorkOrder :exec
UPDATE workorders
SET is_closed = 1
WHERE id = ?;

-- name: GetWorkOrder :one
SELECT * FROM workorders
WHERE id = ?;

-- name: GetAllWorkOrders :many
SELECT * FROM workorders;