-- +goose Up
CREATE TABLE ptacs (
    room TEXT PRIMARY KEY,
    brand TEXT NOT NULL,
    model TEXT NOT NULL,
    last_service TEXT NOT NULL
);

-- +goose Down
DROP TABLE ptacs;