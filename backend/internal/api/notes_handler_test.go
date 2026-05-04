package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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
