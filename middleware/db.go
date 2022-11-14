package middleware

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)


func SQLiteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("sqlite3", "./sqlitedb/game_collection.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		c.Set("DBClient",  db)
		c.Next()
	}
}
