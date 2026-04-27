package api

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter(svc NoteService) chi.Router {
	r := chi.NewRouter()
	h := &notesHandler{svc: svc}
	r.Post("/notes", h.create)
	return r
}
