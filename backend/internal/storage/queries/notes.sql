-- name: CreateNote :exec
INSERT INTO notes (id, path, content, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateNote :one
UPDATE notes
SET path = ?, content = ?, updated_at = ?
WHERE id = ?
RETURNING id, path, content, created_at, updated_at;

-- name: DeleteNote :execrows
DELETE FROM notes WHERE id = ?;
