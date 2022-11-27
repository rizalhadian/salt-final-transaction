package pkg_database_mysql

import (
	"database/sql"
	"fmt"
	"time"
)

func InitDBMysql() *sql.DB {
	dbConn := "mysql"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "2022_salt_final"

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open(dbConn, connection)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(60 * time.Second)

	return db
}
