-- +goose Up
CREATE TABLE ptacs (
    room INTEGER PRIMARY KEY,
    brand TEXT NOT NULL,
    model TEXT NOT NULL,
    last_service INTEGER NOT NULL
);

-- +goose Down
DROP TABLE ptacs;