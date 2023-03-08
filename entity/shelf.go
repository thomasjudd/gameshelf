package entity

import (
	"fmt"
)

type Shelf struct {
	Games []Game`json:"games"`
	ShelfId int `json:"id"`
	Name string `json:"name"`
	LEDOn bool `json:"led_on"`
}

func GetAllShelves() []Shelf {
	query := "SELECT * FROM shelf"
	rows, err := DB.Query(query)
	if err != nil {
		panic(err)
	}

	shelves := []Shelf{}
	var currShelf Shelf
	for rows.Next() {
		err := rows.Scan(&(currShelf).ShelfId, &(currShelf).Name)
		if err != nil {
			panic(err)
		}
		shelves = append(shelves, currShelf)
	}
	return shelves
}

func GetShelfById(shelfId int) Shelf {
	var shelf Shelf
	query := "SELECT * FROM shelf WHERE id = ?;"
	row := DB.QueryRow(query, fmt.Sprintf("%v", shelfId))
	err := row.Scan(&shelf.ShelfId, &shelf.Name) 
	if err != nil {
		panic(err)
	}

	games := GetGamesByShelfId(shelfId)
	shelf.Games = games

	return shelf
}

func GetShelfByName(shelfName string) Shelf {
	var shelf Shelf
	query :=  "SELECT * FROM shelf WHERE name = ?;"
	row  := DB.QueryRow(query)
	err := row.Scan(&shelf.ShelfId, &shelf.Name)
	if err != nil {
		panic(err)
	}

	return shelf
}
