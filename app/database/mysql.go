package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbms         = os.Getenv("DBMS")
	host         = os.Getenv("DB_HOST")
	port         = os.Getenv("DB_PORT")
	user         = os.Getenv("DB_USER")
	password     = os.Getenv("DB_PASSWORD")
	databaseName = os.Getenv("DB_NAME")
)

func Connect() (*sql.DB, error) {
	var dbSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, databaseName)

	db, err := sql.Open(dbms, dbSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
