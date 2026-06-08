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
	q     *gen.Queries
	clock func() time.Time
}

func NewSqliteNotesRepository(db *sql.DB) *SqliteNotesRepository {
	return &SqliteNotesRepository{
		q:     gen.New(db),
		clock: time.Now,
	}
}

func NewSqliteNotesRepositoryWithClock(db *sql.DB, clock func() time.Time) *SqliteNotesRepository {
	return &SqliteNotesRepository{
		q:     gen.New(db),
		clock: clock,
	}
}

func (r *SqliteNotesRepository) Save(ctx context.Context, note notes.Note) error {
	note.CreatedAt = note.CreatedAt.UTC().Truncate(time.Millisecond)
	note.UpdatedAt = note.UpdatedAt.UTC().Truncate(time.Millisecond)

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
	now := r.clock()
	var row gen.Note
	var err error

	lastModifiedAt = lastModifiedAt.UTC().Truncate(time.Millisecond)
	now = now.UTC().Truncate(time.Millisecond)

	if force {
		row, err = r.q.UpdateNoteForced(ctx, gen.UpdateNoteForcedParams{
			ID:        id,
			Path:      path,
			Content:   content,
			UpdatedAt: now,
		})
	} else {
		row, err = r.q.UpdateNoteWithVersionCheck(ctx, gen.UpdateNoteWithVersionCheckParams{
			ID:            id,
			Path:          path,
			Content:       content,
			UpdatedAt:     now,
			PrevUpdatedAt: lastModifiedAt,
		})
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if !force {
				_, checkErr := r.q.GetNoteByID(ctx, id)
				if errors.Is(checkErr, sql.ErrNoRows) {
					return notes.Note{}, notes.ErrNotFound
				}
				return notes.Note{}, notes.ErrConflict
			}
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
		CreatedAt: row.CreatedAt.UTC().Truncate(time.Millisecond),
		UpdatedAt: row.UpdatedAt.UTC().Truncate(time.Millisecond),
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
		if errors.Is(err, sql.ErrNoRows) {
			return notes.Note{}, notes.ErrNotFound
		}
		return notes.Note{}, err
	}
	return notes.Note{
		ID:        row.ID,
		Path:      row.Path,
		Content:   row.Content,
		CreatedAt: row.CreatedAt.UTC().Truncate(time.Millisecond),
		UpdatedAt: row.UpdatedAt.UTC().Truncate(time.Millisecond),
	}, nil
}

func (r *SqliteNotesRepository) SearchNotes(ctx context.Context, query string, offset, limit int32) ([]notes.SearchResult, error) {
	rows, err := r.q.SearchNotes(ctx, gen.SearchNotesParams{
		Content: "%" + query + "%",
		Limit:   int64(limit),
		Offset:  int64(offset),
	})
	if err != nil {
		return nil, err
	}
	out := make([]notes.SearchResult, 0, len(rows))
	for _, row := range rows {
		out = append(out, notes.SearchResult{
			ID:        row.ID,
			Path:      row.Path,
			UpdatedAt: row.UpdatedAt,
		})
	}
	return out, nil
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
			CreatedAt: row.CreatedAt.UTC().Truncate(time.Millisecond),
			UpdatedAt: row.UpdatedAt.UTC().Truncate(time.Millisecond),
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
