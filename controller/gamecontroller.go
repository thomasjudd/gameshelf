package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"gameshelf/entity"
)

func GameGet(c *gin.Context) {
	var game entity.Game
	db := c.MustGet("DBClient").(*sql.DB)
	gameId := c.Param("gameid")
	//	query := "SELECT * FROM game where game_id = ?;"
	query := "SELECT game.id, game.name, shelf.name FROM game LEFT JOIN shelf ON game.shelf_id = shelf.id WHERE game.id = ?;"
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

func GameNewGet(c *gin.Context) {
	c.HTML(http.StatusOK, "game_new.tmpl", gin.H{})
}

func GameNewPost(c *gin.Context) {
	db := c.MustGet("DBClient").(*sql.DB)
	name := c.PostForm("name")
	location := c.PostForm("location")

	var shelfId int
	shelfIdQuery := "SELECT id FROM shelf WHERE name = ?;"
	row := db.QueryRow(shelfIdQuery, location)
	err := row.Scan(&shelfId)
	if err != nil {
		panic(err)
	}

	query := `INSERT INTO game (game_id, name, shelf_id) VALUES(NULL, ?, ?);`
	statement, err := db.Prepare(query)
	defer statement.Close()
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(name, shelfId)
	if err != nil {
		panic(err)
	}
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

func GamesGet(c *gin.Context) {
	db := c.MustGet("DBClient").(*sql.DB)
//	query := "SELECT game_id, name, location FROM game"
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

	c.HTML(http.StatusOK, "games.tmpl", gin.H{
		"virtualShelf": virtualShelf,
	})
}
