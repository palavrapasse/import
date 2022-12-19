package internal

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseContext struct {
	DB       *sql.DB
	FilePath string
}

const (
	_sqliteDriverName = "sqlite3"
)

func NewDatabaseContext(fp string) (DatabaseContext, error) {
	db, err := sql.Open(_sqliteDriverName, fp)

	return DatabaseContext{
		DB:       db,
		FilePath: fp,
	}, err
}
