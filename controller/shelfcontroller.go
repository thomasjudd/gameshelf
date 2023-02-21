package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"gameshelf/entity"
	"fmt"
)

func ShelvesGet(c *gin.Context) {
	fmt.Println("shelvesget")
	shelves := entity.GetAllShelves()
	fmt.Println("got shelves")

	for i := range shelves {
		shelves[i].Games = entity.GetGamesByShelfId(shelves[i].ShelfId)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Shelves": shelves,
	})
}

func ShelfGet(c *gin.Context) {
	var shelf entity.Shelf
	shelfId  := c.Param("shelfid")
	shelf = entity.GetShelfById(shelfId)

	c.HTML(http.StatusOK, "shelf.tmpl", gin.H{
		"Shelf": shelf.Name,
	})
}
