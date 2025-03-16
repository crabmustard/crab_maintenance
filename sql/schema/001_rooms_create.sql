-- +goose Up
CREATE TABLE rooms (
    room INTEGER PRIMARY KEY,
    layout TEXT NOT NULL,
    guest TEXT,
    occupied INTEGER NOT NULL DEFAULT 0
);

-- +goose Down
DROP TABLE rooms;