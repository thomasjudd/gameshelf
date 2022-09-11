package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

type Game struct {
	gameId   int    `json:"game_id"`
	name     string `json:"name"`
	location string `json:"location"`
}

var DB *sql.DB

func indexGet(c *gin.Context) {
	//c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func searchGet(c *gin.Context) {
	c.HTML(http.StatusOK, "search.tmpl", gin.H{})
}

func searchPost(c *gin.Context) {
	var game Game
	gameName := c.PostForm("game_name")
	query := "SELECT * FROM game_fts WHERE name MATCH ?;"
	rows, err := DB.Query(query, gameName)
	//	rows, err  := DB.QueryRow(query, gameName)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var name string
		err := rows.Scan(&game.gameId, &game.name, &game.location)
		if err != nil {
			panic(err)
		}
		fmt.Println(name)
	}
	//	if err != nil {
	//		panic(err)
	//  }
	if game.name != "" {
		c.Redirect(301, fmt.Sprintf("/game/%v", game.gameId))
	} else {
		c.JSON(404, gin.H{"msg": "Game Not Found"})
	}
}

func gameGet(c *gin.Context) {
	var game Game
	gameId := c.Param("gameid")
	query := "SELECT * FROM game where game_id = ?;"
	row := DB.QueryRow(query, gameId)
	err := row.Scan(&game.gameId, &game.name, &game.location)

	if err != nil {
		panic(err)
	}
	fmt.Println(game.location)
	c.HTML(http.StatusOK, "game.tmpl", gin.H{
		"game_name":     game.name,
		"game_location": game.location,
	})
}

func gameNewPost(c *gin.Context) {
	name := c.PostForm("name")
	location := c.PostForm("location")
	fmt.Println(location)
	query := `INSERT INTO game (game_id, name, location) VALUES(NULL, ?, ?);`
	statement, err := DB.Prepare(query)
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

func gameNewGet(c *gin.Context) {
	c.HTML(http.StatusOK, "game_new.tmpl", gin.H{})
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./sqlitedb/game_collection.db")
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/", indexGet)

	router.GET("/search", searchGet)
	router.POST("/search", searchPost)

	router.GET("/game/:gameid", gameGet)
	router.GET("/game/new", gameNewGet)
	router.POST("/game/new", gameNewPost)

	router.Run(":80")
}
