package api_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ifsc-ES2/projeto-markupp/backend/internal/api"
	"github.com/ifsc-ES2/projeto-markupp/backend/internal/notes"
)

type fakeService struct {
	note notes.Note
	err  error

	updateNote notes.Note
	updateErr  error
	updateArgs struct {
		id      string
		path    string
		content string
		called  bool
	}

	deleteErr    error
	deleteCalled bool
	deletedID    string

	listResult []notes.Note
	listErr    error
	listCalled bool
}

func (f *fakeService) Create(ctx context.Context, path, content string) (notes.Note, error) {
	return f.note, f.err
}

func (f *fakeService) Update(ctx context.Context, id, path, content string) (notes.Note, error) {
	f.updateArgs.called = true
	f.updateArgs.id = id
	f.updateArgs.path = path
	f.updateArgs.content = content
	return f.updateNote, f.updateErr
}

func (f *fakeService) Delete(ctx context.Context, id string) error {
	f.deleteCalled = true
	f.deletedID = id
	return f.deleteErr
}

func (f *fakeService) GetNoteById(ctx context.Context, id string) (notes.Note, error) {
	return f.note, f.err
}

func (f *fakeService) ListNotes(ctx context.Context) ([]notes.Note, error) {
	f.listCalled = true
	return f.listResult, f.listErr
}

func doPost(t *testing.T, svc api.NoteService, body string) *httptest.ResponseRecorder {
	t.Helper()
	router := api.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPost, "/notes", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func doPut(t *testing.T, svc api.NoteService, id, body string) *httptest.ResponseRecorder {
	t.Helper()
	router := api.NewRouter(svc)
	req := httptest.NewRequest(http.MethodPut, "/notes/"+id, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func doDelete(t *testing.T, svc api.NoteService, id string) *httptest.ResponseRecorder {
	t.Helper()
	router := api.NewRouter(svc)
	req := httptest.NewRequest(http.MethodDelete, "/notes/"+id, nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func doGet(t *testing.T, svc api.NoteService, id string) *httptest.ResponseRecorder {
	t.Helper()
	router := api.NewRouter(svc)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/notes/%s", id), nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func doListNotes(t *testing.T, svc api.NoteService) *httptest.ResponseRecorder {
	t.Helper()
	router := api.NewRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/notes", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func TestPostNotes_BodyValido_Retorna201(t *testing.T) {
	svc := &fakeService{
		note: notes.Note{ID: "id-x", Path: "x.md", Content: "y"},
	}
	rec := doPost(t, svc, `{"path":"x.md","content":"y"}`)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "id-x", resp["id"])
	assert.Equal(t, "x.md", resp["path"])
	assert.Equal(t, "y", resp["content"])
}

func TestPostNotes_JSONMalformado_Retorna400InvalidRequest(t *testing.T) {
	svc := &fakeService{}
	rec := doPost(t, svc, `{"path"`)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "invalid_request", resp["error"])
}

func TestPostNotes_ErrInvalidPath_Retorna400(t *testing.T) {
	svc := &fakeService{err: notes.ErrInvalidPath}
	rec := doPost(t, svc, `{"path":"","content":"x"}`)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "invalid_path", resp["error"])
}

func TestPostNotes_ErrInvalidContent_Retorna400(t *testing.T) {
	svc := &fakeService{err: notes.ErrInvalidContent}
	rec := doPost(t, svc, `{"path":"x.md","content":"y"}`)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "invalid_content", resp["error"])
}

func TestPostNotes_ErrDuplicatePath_Retorna409(t *testing.T) {
	wrapped := fmt.Errorf("salvar nota %q: %w", "id-interno-secreto", notes.ErrDuplicatePath)
	svc := &fakeService{err: wrapped}
	rec := doPost(t, svc, `{"path":"x.md","content":"y"}`)

	assert.Equal(t, http.StatusConflict, rec.Code)

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "duplicate_path", resp["error"])
	assert.Equal(t, notes.ErrDuplicatePath.Error(), resp["message"])
	assert.NotContains(t, resp["message"], "id-interno-secreto")
}

func TestPutNotes_BodyValido_Retorna200(t *testing.T) {
	svc := &fakeService{
		updateNote: notes.Note{ID: "id-1", Path: "novo.md", Content: "novo"},
	}
	rec := doPut(t, svc, "id-1", `{"path":"novo.md","content":"novo"}`)

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "id-1", resp["id"])
	assert.Equal(t, "novo.md", resp["path"])
	assert.Equal(t, "novo", resp["content"])

	assert.Equal(t, "id-1", svc.updateArgs.id)
	assert.Equal(t, "novo.md", svc.updateArgs.path)
	assert.Equal(t, "novo", svc.updateArgs.content)
}

func TestPutNotes_JSONMalformado_Retorna400(t *testing.T) {
	svc := &fakeService{}
	rec := doPut(t, svc, "id-1", `{"path"`)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "invalid_request", resp["error"])
}

func TestPutNotes_ErrNotFound_Retorna404(t *testing.T) {
	svc := &fakeService{updateErr: notes.ErrNotFound}
	rec := doPut(t, svc, "fantasma", `{"path":"x.md","content":"y"}`)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "not_found", resp["error"])
}

func TestPutNotes_ErrDuplicatePath_Retorna409(t *testing.T) {
	svc := &fakeService{updateErr: notes.ErrDuplicatePath}
	rec := doPut(t, svc, "id-1", `{"path":"existente.md","content":"y"}`)

	assert.Equal(t, http.StatusConflict, rec.Code)
	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "duplicate_path", resp["error"])
}

func TestPutNotes_ErrInvalidPath_Retorna400(t *testing.T) {
	svc := &fakeService{updateErr: notes.ErrInvalidPath}
	rec := doPut(t, svc, "id-1", `{"path":"","content":"y"}`)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "invalid_path", resp["error"])
}

func TestDeleteNotes_Sucesso_Retorna204(t *testing.T) {
	svc := &fakeService{}
	rec := doDelete(t, svc, "id-1")

	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Empty(t, rec.Body.Bytes())
	assert.True(t, svc.deleteCalled)
	assert.Equal(t, "id-1", svc.deletedID)
}

func TestDeleteNotes_ErrNotFound_Retorna404(t *testing.T) {
	svc := &fakeService{deleteErr: notes.ErrNotFound}
	rec := doDelete(t, svc, "fantasma")

	assert.Equal(t, http.StatusNotFound, rec.Code)
	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "not_found", resp["error"])
}

func TestGetNotes_IDValido_Retorna200(t *testing.T) {
	svc := &fakeService{
		note: notes.Note{ID: "id-x", Path: "x.md", Content: "y"},
	}
	rec := doGet(t, svc, "id-x")

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "id-x", resp["id"])
	assert.Equal(t, "x.md", resp["path"])
	assert.Equal(t, "y", resp["content"])
}

func TestGetNotes_IDVazio_Retorna404(t *testing.T) {
	svc := &fakeService{}
	rec := doGet(t, svc, "")

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetNotes_IDNaoEncontrado_Retorna404(t *testing.T) {
	svc := &fakeService{err: sql.ErrNoRows}
	rec := doGet(t, svc, "id-inexistente")

	assert.Equal(t, http.StatusNotFound, rec.Code)

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "not_found", resp["error"])
}

type stubRepository struct {
	noteToReturn notes.Note
	getError     error

	listResult []notes.Note
	listErr    error
}

func (s *stubRepository) Save(ctx context.Context, n notes.Note) error {
	return nil
}

func (s *stubRepository) Update(ctx context.Context, id, path, content string, updatedAt time.Time) (notes.Note, error) {
	return notes.Note{}, nil
}

func (s *stubRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (s *stubRepository) GetNoteByID(ctx context.Context, id string) (notes.Note, error) {
	return s.noteToReturn, s.getError
}

func (s *stubRepository) ListNotes(ctx context.Context) ([]notes.Note, error) {
	return s.listResult, s.listErr
}

func TestGetNotes_IDNaoEncontrado_ComServiceReal_Retorna404(t *testing.T) {
	repo := &stubRepository{getError: sql.ErrNoRows}
	svc := notes.NewService(repo, 1000)

	router := api.NewRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/notes/id-inexistente", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "not_found", resp["error"])
}

func TestGetNotes_IDVazio_ComServiceReal_Retorna400(t *testing.T) {
	repo := &stubRepository{}
	svc := notes.NewService(repo, 1000)

	router := api.NewRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/notes/", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetNotes_IDValido_ComServiceReal_Retorna200(t *testing.T) {
	notaEsperada := notes.Note{
		ID:      "id-123",
		Path:    "file.md",
		Content: "conteudo",
	}
	repo := &stubRepository{noteToReturn: notaEsperada}
	svc := notes.NewService(repo, 1000)

	router := api.NewRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/notes/id-123", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "id-123", resp["id"])
	assert.Equal(t, "file.md", resp["path"])
	assert.Equal(t, "conteudo", resp["content"])
}

func TestListNotes_DBVazio_Retorna200ComArrayVazio(t *testing.T) {
	svc := &fakeService{listResult: []notes.Note{}}
	rec := doListNotes(t, svc)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var resp []map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Empty(t, resp)
	assert.True(t, svc.listCalled)
}

func TestListNotes_ComNotas_RetornaArrayCompleto(t *testing.T) {
	svc := &fakeService{listResult: []notes.Note{
		{ID: "id-1", Path: "a.md", Content: "aaa"},
		{ID: "id-2", Path: "b.md", Content: "bbb"},
	}}
	rec := doListNotes(t, svc)

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp []map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	require.Len(t, resp, 2)
	assert.Equal(t, "id-1", resp[0]["id"])
	assert.Equal(t, "a.md", resp[0]["path"])
	assert.Equal(t, "aaa", resp[0]["content"])
	assert.Equal(t, "id-2", resp[1]["id"])
}

func TestListNotes_ServiceRetornaErro_Retorna500(t *testing.T) {
	svc := &fakeService{listErr: errors.New("erro inesperado")}
	rec := doListNotes(t, svc)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "internal", resp["error"])
}
