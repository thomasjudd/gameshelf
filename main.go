package main

import (
	"gameshelf/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
)

func main() {
	html, err := template.ParseGlob("templates/*/*.tmpl")
	if err != nil {
		panic(err)
	}

	app := gin.Default()
	app.SetHTMLTemplate(html)

	app.Static("/static/js", "./static/js")
	app.Static("/static/css", "./static/css")

	app.GET("/", controller.ShelvesGet)
	app.GET("/shelf/:shelfid", controller.ShelfGet)

	app.GET("/game/:gameid", controller.GameGet)
	app.DELETE("/game/:gameid", controller.GameDelete)
	app.POST("/game/new", controller.GameNewPost)

	app.Run(":8181")
}
