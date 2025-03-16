-- name: CreatePtac :one
INSERT INTO ptacs (room, brand, model, last_service)
VALUES (
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: GetPtac :one
SELECT * FROM ptacs WHERE room = ?;

-- name: GetPtacCount :one
SELECT COUNT(room) FROM ptacs;

-- name: GetAllPtac :many
SELECT * FROM ptacs;


-- name: ClearPtacList :exec
DELETE FROM ptacs;