package sql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gopub/log"
)

type (
	DB          = sql.DB
	Tx          = sql.Tx
	TxOptions   = sql.TxOptions
	Stmt        = sql.Stmt
	Row         = sql.Row
	Rows        = sql.Rows
	Conn        = sql.Conn
	Result      = sql.Result
	NullInt64   = sql.NullInt64
	NullTime    = sql.NullTime
	NullBool    = sql.NullBool
	NullFloat64 = sql.NullFloat64
	NullInt32   = sql.NullInt32
	NullString  = sql.NullString
	Scanner     = sql.Scanner
)

var (
	ErrNoRows   = sql.ErrNoRows
	ErrTxDone   = sql.ErrTxDone
	ErrConnDone = sql.ErrConnDone
)

type ColumnScanner interface {
	Scan(dest ...interface{}) error
}

type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func OpenPostgres(dbURL string) *sql.DB {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Panicf("Open %s: %+v", dbURL, err)
	}

	err = db.Ping()
	if err != nil {
		log.Panicf("Ping %s: %+v", dbURL, err)
	}
	return db
}

func BuildPostgresURL(name, host string, port int, user, password string, sslEnabled bool) string {
	if host == "" {
		host = "localhost"
	}
	if port == 0 {
		port = 5432
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, name)
	if !sslEnabled {
		url = url + "?sslmode=disable"
	}
	return url
}

func Escape(s string) string {
	s = strings.Replace(s, ",", "\\,", -1)
	s = strings.Replace(s, "(", "\\(", -1)
	s = strings.Replace(s, ")", "\\)", -1)
	s = strings.Replace(s, "\"", "\\\"", -1)
	return s
}
