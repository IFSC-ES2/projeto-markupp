package storage

import (
	"context"
	"database/sql"
	"errors"

	sqlite "modernc.org/sqlite"

	"github.com/ifsc-ES2/projeto-markupp/backend/internal/notes"
	"github.com/ifsc-ES2/projeto-markupp/backend/internal/storage/gen"
)

const sqliteUniqueConstraintCode = 2067

type SqliteNotesRepository struct {
	q *gen.Queries
}

func NewSqliteNotesRepository(db *sql.DB) *SqliteNotesRepository {
	return &SqliteNotesRepository{q: gen.New(db)}
}

func (r *SqliteNotesRepository) Save(ctx context.Context, note notes.Note) error {
	err := r.q.CreateNote(ctx, gen.CreateNoteParams{
		ID:        note.ID,
		Path:      note.Path,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	})
	if err == nil {
		return nil
	}
	if isUniqueConstraintViolation(err) {
		return notes.ErrDuplicatePath
	}
	return err
}

func isUniqueConstraintViolation(err error) bool {
	var sqliteErr *sqlite.Error
	if !errors.As(err, &sqliteErr) {
		return false
	}
	return sqliteErr.Code() == sqliteUniqueConstraintCode
}
