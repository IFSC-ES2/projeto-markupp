package storage_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ifsc-ES2/projeto-markupp/markupp/internal/notes"
	"github.com/ifsc-ES2/projeto-markupp/markupp/internal/storage"
	_ "modernc.org/sqlite"
)

func TestSearchNotes_ComResultados_RetornaPaginado(t *testing.T) {
	db := setupIntegrationTestDB(t)
	defer db.Close()
	repo := storage.NewSqliteNotesRepository(db)
	ctx := context.Background()

	// Inserir dados de teste
	now := time.Now()
	notesData := []notes.Note{
		{ID: "1", Path: "golang1.md", Content: "golang tutorial", CreatedAt: now, UpdatedAt: now},
		{ID: "2", Path: "golang2.md", Content: "golang tips", CreatedAt: now, UpdatedAt: now},
		{ID: "3", Path: "python.md", Content: "python guide", CreatedAt: now, UpdatedAt: now},
		{ID: "4", Path: "golang3.md", Content: "golang advanced", CreatedAt: now, UpdatedAt: now},
	}
	for _, note := range notesData {
		err := repo.Save(ctx, note)
		require.NoError(t, err)
	}

	// Buscar por "golang"
	results, err := repo.SearchNotes(ctx, "%golang%", 0, 10)

	require.NoError(t, err)
	require.Len(t, results, 3)
	// Verificar que retorna apenas id, path e updatedAt
	assert.Equal(t, "1", results[0].ID)
	assert.Equal(t, "golang1.md", results[0].Path)
	assert.Equal(t, now.Unix(), results[0].UpdatedAt.Unix())
}

func TestSearchNotes_ComPaginacao_RetornaApenasLimitAndOffset(t *testing.T) {
	db := setupIntegrationTestDB(t)
	defer db.Close()
	repo := storage.NewSqliteNotesRepository(db)
	ctx := context.Background()

	// Inserir 5 notas com "golang"
	now := time.Now()
	for i := 1; i <= 5; i++ {
		note := notes.Note{
			ID:        string(rune('0' + i)),
			Path:      "golang" + string(rune('0'+i)) + ".md",
			Content:   "golang content " + string(rune('0'+i)),
			CreatedAt: now,
			UpdatedAt: now,
		}
		err := repo.Save(ctx, note)
		require.NoError(t, err)
	}

	// Buscar com offset=1 e limit=2
	results, err := repo.SearchNotes(ctx, "%golang%", 1, 2)

	require.NoError(t, err)
	require.Len(t, results, 2)
}

func TestSearchNotes_OffsetMaiorQueTotal_RetornaVazio(t *testing.T) {
	db := setupIntegrationTestDB(t)
	defer db.Close()
	repo := storage.NewSqliteNotesRepository(db)
	ctx := context.Background()

	now := time.Now()
	note := notes.Note{
		ID:        "1",
		Path:      "golang.md",
		Content:   "golang tutorial",
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := repo.Save(ctx, note)
	require.NoError(t, err)

	// Offset > total de resultados
	results, err := repo.SearchNotes(ctx, "%golang%", 100, 10)

	require.NoError(t, err)
	assert.Empty(t, results)
}

func TestSearchNotes_NaoEncontra_RetornaVazio(t *testing.T) {
	db := setupIntegrationTestDB(t)
	defer db.Close()
	repo := storage.NewSqliteNotesRepository(db)
	ctx := context.Background()

	now := time.Now()
	note := notes.Note{
		ID:        "1",
		Path:      "python.md",
		Content:   "python tutorial",
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := repo.Save(ctx, note)
	require.NoError(t, err)

	// Buscar por algo que não existe
	results, err := repo.SearchNotes(ctx, "%golang%", 0, 10)

	require.NoError(t, err)
	assert.Empty(t, results)
}

func TestSearchNotes_LikeEhCaseInsensitive(t *testing.T) {
	db := setupIntegrationTestDB(t)
	defer db.Close()
	repo := storage.NewSqliteNotesRepository(db)
	ctx := context.Background()

	now := time.Now()
	notesData := []notes.Note{
		{ID: "1", Path: "a.md", Content: "Golang Tutorial", CreatedAt: now, UpdatedAt: now},
		{ID: "2", Path: "b.md", Content: "golang tips", CreatedAt: now, UpdatedAt: now},
	}
	for _, note := range notesData {
		err := repo.Save(ctx, note)
		require.NoError(t, err)
	}

	// Buscar por "golang" (minúscula) - LIKE é case-insensitive
	results, err := repo.SearchNotes(ctx, "%golang%", 0, 10)

	require.NoError(t, err)
	require.Len(t, results, 2)
}

func setupIntegrationTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)

	// Criar tabela de testes
	_, err = db.Exec(`
		CREATE TABLE notes (
			id TEXT PRIMARY KEY,
			path TEXT UNIQUE NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		)
	`)
	require.NoError(t, err)

	return db
}
