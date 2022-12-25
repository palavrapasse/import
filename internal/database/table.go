package database

import (
	"database/sql"
	"fmt"
	"strings"
)

const (
	unknownTableName = "void"
)

var unknownTableFields = []Field{}
var unknownRecordValues = []any{}

type Table interface {
	Name() string
	Records() []Record
	Fields() []Field
	Copy(Records) Table
	PrepareInsertStatement(tx *sql.Tx) (*sql.Stmt, error)
}

type DatabaseTable struct {
	Table
	Records
}

type PrimaryTable DatabaseTable
type ForeignTable DatabaseTable

type BadActorTable = PrimaryTable
type CredentialsTable = PrimaryTable
type HashCredentialsTable = ForeignTable
type HashUserTable = ForeignTable
type LeakBadActorTable = ForeignTable
type LeakCredentialsTable = ForeignTable
type LeakPlatformTable = ForeignTable
type LeakTable = PrimaryTable
type LeakUserTable = ForeignTable
type PlatformTable = PrimaryTable
type UserCredentialsTable = ForeignTable
type UserTable = PrimaryTable

func (t DatabaseTable) Name() string {
	rs := t.Records

	if len(rs) > 0 {
		return strings.Split(fmt.Sprintf("%T", rs[0]), ".")[1]
	} else {
		return unknownTableName
	}
}

func (t DatabaseTable) Fields() []Field {
	rs := t.Records

	if len(rs) > 0 {
		return Fields(rs[0])
	} else {
		return unknownTableFields
	}
}

func (pt PrimaryTable) PrepareInsertStatement(tx *sql.Tx) (*sql.Stmt, error) {
	return tx.Prepare(pt.prepareInsertStatementString())

}

func (ft ForeignTable) PrepareInsertStatement(tx *sql.Tx) (*sql.Stmt, error) {
	return tx.Prepare(ft.prepareInsertStatementString())
}

func (pt PrimaryTable) prepareInsertStatementString() string {
	tableName := pt.Name()
	tableFields := pt.Fields()[1:]

	return prepareInsertStatementString(tableName, tableFields)
}

func (ft ForeignTable) prepareInsertStatementString() string {
	tableName := ft.Name()
	tableFields := ft.Fields()

	return prepareInsertStatementString(tableName, tableFields)
}

func prepareInsertStatementString(tableName string, tableFields []Field) string {
	tablePlaceholders := stringSliceMap(func(v any) string { return "?" }, tableFields)

	tableFieldsJoin := strings.Join(toStringSlice(tableFields), ", ")
	tablePlaceholdersJoin := strings.Join(toStringSlice(tablePlaceholders), ", ")

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, tableFieldsJoin, tablePlaceholdersJoin)
}

func toStringSlice[T any](s []T) []string {
	return stringSliceMap(
		func(v any) string {
			return fmt.Sprintf("%v", v)
		}, s,
	)
}

func stringSliceMap[T any](m func(v any) string, s []T) []string {
	ss := make([]string, len(s))

	for i, v := range s {
		ss[i] = m(v)
	}

	return ss
}
