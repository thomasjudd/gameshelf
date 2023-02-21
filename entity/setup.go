package entity

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)


func ConnectDatabase() *sql.DB {
		db, err := sql.Open("sqlite3", "./sqlitedb/gameshelf.db")
		if err != nil {
			panic(err)
		}
		return db
}


var DB *sql.DB = ConnectDatabase()
