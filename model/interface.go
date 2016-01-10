package model

import "github.com/jmoiron/sqlx"
import "database/sql"

type ExemplarSqlx interface {
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	NamedExec(query string, arg interface{}) (sql.Result, error)
}
