package storage

import (
	"database/sql"
	"errors"
)

func OpenDB(path string) (*sql.DB, error) {
	return nil, errors.New("não implementado")
}

func Migrate(db *sql.DB) error {
	return errors.New("não implementado")
}
