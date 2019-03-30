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

func (t *Tiles) Remove(i, j int) Tiles {
	tt := *t
	if len(tt) == 0 {
		return Tiles{}
	}
	if i < 0 {
		i = 0
	}
	if j < 1 {
		j = 1
	}
	if max := len(tt) - 1; i > max {
		i = max
	}
	if max := len(tt); j > max {
		j = max
	}
	if i > j {
		i = j - 1
	}
	tiles := tt[i:j]
	var ttt Tiles
	ttt = append(ttt, tt[:i]...)
	ttt = append(ttt, tt[j:]...)
	new := make(Tiles, len(ttt))
	copy(new, ttt)
	*t = new
	return tiles
}

func (t *Tiles) Add(tiles Tiles) {
	tt := *t
	tt = append(tt, tiles...)
	n := make(Tiles, len(tt))
	copy(n, tt)
	*t = n
}

func (t *Tiles) SumOfAll() int {
	tt := *t
	var sum int
	for _, tile := range tt {
		sum += tile.Number
	}
	return sum
}

func (t *Tiles) SumOfLowerThree() int {
	tt := *t
	tt.Sort()
	max := 3
	if length := len(tt); length < max {
		max = length
	}
	var sum int
	for i := 0; i < max; i++ {
		sum += tt[i].Number
	}
	return sum
}

func (t *Tiles) SumOfUpperThree() int {
	tt := *t
	tt.Sort()
	min := 3
	if length := len(tt); length < min {
		min = length
	}
	var sum int
	for i := 0; i < min; i++ {
		sum += tt[len(tt)-1-i].Number
	}
	return sum
}

func (t *Tiles) SumOfRed() int {
	tt := *t
	var sum int
	for _, tile := range tt {
		if tile.Color == Red {
			sum += tile.Number
		}
	}
	return sum
}

func (t *Tiles) SumOfBlue() int {
	tt := *t
	var sum int
	for _, tile := range tt {
		if tile.Color == Blue {
			sum += tile.Number
		}
	}
	return sum
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
