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

	app.Static("/static/js", "./static/js")
	app.Static("/static/css", "./static/css")

	app.GET("/", controller.ShelvesGet)

	app.GET("/game/:gameid", controller.GameGet)
	app.GET("/game/new", controller.GameNewGet)
	app.POST("/game/new", controller.GameNewPost)

	app.POST("/game/delete", controller.GameDelete)

//	app.GET("/shelf", controller.ShelfGet)

	app.Run(":8181")
}
