package main

import (
	"reflect"
	"testing"
)

func TestTilesSort(t *testing.T) {
	cases := []struct {
		name     string
		tiles    Tiles
		expected Tiles
	}{
		{
			name: "all tiles",
			tiles: Tiles{
				{Blue, 1},
				{Blue, 0},
				{Blue, 2},
				{Blue, 3},
				{Red, 4},
				{Yellow, 5},
				{Blue, 4},
				{Red, 6},
				{Red, 7},
				{Blue, 6},
				{Red, 0},
				{Red, 2},
				{Red, 8},
				{Blue, 7},
				{Blue, 8},
				{Red, 3},
				{Red, 9},
				{Yellow, 5},
				{Blue, 9},
				{Red, 1},
			},
			expected: NewTiles(),
		},
		{
			name: "5 tiles",
			tiles: Tiles{
				{Red, 8},
				{Blue, 0},
				{Yellow, 5},
				{Blue, 8},
				{Red, 3},
			},
			expected: Tiles{
				{Blue, 0},
				{Red, 3},
				{Yellow, 5},
				{Red, 8},
				{Blue, 8},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.tiles.Sort()
			actual := tc.tiles
			for i, v := range actual {
				if v != tc.expected[i] {
					t.Errorf("got %#v, expected %#v at index %d", v, tc.expected[i], i)
				}
			}
		})
	}
}

func TestTilesRemove(t *testing.T) {
	cases := []struct {
		name     string
		i        int
		j        int
		expected Tiles
		removeed Tiles
	}{
		{
			name: "remove 1 tile from the head",
			i:    0,
			j:    1,
			expected: Tiles{
				{Red, 0},
			},
			removeed: Tiles{
				{Blue, 0},
				{Red, 1},
				{Blue, 1},
				{Red, 2},
				{Blue, 2},
				{Red, 3},
				{Blue, 3},
				{Red, 4},
				{Blue, 4},
				{Yellow, 5},
				{Yellow, 5},
				{Red, 6},
				{Blue, 6},
				{Red, 7},
				{Blue, 7},
				{Red, 8},
				{Blue, 8},
				{Red, 9},
				{Blue, 9},
			},
		},
		{
			name: "remove 2 tiles from the tail",
			i:    17,
			j:    19,
			expected: Tiles{
				{Red, 9},
				{Blue, 9},
			},
			removeed: Tiles{
				{Blue, 0},
				{Red, 1},
				{Blue, 1},
				{Red, 2},
				{Blue, 2},
				{Red, 3},
				{Blue, 3},
				{Red, 4},
				{Blue, 4},
				{Yellow, 5},
				{Yellow, 5},
				{Red, 6},
				{Blue, 6},
				{Red, 7},
				{Blue, 7},
				{Red, 8},
				{Blue, 8},
			},
		},
		{
			name: "remove 3 tiles from 2nd from the head",
			i:    1,
			j:    4,
			expected: Tiles{
				{Red, 1},
				{Blue, 1},
				{Red, 2},
			},
			removeed: Tiles{
				{Blue, 0},
				{Blue, 2},
				{Red, 3},
				{Blue, 3},
				{Red, 4},
				{Blue, 4},
				{Yellow, 5},
				{Yellow, 5},
				{Red, 6},
				{Blue, 6},
				{Red, 7},
				{Blue, 7},
				{Red, 8},
				{Blue, 8},
			},
		},
		{
			name: "remove 4 tiles from 2nd from the tail",
			i:    9,
			j:    13,
			expected: Tiles{
				{Blue, 6},
				{Red, 7},
				{Blue, 7},
				{Red, 8},
			},
			removeed: Tiles{
				{Blue, 0},
				{Blue, 2},
				{Red, 3},
				{Blue, 3},
				{Red, 4},
				{Blue, 4},
				{Yellow, 5},
				{Yellow, 5},
				{Red, 6},
				{Blue, 8},
			},
		},
		{
			name: "remove 1 tile from the middle",
			i:    2,
			j:    3,
			expected: Tiles{
				{Red, 3},
			},
			removeed: Tiles{
				{Blue, 0},
				{Blue, 2},
				{Blue, 3},
				{Red, 4},
				{Blue, 4},
				{Yellow, 5},
				{Yellow, 5},
				{Red, 6},
				{Blue, 8},
			},
		},
		{
			name: "remove 10 tiles but remaining only 9 tiles",
			i:    0,
			j:    10,
			expected: Tiles{
				{Blue, 0},
				{Blue, 2},
				{Blue, 3},
				{Red, 4},
				{Blue, 4},
				{Yellow, 5},
				{Yellow, 5},
				{Red, 6},
				{Blue, 8},
			},
			removeed: Tiles{},
		},
		{
			name:     "remove but no tile",
			i:        0,
			j:        1,
			expected: Tiles{},
			removeed: Tiles{},
		},
	}
	tiles := NewTiles()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := tiles.Remove(tc.i, tc.j); len(actual) != len(tc.expected) {
				t.Errorf("expected the number of all required values to match %d got %d ", len(tc.expected), len(actual))
			} else {
				for i, v := range actual {
					if v != tc.expected[i] {
						t.Errorf("got %#v, expected %#v at index %d", v, tc.expected[i], i)
					}
				}
			}
			if len(tiles) != len(tc.removeed) {
				t.Errorf("expected the number of all required values to match %d got %d ", len(tc.removeed), len(tiles))
			} else {
				for i, v := range tiles {
					if v != tc.removeed[i] {
						t.Errorf("got %#v, expected %#v at index %d", v, tc.removeed[i], i)
					}
				}
			}
		})
	}
}

func TestTilesAdd(t *testing.T) {
	cases := []struct {
		name     string
		tiles    Tiles
		expected Tiles
	}{
		{
			name:     "add no tile",
			tiles:    Tiles{},
			expected: Tiles{},
		},
		{
			name: "add 1 tile",
			tiles: Tiles{
				{Red, 0},
			},
			expected: Tiles{
				{Red, 0},
			},
		},
		{
			name: "add 2 tiles",
			tiles: Tiles{
				{Blue, 0},
				{Red, 1},
			},
			expected: Tiles{
				{Red, 0},
				{Blue, 0},
				{Red, 1},
			},
		},
		{
			name: "add 3 tiles",
			tiles: Tiles{
				{Blue, 1},
				{Red, 2},
				{Blue, 2},
			},
			expected: Tiles{
				{Red, 0},
				{Blue, 0},
				{Red, 1},
				{Blue, 1},
				{Red, 2},
				{Blue, 2},
			},
		},
		{
			name: "add 4 tiles",
			tiles: Tiles{
				{Red, 3},
				{Blue, 3},
				{Red, 4},
				{Blue, 4},
			},
			expected: Tiles{
				{Red, 0},
				{Blue, 0},
				{Red, 1},
				{Blue, 1},
				{Red, 2},
				{Blue, 2},
				{Red, 3},
				{Blue, 3},
				{Red, 4},
				{Blue, 4},
			},
		},
		{
			name: "add 5 tiles",
			tiles: Tiles{
				{Yellow, 5},
				{Yellow, 5},
				{Red, 6},
				{Blue, 6},
				{Red, 7},
			},
			expected: Tiles{
				{Red, 0},
				{Blue, 0},
				{Red, 1},
				{Blue, 1},
				{Red, 2},
				{Blue, 2},
				{Red, 3},
				{Blue, 3},
				{Red, 4},
				{Blue, 4},
				{Yellow, 5},
				{Yellow, 5},
				{Red, 6},
				{Blue, 6},
				{Red, 7},
			},
		},
		{
			name: "add 5 tiles again",
			tiles: Tiles{
				{Blue, 7},
				{Red, 8},
				{Blue, 8},
				{Red, 9},
				{Blue, 9},
			},
			expected: Tiles{
				{Red, 0},
				{Blue, 0},
				{Red, 1},
				{Blue, 1},
				{Red, 2},
				{Blue, 2},
				{Red, 3},
				{Blue, 3},
				{Red, 4},
				{Blue, 4},
				{Yellow, 5},
				{Yellow, 5},
				{Red, 6},
				{Blue, 6},
				{Red, 7},
				{Blue, 7},
				{Red, 8},
				{Blue, 8},
				{Red, 9},
				{Blue, 9},
			},
		},
	}
	tiles := Tiles{}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tiles.Add(tc.tiles)
			if actual := tiles; len(actual) != len(tc.expected) {
				t.Errorf("expected the number of all required values to match %d got %d ", len(tc.expected), len(actual))
			} else {
				for i, v := range actual {
					if v != tc.expected[i] {
						t.Errorf("got %#v, expected %#v at index %d", v, tc.expected[i], i)
					}
				}
			}
		})
	}
}

func TestTilesSumOfAll(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"all tiles": {
			tiles:    NewTiles(),
			expected: 90,
		},
		"one color": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 10,
		},
		"two colors": {
			tiles:    Tiles{{Red, 0}, {Blue, 1}, {Red, 2}, {Blue, 3}, {Red, 4}},
			expected: 10,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.SumOfAll(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesSumOfLowerThree(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"all tiles": {
			tiles:    NewTiles(),
			expected: 1,
		},
		"two tiles": {
			tiles:    Tiles{{Red, 2}, {Blue, 3}},
			expected: 5,
		},
		"one color": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 3,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.SumOfLowerThree(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesSumOfCenterThree(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"all tiles": {
			tiles:    NewTiles(),
			expected: 14,
		},
		"one color": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 6,
		},
		"two colors": {
			tiles:    Tiles{{Red, 0}, {Blue, 1}, {Red, 2}, {Blue, 3}, {Red, 4}},
			expected: 6,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.SumOfCenterThree(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesSumOfUpperThree(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"all tiles": {
			tiles:    NewTiles(),
			expected: 26,
		},
		"two tiles": {
			tiles:    Tiles{{Red, 2}, {Blue, 3}},
			expected: 5,
		},
		"one color": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 9,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.SumOfUpperThree(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesSumOfRed(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"only red": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 10,
		},
		"only blue": {
			tiles:    Tiles{{Blue, 0}, {Blue, 2}, {Blue, 3}, {Blue, 4}, {Blue, 5}},
			expected: 0,
		},
		"two colors": {
			tiles:    Tiles{{Red, 0}, {Blue, 1}, {Red, 2}, {Blue, 3}, {Red, 4}},
			expected: 6,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.SumOfRed(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesSumOfBlue(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"only red": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 0,
		},
		"only blue": {
			tiles:    Tiles{{Blue, 0}, {Blue, 1}, {Blue, 2}, {Blue, 3}, {Blue, 4}},
			expected: 10,
		},
		"two colors": {
			tiles:    Tiles{{Red, 0}, {Blue, 1}, {Red, 2}, {Blue, 3}, {Red, 4}},
			expected: 4,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.SumOfBlue(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesDifference(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"one color": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 4,
		},
		"two colors": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Blue, 3}, {Blue, 4}},
			expected: 4,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.Difference(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesOdd(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"one color": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 2,
		},
		"two colors": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Blue, 3}, {Blue, 4}},
			expected: 2,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.Odd(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesEven(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"one color": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 3,
		},
		"two colors": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Blue, 3}, {Blue, 4}},
			expected: 3,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.Even(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesNumberPairs(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"no pair": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 0,
		},
		"one color": {
			tiles:    Tiles{{Red, 0}, {Red, 0}, {Red, 1}, {Red, 1}, {Red, 2}},
			expected: 2,
		},
		"two colors": {
			tiles:    Tiles{{Red, 0}, {Blue, 0}, {Red, 1}, {Blue, 1}, {Red, 2}},
			expected: 2,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.NumberPairs(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesColorPairs(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected [][]int
	}{
		"no pair": {
			tiles:    Tiles{{Red, 0}, {Blue, 1}, {Red, 2}, {Blue, 3}, {Red, 4}},
			expected: nil,
		},
		"all": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: [][]int{{0, 1, 2, 3, 4}},
		},
		"one pair": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Blue, 2}, {Red, 3}, {Blue, 4}},
			expected: [][]int{{0, 1}},
		},
		"two pairs": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Blue, 2}, {Blue, 3}, {Blue, 4}},
			expected: [][]int{{0, 1}, {2, 3, 4}},
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.ColorPairs(); !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesCenter(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"red": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 2,
		},
		"blue": {
			tiles:    Tiles{{Blue, 0}, {Blue, 1}, {Blue, 2}, {Blue, 3}, {Blue, 4}},
			expected: 2,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.Center(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesSerial(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected []int
	}{
		"none": {
			tiles:    Tiles{{Red, 0}, {Red, 0}, {Red, 2}, {Red, 2}, {Red, 4}},
			expected: nil,
		},
		"single match": {
			tiles:    Tiles{{Blue, 0}, {Blue, 1}, {Blue, 3}, {Blue, 5}, {Blue, 7}},
			expected: []int{0, 1},
		},
		"multiple matches": {
			tiles:    Tiles{{Red, 0}, {Blue, 1}, {Red, 3}, {Blue, 5}, {Red, 6}},
			expected: []int{0, 1, 3, 4},
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.Serial(); !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesRedTiles(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"only red": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 5,
		},
		"only blue": {
			tiles:    Tiles{{Blue, 0}, {Blue, 1}, {Blue, 2}, {Blue, 3}, {Blue, 4}},
			expected: 0,
		},
		"two colors": {
			tiles:    Tiles{{Red, 0}, {Blue, 1}, {Red, 2}, {Blue, 3}, {Red, 4}},
			expected: 3,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.RedTiles(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesBlueTiles(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		expected int
	}{
		"only red": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			expected: 0,
		},
		"only blue": {
			tiles:    Tiles{{Blue, 0}, {Blue, 1}, {Blue, 2}, {Blue, 3}, {Blue, 4}},
			expected: 5,
		},
		"two colors": {
			tiles:    Tiles{{Red, 0}, {Blue, 1}, {Red, 2}, {Blue, 3}, {Red, 4}},
			expected: 2,
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.BlueTiles(); actual != tc.expected {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}

func TestTilesNumberWhere(t *testing.T) {
	cases := map[string]struct {
		tiles    Tiles
		n        int
		expected []int
	}{
		"none": {
			tiles:    Tiles{{Red, 0}, {Red, 1}, {Red, 2}, {Red, 3}, {Red, 4}},
			n:        5,
			expected: nil,
		},
		"single match": {
			tiles:    Tiles{{Blue, 0}, {Blue, 1}, {Blue, 2}, {Blue, 3}, {Blue, 4}},
			n:        2,
			expected: []int{2},
		},
		"multiple matches": {
			tiles:    Tiles{{Red, 0}, {Blue, 0}, {Red, 1}, {Blue, 1}, {Red, 2}},
			n:        1,
			expected: []int{2, 3},
		},
	}
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			if actual := tc.tiles.NumberWhere(tc.n); !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}
