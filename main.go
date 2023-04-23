package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	_ "github.com/mattn/go-sqlite3"
	"gameshelf/controller"
	//"gameshelf/middleware"
	//"gameshelf/entity"
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
	app.POST("/game/new", controller.GameNewPost)

	app.POST("/game/delete", controller.GameDelete)

	app.Run(":8181")
}
