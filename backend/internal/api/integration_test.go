package api_test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "modernc.org/sqlite"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ifsc-ES2/projeto-markupp/backend/internal/api"
	"github.com/ifsc-ES2/projeto-markupp/backend/internal/notes"
	"github.com/ifsc-ES2/projeto-markupp/backend/internal/storage"
)

const integrationMaxNoteSize = 50 * 1024 * 1024

func setupIntegrationServer(t *testing.T) (*httptest.Server, *sql.DB) {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)
	require.NoError(t, storage.Migrate(db))

	repo := storage.NewSqliteNotesRepository(db)
	svc := notes.NewService(repo, integrationMaxNoteSize)
	router := api.NewRouter(svc)

	server := httptest.NewServer(router)
	t.Cleanup(func() {
		server.Close()
		_ = db.Close()
	})
	return server, db
}

func postNotes(t *testing.T, baseURL, body string) *http.Response {
	t.Helper()
	resp, err := http.Post(baseURL+"/notes", "application/json", strings.NewReader(body))
	require.NoError(t, err)
	return resp
}

func decodeJSON(t *testing.T, body io.Reader) map[string]any {
	t.Helper()
	var result map[string]any
	require.NoError(t, json.NewDecoder(body).Decode(&result))
	return result
}

func TestIntegration_CriarNota_FluxoCompleto(t *testing.T) {
	server, db := setupIntegrationServer(t)

	resp := postNotes(t, server.URL, `{"path":"integracao.md","content":"# teste"}`)
	defer func() { _ = resp.Body.Close() }()

	require.Equal(t, http.StatusCreated, resp.StatusCode)

	body := decodeJSON(t, resp.Body)
	id, _ := body["id"].(string)
	assert.NotEmpty(t, id)
	assert.Equal(t, "integracao.md", body["path"])
	assert.Equal(t, "# teste", body["content"])

	var dbPath, dbContent string
	err := db.QueryRow("SELECT path, content FROM notes WHERE id = ?", id).Scan(&dbPath, &dbContent)
	require.NoError(t, err)
	assert.Equal(t, "integracao.md", dbPath)
	assert.Equal(t, "# teste", dbContent)
}

func TestIntegration_CriarNotaDuplicada_RetornaMensagemLimpa(t *testing.T) {
	server, _ := setupIntegrationServer(t)

	body := `{"path":"dup.md","content":"x"}`
	resp1 := postNotes(t, server.URL, body)
	require.Equal(t, http.StatusCreated, resp1.StatusCode)
	_ = resp1.Body.Close()

	resp2 := postNotes(t, server.URL, body)
	defer func() { _ = resp2.Body.Close() }()
	require.Equal(t, http.StatusConflict, resp2.StatusCode)

	errBody := decodeJSON(t, resp2.Body)
	assert.Equal(t, "duplicate_path", errBody["error"])
	assert.Equal(t, "path já existe", errBody["message"])

	msg, _ := errBody["message"].(string)
	assert.NotContains(t, msg, "salvar nota")
}
