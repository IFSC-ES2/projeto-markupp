package api

import (
	"context"
	"net/http"

	"github.com/ifsc-ES2/projeto-markupp/backend/internal/notes"
)

type NoteService interface {
	Create(ctx context.Context, path, content string) (notes.Note, error)
}

type notesHandler struct {
	svc NoteService
}

func (h *notesHandler) create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
