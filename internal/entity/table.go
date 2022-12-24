package entity

import "database/sql"

type Record interface {
	PrepareInsertStatement(db *sql.DB) *sql.Stmt
}

type Table []Record
