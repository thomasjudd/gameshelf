package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"gameshelf/entity"
)

func ShelvesGet(c *gin.Context) {
	db := c.MustGet("DBClient").(*sql.DB)
//	query := "SELECT id, name, location FROM game"
	query := "SELECT game.id, game.name, shelf.name FROM game LEFT JOIN shelf ON game.shelf_id = shelf.id;"
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

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"virtualShelf": virtualShelf,
	})
}

func ShelfGet(c *gin.Context) {
	db := c.MustGet("DBClient").(*sql.DB)
	fmt.Println("db:  ",  db)
	query := "SELECT id, name FROM shelf WHERE id = ?;"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	currShelf := entity.Shelf{}

	for rows.Next() {
		err := rows.Scan(&(currShelf).ShelfId)
		if err != nil {
			panic(err)
		}
		virtualShelf[currGame.Location] = append(virtualShelf[currGame.Location], currGame)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Shelf": virtualShelf,
	})
}
