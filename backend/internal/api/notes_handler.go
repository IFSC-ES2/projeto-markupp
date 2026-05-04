package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ifsc-ES2/projeto-markupp/backend/internal/notes"
)

type NoteService interface {
	Create(ctx context.Context, path, content string) (notes.Note, error)
	GetNoteById(ctx context.Context, id string) (notes.Note, error)
}

type notesHandler struct {
	svc NoteService
}

type noteResponse struct {
	ID        string    `json:"id"`
	Path      string    `json:"path"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func toResponse(n notes.Note) noteResponse {
	return noteResponse{
		ID:        n.ID,
		Path:      n.Path,
		Content:   n.Content,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}

func (h *notesHandler) create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "invalid_request", "JSON inválido", http.StatusBadRequest)
		return
	}
	note, err := h.svc.Create(r.Context(), req.Path, req.Content)
	if err != nil {
		writeDomainError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, toResponse(note))
}

func (h *notesHandler) get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	note, err := h.svc.GetNoteById(r.Context(), id)
	if err != nil {
		writeDomainError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, toResponse(note))
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, code, message string, status int) {
	writeJSON(w, status, errorResponse{Error: code, Message: message})
}

func writeDomainError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		writeError(w, "not_found", "nota não encontrada", http.StatusNotFound)
	case errors.Is(err, notes.ErrInvalidPath):
		writeError(w, "invalid_path", notes.ErrInvalidPath.Error(), http.StatusBadRequest)
	case errors.Is(err, notes.ErrInvalidContent):
		writeError(w, "invalid_content", notes.ErrInvalidContent.Error(), http.StatusBadRequest)
	case errors.Is(err, notes.ErrDuplicatePath):
		writeError(w, "duplicate_path", notes.ErrDuplicatePath.Error(), http.StatusConflict)
	case errors.Is(err, notes.ErrInvalidId):
		writeError(w, "invalid_id", notes.ErrInvalidId.Error(), http.StatusBadRequest)
	case errors.Is(err, notes.ErrNotFoundId):
		writeError(w, "not_found", "nota não encontrada", http.StatusNotFound)
	default:
		writeError(w, "internal", "erro interno", http.StatusInternalServerError)
	}
}
