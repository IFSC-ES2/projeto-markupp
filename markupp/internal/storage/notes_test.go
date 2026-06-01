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

	"github.com/ifsc-ES2/projeto-markupp/markupp/internal/notes"
	"github.com/ifsc-ES2/projeto-markupp/markupp/internal/storage"
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

func TestSqliteRepo_Update_AtualizaCamposEPreservaCreatedAt(t *testing.T) {
	db := setupTestDB(t)
	repo := storage.NewSqliteNotesRepository(db)
	original := sampleNote()
	require.NoError(t, repo.Save(context.Background(), original))

	got, err := repo.Update(context.Background(), original.ID, "renomeado.md", "novo conteudo", original.UpdatedAt, false)

	require.NoError(t, err)
	assert.Equal(t, original.ID, got.ID)
	assert.Equal(t, "renomeado.md", got.Path)
	assert.Equal(t, "novo conteudo", got.Content)
	assert.True(t, original.CreatedAt.Equal(got.CreatedAt))
	// UpdatedAt deve ter avançado
	assert.True(t, got.UpdatedAt.After(original.UpdatedAt))
}

func TestSqliteRepo_Update_IDInexistente_RetornaErrNotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := storage.NewSqliteNotesRepository(db)

	_, err := repo.Update(context.Background(), "nao-existe", "x.md", "y", time.Now(), false)

	require.Error(t, err)
	assert.True(t, errors.Is(err, notes.ErrNotFound))
}

func TestSqliteRepo_Update_PathDuplicado_RetornaErrDuplicatePath(t *testing.T) {
	db := setupTestDB(t)
	repo := storage.NewSqliteNotesRepository(db)
	n1 := sampleNote()
	n2 := sampleNote()
	n2.ID = "id-test-2"
	n2.Path = "outra.md"
	require.NoError(t, repo.Save(context.Background(), n1))
	require.NoError(t, repo.Save(context.Background(), n2))

	_, err := repo.Update(context.Background(), n2.ID, n1.Path, n2.Content, time.Now(), false)

	require.Error(t, err)
	assert.True(t, errors.Is(err, notes.ErrDuplicatePath))
}

func TestSqliteRepo_Delete_RemoveLinha(t *testing.T) {
	db := setupTestDB(t)
	repo := storage.NewSqliteNotesRepository(db)
	n := sampleNote()
	require.NoError(t, repo.Save(context.Background(), n))

	err := repo.Delete(context.Background(), n.ID)
	require.NoError(t, err)

	var count int
	require.NoError(t, db.QueryRow("SELECT COUNT(*) FROM notes WHERE id = ?", n.ID).Scan(&count))
	assert.Equal(t, 0, count)
}

func TestSqliteRepo_Delete_IDInexistente_RetornaErrNotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := storage.NewSqliteNotesRepository(db)

	err := repo.Delete(context.Background(), "nao-existe")

	require.Error(t, err)
	assert.True(t, errors.Is(err, notes.ErrNotFound))
}

func TestSqliteRepo_ListNotes_DBVazio_RetornaSliceVazio(t *testing.T) {
	db := setupTestDB(t)
	repo := storage.NewSqliteNotesRepository(db)

	got, err := repo.ListNotes(context.Background())

	require.NoError(t, err)
	assert.Empty(t, got)
}

func TestSqliteRepo_ListNotes_RetornaTodasNotasOrdenadasPorPath(t *testing.T) {
	db := setupTestDB(t)
	repo := storage.NewSqliteNotesRepository(db)
	now := time.Date(2026, 4, 27, 10, 0, 0, 0, time.UTC)

	notas := []notes.Note{
		{ID: "id-c", Path: "c.md", Content: "ccc", CreatedAt: now, UpdatedAt: now},
		{ID: "id-a", Path: "a.md", Content: "aaa", CreatedAt: now, UpdatedAt: now},
		{ID: "id-b", Path: "b.md", Content: "bbb", CreatedAt: now, UpdatedAt: now},
	}
	for _, n := range notas {
		require.NoError(t, repo.Save(context.Background(), n))
	}

	got, err := repo.ListNotes(context.Background())

	require.NoError(t, err)
	require.Len(t, got, 3)
	assert.Equal(t, "a.md", got[0].Path)
	assert.Equal(t, "b.md", got[1].Path)
	assert.Equal(t, "c.md", got[2].Path)
	assert.Equal(t, "aaa", got[0].Content)
}
