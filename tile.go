package main

import (
	"fmt"
	"math/rand"
	"sort"
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

func (t *Tiles) Sort() {
	tt := *t
	sort.Slice(tt, func(i, j int) bool {
		if tt[i].Number == tt[j].Number {
			return tt[i].Color < tt[j].Color
		}
		return tt[i].Number < tt[j].Number
	})
	t = &tt
}

func NewTiles() Tiles {
	var tiles Tiles
	for number := 0; number < 10; number++ {
		for _, color := range []Color{Red, Blue} {
			if number == 5 {
				color = Yellow
			}
			tiles = append(tiles, Tile{
				Number: number,
				Color:  color,
			})
		}
	}
	return tiles
}
