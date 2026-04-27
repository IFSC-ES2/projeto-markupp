package api_test

import (
	"bytes"
	"context"
	"encoding/json"
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
}

func (f *fakeService) Create(ctx context.Context, path, content string) (notes.Note, error) {
	return f.note, f.err
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
	svc := &fakeService{err: notes.ErrDuplicatePath}
	rec := doPost(t, svc, `{"path":"x.md","content":"y"}`)

	assert.Equal(t, http.StatusConflict, rec.Code)

	var resp map[string]any
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "duplicate_path", resp["error"])
}
