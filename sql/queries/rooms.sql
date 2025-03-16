-- name: GetRoom :one
SELECT * FROM rooms
WHERE room = ?;

-- name: GetAllRooms :many
SELECT * FROM rooms;


-- name: CreateRoom :one
INSERT INTO rooms (
  room, layout, guest, occupied
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: AddRoomGuest :exec
UPDATE rooms
SET guest = ?,
occupied = 1
WHERE room = ?;

-- name: RemoveRoomGuest :exec
UPDATE rooms
SET guest = ?,
occupied = 0
WHERE room = ?;