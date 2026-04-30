package storage

import (
	"database/sql"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"

	"github.com/ifsc-ES2/projeto-markupp/backend/migrations"
)

func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *sql.DB) error {
	goose.SetBaseFS(migrations.FS)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}
	return goose.Up(db, ".")
}
