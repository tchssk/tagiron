package main

import (
	"fmt"
	"math/rand"
)

type Tile struct {
	Color  Color
	Number int
}

type Tiles []Tile

func (t Tile) String() string {
	return fmt.Sprintf("%s%d", t.Color.String(), t.Number)
}

func (t *Tiles) Shuffle() {
	tt := *t
	for i := range tt {
		j := rand.Intn(i + 1)
		tt[i], tt[j] = tt[j], tt[i]
	}
	t = &tt
}
