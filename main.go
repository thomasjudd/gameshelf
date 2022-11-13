package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	_ "github.com/mattn/go-sqlite3"
	"gameshelf/controller"
	"gameshelf/middleware"
)

func main() {

	html, err := template.ParseGlob("templates/*/*.tmpl")
	if err != nil {
		panic(err)
	}
	app := gin.Default()
	app.Use(middleware.SQLiteMiddleware())
  app.SetHTMLTemplate(html)

	app.GET("/", controller.SearchGet)
	app.POST("/", controller.SearchPost)

	app.GET("/game/:gameid", controller.GameGet)
	app.GET("/game/new", controller.GameNewGet)
	app.POST("/game/new", controller.GameNewPost)

	app.GET("/games", controller.GamesGet)

	app.Run(":8181")
}
