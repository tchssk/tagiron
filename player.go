package main

type Player struct {
	Name      string
	ID        string
	Tiles     Tiles
	Questions Questions
}

type Players []Player

func (p *Players) Add(player Player) {
	pp := *p
	pp = append(pp, player)
	n := make(Players, len(pp))
	copy(n, pp)
	*p = n
}

func NewPlayer(name, id string) Player {
	return Player{
		Name:      name,
		ID:        id,
		Tiles:     Tiles{},
		Questions: Questions{},
	}
}
