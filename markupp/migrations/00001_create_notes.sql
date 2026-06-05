-- +goose Up
CREATE TABLE notes (
    id          TEXT PRIMARY KEY,
    path        TEXT NOT NULL UNIQUE,
    content     TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS notes;
