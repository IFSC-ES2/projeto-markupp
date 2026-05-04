package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/ifsc-ES2/projeto-markupp/backend/internal/notes"
)

type NoteService interface {
	Create(ctx context.Context, path, content string) (notes.Note, error)
	Update(ctx context.Context, id, path, content string) (notes.Note, error)
	Delete(ctx context.Context, id string) error
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

func (h *notesHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "invalid_request", "JSON inválido", http.StatusBadRequest)
		return
	}
	note, err := h.svc.Update(r.Context(), id, req.Path, req.Content)
	if err != nil {
		writeDomainError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, toResponse(note))
}

func (h *notesHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.svc.Delete(r.Context(), id); err != nil {
		writeDomainError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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
	case errors.Is(err, notes.ErrInvalidPath):
		writeError(w, "invalid_path", notes.ErrInvalidPath.Error(), http.StatusBadRequest)
	case errors.Is(err, notes.ErrInvalidContent):
		writeError(w, "invalid_content", notes.ErrInvalidContent.Error(), http.StatusBadRequest)
	case errors.Is(err, notes.ErrDuplicatePath):
		writeError(w, "duplicate_path", notes.ErrDuplicatePath.Error(), http.StatusConflict)
	case errors.Is(err, notes.ErrNotFound):
		writeError(w, "not_found", notes.ErrNotFound.Error(), http.StatusNotFound)
	default:
		writeError(w, "internal", "erro interno", http.StatusInternalServerError)
	}
}
