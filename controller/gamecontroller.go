package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"gameshelf/entity"
)

func GameGet(c *gin.Context) {
	gameId := c.Param("gameid")
	//	query := "SELECT * FROM game where game_id = ?;"
	game := entity.GetGame(gameId)
	c.HTML(http.StatusOK, "game.tmpl", gin.H{
		"game_name":     game.Name,
		"game_location": game.ShelfId,
	})
}

func GameNewGet(c *gin.Context) {
	c.HTML(http.StatusOK, "game_new.tmpl", gin.H{})
}

func GameNewPost(c *gin.Context) {
	name := c.PostForm("name")
	shelfName := c.PostForm("location")

	shelf := entity.GetShelfByName(shelfName)

	newGame := entity.Game{
		Name: name,
		ShelfId: shelf.ShelfId,
	}

	entity.CreateGame(newGame)
}

func GameDelete(c *gin.Context) {
	db := c.MustGet("DBClient").(*sql.DB)
	name := c.PostForm("name")
	query := `DELETE FROM game WHERE name = ?;`
	statement, err := db.Prepare(query)
	defer statement.Close()
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(name)
	if err != nil {
		panic(err)
	}
}
