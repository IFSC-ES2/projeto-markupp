package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ifsc-ES2/projeto-markupp/backend/internal/notes"
)

type SqliteNotesRepository struct{}

func NewSqliteNotesRepository(db *sql.DB) *SqliteNotesRepository {
	return &SqliteNotesRepository{}
}

func (r *SqliteNotesRepository) Save(ctx context.Context, note notes.Note) error {
	return errors.New("não implementado")
}
