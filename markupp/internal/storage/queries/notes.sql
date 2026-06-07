-- name: CreateNote :exec
INSERT INTO notes (id, path, content, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: GetNoteByID :one
SELECT id, path, content, created_at, updated_at FROM notes WHERE id = ?;

-- name: ListNotes :many
SELECT id, path, content, created_at, updated_at FROM notes
ORDER BY path;

-- name: UpdateNoteWithVersionCheck :one
UPDATE notes
SET path = ?, content = ?, updated_at = ?
WHERE id = ? AND updated_at = ?
RETURNING id, path, content, created_at, updated_at;

-- name: UpdateNoteForced :one
UPDATE notes
SET path = ?, content = ?, updated_at = ?
WHERE id = ?
RETURNING id, path, content, created_at, updated_at;

-- name: DeleteNote :execrows
DELETE FROM notes WHERE id = ?;
