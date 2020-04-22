package sql

import (
	"database/sql"
	"fmt"
	"log"
)

func MustPrepare(db *sql.DB, format string, args ...interface{}) *sql.Stmt {
	fmt.Println(fmt.Sprintf(format, args...))
	stmt, err := db.Prepare(fmt.Sprintf(format, args...))
	if err != nil {
		log.Panicf("Prepare: %+v", err)
	}
	return stmt
}
