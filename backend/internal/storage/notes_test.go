package storage_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	_ "modernc.org/sqlite"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ifsc-ES2/projeto-markupp/backend/internal/notes"
	"github.com/ifsc-ES2/projeto-markupp/backend/internal/storage"
)

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })
	require.NoError(t, storage.Migrate(db))
	return db
}

func sampleNote() notes.Note {
	now := time.Date(2026, 4, 27, 10, 0, 0, 0, time.UTC)
	return notes.Note{
		ID:        "id-test-1",
		Path:      "amostra.md",
		Content:   "conteudo de amostra",
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func TestSqliteRepo_Save_PersisteCamposCorretos(t *testing.T) {
	db := setupTestDB(t)
	repo := storage.NewSqliteNotesRepository(db)
	n := sampleNote()

	err := repo.Save(context.Background(), n)
	require.NoError(t, err)

	var gotID, gotPath, gotContent string
	var gotCreated, gotUpdated time.Time
	err = db.QueryRowContext(context.Background(),
		"SELECT id, path, content, created_at, updated_at FROM notes WHERE id = ?",
		n.ID,
	).Scan(&gotID, &gotPath, &gotContent, &gotCreated, &gotUpdated)
	require.NoError(t, err)

	assert.Equal(t, n.ID, gotID)
	assert.Equal(t, n.Path, gotPath)
	assert.Equal(t, n.Content, gotContent)
	assert.True(t, n.CreatedAt.Equal(gotCreated))
	assert.True(t, n.UpdatedAt.Equal(gotUpdated))
}

func TestSqliteRepo_Save_PathDuplicado_RetornaErrDuplicatePath(t *testing.T) {
	db := setupTestDB(t)
	repo := storage.NewSqliteNotesRepository(db)
	n1 := sampleNote()
	n2 := sampleNote()
	n2.ID = "id-test-2"

	require.NoError(t, repo.Save(context.Background(), n1))

	err := repo.Save(context.Background(), n2)
	require.Error(t, err)
	assert.True(t, errors.Is(err, notes.ErrDuplicatePath))
}

func TestSqliteRepo_Save_Sucesso_RetornaNil(t *testing.T) {
	db := setupTestDB(t)
	repo := storage.NewSqliteNotesRepository(db)

	err := repo.Save(context.Background(), sampleNote())
	assert.NoError(t, err)
}
