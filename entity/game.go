package entity

import "fmt"

type Game struct {
	GameId   int    `json:"game_id"`
	Name     string `json:"name"`
	ShelfId  int `json:"shelf_id"`
}

func GetGame(gameId int) Game {
	var game Game
	query := "SELECT * FROM game WHERE id = ?;"
	row := DB.QueryRow(query, gameId)
	err := row.Scan(&(game).GameId, &(game).Name, &(game).ShelfId)
	if err != nil {
		panic(err)
	}
	return game
}

func GetGamesByShelfId(shelfId int) []Game {
	query := "SELECT id, name FROM game WHERE shelf_id = ?;"
	rows, err := DB.Query(query, fmt.Sprintf("%v",shelfId))
	if err != nil {
		panic(err)
	}

	games := []Game{}
	var currGame Game

	for rows.Next() {
		err := rows.Scan(&(currGame).GameId, &(currGame).Name)
		if err != nil {
			panic(err)
		}
		currGame.ShelfId = shelfId
		games = append(games, currGame)
	}

	return games
}


func CreateGame(game Game) {
	query := "INSERT into  game (id, name, shelf_id)  VALUES(NULL,  ?, ?)"
	statement, err  := DB.Prepare(query)
	defer statement.Close()
	if  err != nil {
		panic(err)
	}

	_, err = statement.Exec(game.Name, game.ShelfId)
	if err != nil {
		panic(err)
	}
}
