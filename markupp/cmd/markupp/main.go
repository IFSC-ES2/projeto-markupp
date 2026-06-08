package main

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/ifsc-ES2/projeto-markupp/markupp/internal/api"
	"github.com/ifsc-ES2/projeto-markupp/markupp/internal/config"
	"github.com/ifsc-ES2/projeto-markupp/markupp/internal/notes"
	"github.com/ifsc-ES2/projeto-markupp/markupp/internal/storage"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("carregar config", "err", err)
		os.Exit(1)
	}

	db, err := storage.OpenDB(cfg.DBPath)
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
	svc := notes.NewService(repo, cfg.MaxNoteSize)
	router := api.NewRouter(svc)

	addr := ":" + strconv.Itoa(cfg.Port)
	logger.Info("servidor subindo", "addr", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Error("servidor", "err", err)
		os.Exit(1)
	}
}
