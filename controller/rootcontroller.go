package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

func IndexGet (c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
