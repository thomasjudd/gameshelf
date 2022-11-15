package model

import (
	"gameshelf/entity"
	"database/sql"
)

func FindGameByName(dbClient *sql.DB, gameName string) entity.Game {
	var game entity.Game
	query := "SELECT * FROM game_fts WHERE name MATCH ?;"
	rows, err := dbClient.Query(query, gameName)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&game.GameId, &game.Name, &game.Location)
		if err != nil {
			panic(err)
		}
	}

	return game
}
