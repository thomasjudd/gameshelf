package middleware

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)


func SQLiteMiddleware() gin.HandlerFunc {
	db, err := sql.Open("sqlite3", "./sqlitedb/game_collection.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return func(c *gin.Context) {
		c.Set("DBClient",  db)
		c.Next()
	}
}
