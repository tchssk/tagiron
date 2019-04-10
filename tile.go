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

func (t *Tiles) SumOfCenterThree() int {
	tt := *t
	tt.Sort()
	min := len(tt)/2 - 1
	if min < 0 {
		min = 0
	}
	max := min + 3
	if length := len(tt); length < max {
		max = length
	}
	var sum int
	for i := min; i < max; i++ {
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

func (t *Tiles) Difference() int {
	tt := *t
	tt.Sort()
	length := len(tt)
	if length < 1 {
		return 0
	}
	return tt[length-1].Number - tt[0].Number
}

func (t *Tiles) Odd() int {
	tt := *t
	var n int
	for _, tile := range tt {
		if tile.Number%2 != 0 {
			n++
		}
	}
	return n
}

func (t *Tiles) Even() int {
	tt := *t
	var n int
	for _, tile := range tt {
		if tile.Number%2 == 0 {
			n++
		}
	}
	return n
}

func (t *Tiles) NumberPairs() int {
	tt := *t
	var n int
	for i := 0; i < len(tt)-1; i++ {
		j := i + 1
		if tt[i].Number == tt[j].Number {
			n++
		}
	}
	return n
}

func (t *Tiles) ColorPairs() [][]int {
	tt := *t
	var pairs [][]int
	var pair []int
	for i := 0; i < len(tt)-1; i++ {
		j := i + 1
		if tt[i].Color == tt[j].Color {
			if pair == nil {
				pair = append(pair, i)
			}
			pair = append(pair, j)
			if j != len(tt)-1 {
				continue
			}
		}
		if len(pair) != 0 {
			pairs = append(pairs, pair)
			pair = nil
		}
	}
	return pairs
}

func (t *Tiles) Center() int {
	tt := *t
	if len(tt) == 0 {
		return 0
	}
	return tt[len(tt)/2].Number
}

func (t *Tiles) Serial() []int {
	tt := *t
	var indexes []int
	var ongoing bool
	for i := 0; i < len(tt)-1; i++ {
		j := i + 1
		if tt[i].Number+1 == tt[j].Number {
			if ongoing != true {
				indexes = append(indexes, i)
			}
			indexes = append(indexes, j)
			continue
		}
		ongoing = false
	}
	return indexes
}

func (t *Tiles) RedTiles() int {
	tt := *t
	var n int
	for _, tile := range tt {
		if tile.Color == Red {
			n++
		}
	}
	return n
}

func (t *Tiles) BlueTiles() int {
	tt := *t
	var n int
	for _, tile := range tt {
		if tile.Color == Blue {
			n++
		}
	}
	return n
}

func (t *Tiles) NumberWhere(n int) []int {
	tt := *t
	tt.Sort()
	var indexes []int
	for i, tile := range tt {
		if tile.Number == n {
			indexes = append(indexes, i)
		}
	}
	return indexes
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
