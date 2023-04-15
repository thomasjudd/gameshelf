package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"gameshelf/entity"
	"strconv"
	"fmt"
)

func GameGet(c *gin.Context) {
	gameId := c.Param("gameid")
	gameIdInt, err :=  strconv.Atoi(gameId)
	if err != nil {
		panic(err)
	}

	game := entity.GetGame(gameIdInt)
	c.HTML(http.StatusOK, "game.tmpl", gin.H{
		"game": game,
	})
}

func GameNewGet(c *gin.Context) {
	c.HTML(http.StatusOK, "game_new.tmpl", gin.H{})
}

func GameNewPost(c *gin.Context) {
  var game entity.Game
	err := c.BindJSON(&game)
	fmt.Println("game name: ", game.Name)
	fmt.Println("shelf id: ", game.ShelfId)
  if err != nil {
		panic(err)
  }

	fmt.Println("game name: ", game.Name)
	fmt.Println("game shelf id: ", game.ShelfId)

	entity.CreateGame(game)
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
