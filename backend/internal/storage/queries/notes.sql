-- name: CreateNote :exec
INSERT INTO notes (id, path, content, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: GetNoteByID :one
SELECT id, path, content, created_at, updated_at FROM notes WHERE id = ?;