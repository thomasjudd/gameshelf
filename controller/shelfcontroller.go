package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"gameshelf/entity"
)

func ShelfGet(c *gin.Context) {
	db := c.MustGet("DBClient").(*sql.DB)
	fmt.Println("db:  ",  db)
	query := "SELECT game_id, name, location FROM game"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

  virtualShelf := make(map[string][]entity.Game)
	currGame := entity.Game{}

	for rows.Next() {
		err := rows.Scan(&(currGame).GameId, &(currGame).Name, &(currGame).Location)
		if err != nil {
			panic(err)
		}
		virtualShelf[currGame.Location] = append(virtualShelf[currGame.Location], currGame)
	}

	c.HTML(http.StatusOK, "shelf.tmpl", gin.H{
		"virtualShelf": virtualShelf,
	})
}
