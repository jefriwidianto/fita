package Config

import (
	"database/sql"
)

type ConfigSettingSql struct {}

type Db interface {
	InitDB()
}

type DbSqlConfigName string

const (
	// Database Connection Constant name
	DATABASE_MAIN  DbSqlConfigName = "sqlite3"
)

//maping all connection DB sql
var SqlConnection *sql.DB

func (d DbSqlConfigName) Get() *sql.DB {
	return SqlConnection
}
