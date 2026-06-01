package storage

import (
	"context"
	"database/sql"
	"errors"
	"time"

	sqlite "modernc.org/sqlite"

	"github.com/ifsc-ES2/projeto-markupp/markupp/internal/notes"
	"github.com/ifsc-ES2/projeto-markupp/markupp/internal/storage/gen"
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

func (r *SqliteNotesRepository) Update(ctx context.Context, id, path, content string, lastModifiedAt time.Time, force bool) (notes.Note, error) {
	row, err := r.q.UpdateNote(ctx, gen.UpdateNoteParams{
		ID:        id,
		Path:      path,
		Content:   content,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return notes.Note{}, notes.ErrNotFound
		}
		if isUniqueConstraintViolation(err) {
			return notes.Note{}, notes.ErrDuplicatePath
		}
		return notes.Note{}, err
	}

	return notes.Note{
		ID:        row.ID,
		Path:      row.Path,
		Content:   row.Content,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}

func (r *SqliteNotesRepository) Delete(ctx context.Context, id string) error {
	rows, err := r.q.DeleteNote(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return notes.ErrNotFound
	}
	return nil
}

func (r *SqliteNotesRepository) GetNoteByID(ctx context.Context, id string) (notes.Note, error) {
	row, err := r.q.GetNoteByID(ctx, id)
	if err != nil {
		return notes.Note{}, err
	}
	return notes.Note{
		ID:        row.ID,
		Path:      row.Path,
		Content:   row.Content,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}

func (r *SqliteNotesRepository) ListNotes(ctx context.Context) ([]notes.Note, error) {
	rows, err := r.q.ListNotes(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]notes.Note, 0, len(rows))
	for _, row := range rows {
		out = append(out, notes.Note{
			ID:        row.ID,
			Path:      row.Path,
			Content:   row.Content,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		})
	}
	return out, nil
}

func isUniqueConstraintViolation(err error) bool {
	var sqliteErr *sqlite.Error
	if !errors.As(err, &sqliteErr) {
		return false
	}
	return sqliteErr.Code() == sqliteUniqueConstraintCode
}
