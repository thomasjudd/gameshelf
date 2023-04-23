package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
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

func GameNewPost(c *gin.Context) {
  var game entity.Game
  err := c.BindJSON(&game)
  if err != nil {
    	panic(err)
  }

  entity.CreateGame(game)
	game = entity.GetGameByName(game.Name)
  c.JSON(http.StatusCreated, game)
}

func GameDelete(c *gin.Context) {
	gameId := c.Param("gameid")
	gameIdInt, err := strconv.Atoi(gameId)
	fmt.Println("target to delete ", gameIdInt)
	if err != nil {
		panic(err)
	}
  entity.DeleteGameById(gameIdInt)
}
