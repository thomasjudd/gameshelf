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
    fmt.Println("game name: ", game.Name)
    fmt.Println("shelf id: ", game.ShelfId)
  if err != nil {
    	panic(err)
  }

    fmt.Println("game name: ", game.Name)
    fmt.Println("game shelf id: ", game.ShelfId)

    entity.CreateGame(game)
    c.JSON(http.StatusCreated, game)
}

func GameDelete(c *gin.Context) {
    var game entity.Game

    err := c.BindJSON(&game)
    if err != nil {
     	panic(err)
    }

    fmt.Println("game name: ", game.Name)

  entity.DeleteGame(game)
}
