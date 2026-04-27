-- name: CreateNote :exec
INSERT INTO notes (id, path, content, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);
