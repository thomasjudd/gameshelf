package entity

type Game struct {
	GameId   int    `json:"game_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
