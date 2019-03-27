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
	questions := stock.Remove(0, 6)
	return Game{
		Tiles:     tiles,
		Stock:     stock,
		Questions: questions,
		Players:   Players{},
	}
}
