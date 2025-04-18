-- name: CreatePtac :one
INSERT INTO ptacs (room, brand, model, last_service)
VALUES (
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: GetPtacRoom :one
SELECT * FROM ptacs WHERE room = ?;


-- name: GetPtacsToClean :many
SELECT * FROM ptacs
ORDER BY last_service ASC
LIMIT ?;

-- name: GetPtacCount :one
SELECT COUNT(room) FROM ptacs;

-- name: GetAllPtac :many
SELECT * FROM ptacs
ORDER BY room;


-- name: ClearPtacList :exec
DELETE FROM ptacs;