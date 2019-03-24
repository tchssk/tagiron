package main

type Player struct {
	Name      string
	ID        string
	Tiles     Tiles
	Questions Questions
}

type Players []Player

func NewPlayer(name, id string) Player {
	return Player{
		Name:      name,
		ID:        id,
		Tiles:     Tiles{},
		Questions: Questions{},
	}
}
