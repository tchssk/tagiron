package main

type Game struct {
	Tiles     Tiles
	Stock     Questions
	Questions Questions
	Players   Players
}

func NewGame() Game {
	var (
		tiles = NewTiles()
		stock = NewQuestions()
	)
	tiles.Shuffle()
	stock.Shuffle()
	return Game{
		Tiles:     tiles,
		Stock:     stock,
		Questions: Questions{},
		Players:   Players{},
	}
}
