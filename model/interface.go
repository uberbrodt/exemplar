package model

import "github.com/jmoiron/sqlx"

type ExemplarSqlx interface {
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
}
