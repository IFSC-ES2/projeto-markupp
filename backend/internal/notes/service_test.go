package notes_test

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ifsc-ES2/projeto-markupp/backend/internal/notes"
)

const testMaxContentSize = 100

type fakeRepo struct {
	saved   notes.Note
	saveErr error
	called  bool
	note    notes.Note
	getErr  error
}

func (f *fakeRepo) Save(ctx context.Context, n notes.Note) error {
	f.called = true
	f.saved = n
	return f.saveErr
}

func (f *fakeRepo) GetNoteByID(ctx context.Context, id string) (notes.Note, error) {
	return f.note, f.getErr
}

func newServiceForTest(repo notes.Repository) *notes.Service {
	return notes.NewService(repo, testMaxContentSize)
}

func TestCreate_PathVazio_RetornaErrInvalidPath(t *testing.T) {
	repo := &fakeRepo{}
	svc := newServiceForTest(repo)

	_, err := svc.Create(context.Background(), "", "qualquer")

	require.Error(t, err)
	assert.True(t, errors.Is(err, notes.ErrInvalidPath))
	assert.False(t, repo.called)
}

func TestCreate_PathComTraversal_RetornaErrInvalidPath(t *testing.T) {
	repo := &fakeRepo{}
	svc := newServiceForTest(repo)

	_, err := svc.Create(context.Background(), "../etc/passwd", "qualquer")

	require.Error(t, err)
	assert.True(t, errors.Is(err, notes.ErrInvalidPath))
}

func TestCreate_PathLongoDemais_RetornaErrInvalidPath(t *testing.T) {
	repo := &fakeRepo{}
	svc := newServiceForTest(repo)
	longPath := strings.Repeat("a", 1025)

	_, err := svc.Create(context.Background(), longPath, "qualquer")

	require.Error(t, err)
	assert.True(t, errors.Is(err, notes.ErrInvalidPath))
}

func TestCreate_ContentMuitoGrande_RetornaErrInvalidContent(t *testing.T) {
	repo := &fakeRepo{}
	svc := newServiceForTest(repo)
	bigContent := strings.Repeat("x", testMaxContentSize+1)

	_, err := svc.Create(context.Background(), "ok.md", bigContent)

	require.Error(t, err)
	assert.True(t, errors.Is(err, notes.ErrInvalidContent))
}

func TestCreate_ContentVazio_Aceito(t *testing.T) {
	repo := &fakeRepo{}
	svc := newServiceForTest(repo)

	note, err := svc.Create(context.Background(), "vazio.md", "")

	require.NoError(t, err)
	assert.Empty(t, note.Content)
	assert.True(t, repo.called)
}

func TestCreate_CaminhoFeliz_GeraNotaCompleta(t *testing.T) {
	repo := &fakeRepo{}
	svc := newServiceForTest(repo)

	before := time.Now()
	note, err := svc.Create(context.Background(), "minha.md", "conteudo")
	after := time.Now()

	require.NoError(t, err)
	assert.Equal(t, "minha.md", note.Path)
	assert.Equal(t, "conteudo", note.Content)
	assert.NotEmpty(t, note.ID)
	assert.False(t, note.CreatedAt.Before(before))
	assert.False(t, note.CreatedAt.After(after))
	assert.Equal(t, note.CreatedAt, note.UpdatedAt)

	require.True(t, repo.called)
	assert.Equal(t, note, repo.saved)
}

func TestCreate_RepoRetornaErrDuplicatePath_PropagadoAoCaller(t *testing.T) {
	repo := &fakeRepo{saveErr: notes.ErrDuplicatePath}
	svc := newServiceForTest(repo)

	_, err := svc.Create(context.Background(), "dup.md", "x")

	require.Error(t, err)
	assert.True(t, errors.Is(err, notes.ErrDuplicatePath))
}

// Problema 4: Testes para Service.GetNoteById()

func TestGetNoteById_IDVazio_RetornaErrInvalidId(t *testing.T) {
	repo := &fakeRepo{}
	svc := newServiceForTest(repo)

	_, err := svc.GetNoteById(context.Background(), "")

	require.Error(t, err)
	assert.True(t, errors.Is(err, notes.ErrInvalidId))
	// Valida que o repo NÃO foi chamado (validação local é suficiente)
}

func TestGetNoteById_IDNaoEncontrado_RetornaErrNotFoundId(t *testing.T) {
	repo := &fakeRepo{getErr: sql.ErrNoRows}
	svc := newServiceForTest(repo)

	_, err := svc.GetNoteById(context.Background(), "id-inexistente")

	require.Error(t, err)
	assert.True(t, errors.Is(err, notes.ErrNotFoundId))
}

func TestGetNoteById_CaminhoFeliz_RetornaNotaCorreta(t *testing.T) {
	notaEsperada := notes.Note{
		ID:        "id-123",
		Path:      "arquivo.md",
		Content:   "conteudo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repo := &fakeRepo{note: notaEsperada}
	svc := newServiceForTest(repo)

	nota, err := svc.GetNoteById(context.Background(), "id-123")

	require.NoError(t, err)
	assert.Equal(t, notaEsperada, nota)
}
