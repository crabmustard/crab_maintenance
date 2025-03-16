-- +goose Up
CREATE TABLE workorders (
    id UUID PRIMARY KEY,
    room INTEGER NOT NULL,
    is_closed INTEGER NOT NULL DEFAULT 0,
    code INTEGER NOT NULL,
    problem TEXT NOT NULL
);

-- +goose Down
DROP TABLE workorders;