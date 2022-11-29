package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"gameshelf/entity"
	"gameshelf/model"
)

func GameGet(c *gin.Context) {
	var game entity.Game
	db := c.MustGet("DBClient").(*sql.DB)
	gameId := c.Param("gameid")
	query := "SELECT * FROM game where game_id = ?;"
	row := db.QueryRow(query, gameId)
	err := row.Scan(&game.GameId, &game.Name, &game.Location)
	if err != nil {
		panic(err)
	}
	fmt.Println(game.Location)
	c.HTML(http.StatusOK, "game.tmpl", gin.H{
		"game_name":     game.Name,
		"game_location": game.Location,
	})
}

func GamesManage(c *gin.Context) {
	c.HTML(http.StatusOK, "games_manage.tmpl", gin.H{})
}

func GameNewPost(c *gin.Context) {
	db := c.MustGet("DBClient").(*sql.DB)
	name := c.PostForm("name")
	location := c.PostForm("location")
	fmt.Println(location)
	query := `INSERT INTO game (game_id, name, location) VALUES(NULL, ?, ?);`
	statement, err := db.Prepare(query)
	fmt.Println(statement)
	defer statement.Close()
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(name, location)
	if err != nil {
		panic(err)
	}
}

func GamesGet(c *gin.Context) {
	db := c.MustGet("DBClient").(*sql.DB)
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

	c.HTML(http.StatusOK, "games.tmpl", gin.H{
		"virtualShelf": virtualShelf,
	})
}
