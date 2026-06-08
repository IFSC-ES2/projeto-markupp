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
SET path = sqlc.arg(path), content = sqlc.arg(content), updated_at = sqlc.arg(updated_at)
WHERE id = sqlc.arg(id) AND updated_at = sqlc.arg(prev_updated_at)
RETURNING id, path, content, created_at, updated_at;

-- name: UpdateNoteForced :one
UPDATE notes
SET path = ?, content = ?, updated_at = ?
WHERE id = ?
RETURNING id, path, content, created_at, updated_at;

-- name: DeleteNote :execrows
DELETE FROM notes WHERE id = ?;

-- name: SearchNotes :many
SELECT id, path, updated_at FROM notes
WHERE content LIKE ?
ORDER BY updated_at DESC
LIMIT ? OFFSET ?;
