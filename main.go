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

	app.GET("/", controller.IndexGet)

	app.GET("/game/:gameid", controller.GameGet)
	app.POST("/game/new", controller.GameNew)
	app.POST("/game/delete", controller.GameDelete)
	app.GET("/games/manage", controller.GamesManage)

	app.GET("/shelf", controller.ShelfGet)

	app.Run(":8181")
}
