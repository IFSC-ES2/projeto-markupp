package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ifsc-ES2/projeto-markupp/backend/internal/api"
	"github.com/ifsc-ES2/projeto-markupp/backend/internal/notes"
	"github.com/ifsc-ES2/projeto-markupp/backend/internal/storage"
)

const (
	serverPort  = "8080"
	dbPath      = "./markupp.db"
	maxNoteSize = 50 * 1024 * 1024
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := storage.OpenDB(dbPath)
	if err != nil {
		logger.Error("abrir db", "err", err)
		os.Exit(1)
	}
	defer func() { _ = db.Close() }()

	if err := storage.Migrate(db); err != nil {
		logger.Error("aplicar migrations", "err", err)
		os.Exit(1)
	}

	repo := storage.NewSqliteNotesRepository(db)
	svc := notes.NewService(repo, maxNoteSize)
	router := api.NewRouter(svc)

	addr := ":" + serverPort
	logger.Info("servidor subindo", "addr", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Error("servidor", "err", err)
		os.Exit(1)
	}
}
