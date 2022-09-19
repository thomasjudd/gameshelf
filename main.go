package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

type Game struct {
	GameId   int    `json:"game_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

var DB *sql.DB

func indexGet(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"pages": []string{"search","games"},
	})
}

func searchGet(c *gin.Context) {
	c.HTML(http.StatusOK, "search.tmpl", gin.H{})
}

func searchPost(c *gin.Context) {
	var game Game
	gameName := c.PostForm("game_name")
	query := "SELECT * FROM game_fts WHERE name MATCH ?;"
	rows, err := DB.Query(query, gameName)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var name string
		err := rows.Scan(&game.GameId, &game.Name, &game.Location)
		if err != nil {
			panic(err)
		}
		fmt.Println(name)
	}
	if game.Name != "" {
		c.Redirect(301, fmt.Sprintf("/game/%v", game.GameId))
	} else {
		c.JSON(404, gin.H{"msg": "Game Not Found"})
	}
}

func gameGet(c *gin.Context) {
	var game Game
	gameId := c.Param("gameid")
	query := "SELECT * FROM game where game_id = ?;"
	row := DB.QueryRow(query, gameId)
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

func gamesGet(c *gin.Context) {
	query := "SELECT game_id, name FROM game"
	rows, err := DB.Query(query)
	if err != nil {
		panic(err)
	}

	var games []Game
	currGame := Game{}

	for rows.Next() {
		err := rows.Scan(&(currGame).GameId, &(currGame).Name)
		if err != nil {
			panic(err)
		}

		games = append(games, currGame)
	}

	c.HTML(http.StatusOK, "games.tmpl", gin.H{
		"games": games,
	})
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

	router.GET("/games", gamesGet)

	router.Run(":8181")
}
