package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/palavrapasse/import/internal/entity"
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

func (ctx DatabaseContext) Insert(t DatabaseTable) (DatabaseTable, error) {
	var tx *sql.Tx
	var stmt *sql.Stmt

	var updatedRecords Records

	tx, err := ctx.DB.Begin()

	defer func() {
		if err != nil {
			if stmt != nil {
				stmt.Close()
			}

			tx.Rollback()
		}
	}()

	stmt, err = t.PrepareInsertStatement(tx)

	if err == nil {
		records := t

		for _, r := range records {
			var res sql.Result
			var lid int64

			res, err = stmt.Exec(Values(r))
			lid, err = res.LastInsertId()

			if err != nil {
				updatedRecords = append(updatedRecords, CopyWithNewKey(r, entity.AutoGenKey(lid)))
			}
		}
	}

	return t.Copy(updatedRecords), err
}
