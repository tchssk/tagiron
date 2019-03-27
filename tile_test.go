package main

import (
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

func TestTilesPull(t *testing.T) {
	cases := []struct {
		name     string
		i        int
		j        int
		expected Tiles
		pulled   Tiles
	}{
		{
			name: "pull 1 tile from the head",
			i:    0,
			j:    1,
			expected: Tiles{
				{Red, 0},
			},
			pulled: Tiles{
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
			name: "pull 2 tiles from the tail",
			i:    17,
			j:    19,
			expected: Tiles{
				{Red, 9},
				{Blue, 9},
			},
			pulled: Tiles{
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
			name: "pull 3 tiles from 2nd from the head",
			i:    1,
			j:    4,
			expected: Tiles{
				{Red, 1},
				{Blue, 1},
				{Red, 2},
			},
			pulled: Tiles{
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
			name: "pull 4 tiles from 2nd from the tail",
			i:    9,
			j:    13,
			expected: Tiles{
				{Blue, 6},
				{Red, 7},
				{Blue, 7},
				{Red, 8},
			},
			pulled: Tiles{
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
			name: "pull 1 tile from the middle",
			i:    2,
			j:    3,
			expected: Tiles{
				{Red, 3},
			},
			pulled: Tiles{
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
			name: "pull 10 tiles but remaining only 9 tiles",
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
			pulled: Tiles{},
		},
		{
			name:     "pull but no tile",
			i:        0,
			j:        1,
			expected: Tiles{},
			pulled:   Tiles{},
		},
	}
	tiles := NewTiles()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := tiles.Pull(tc.i, tc.j); len(actual) != len(tc.expected) {
				t.Errorf("expected the number of all required values to match %d got %d ", len(tc.expected), len(actual))
			} else {
				for i, v := range actual {
					if v != tc.expected[i] {
						t.Errorf("got %#v, expected %#v at index %d", v, tc.expected[i], i)
					}
				}
			}
			if len(tiles) != len(tc.pulled) {
				t.Errorf("expected the number of all required values to match %d got %d ", len(tc.pulled), len(tiles))
			} else {
				for i, v := range tiles {
					if v != tc.pulled[i] {
						t.Errorf("got %#v, expected %#v at index %d", v, tc.pulled[i], i)
					}
				}
			}
		})
	}
}

func TestTilesPush(t *testing.T) {
	cases := []struct {
		name     string
		tiles    Tiles
		expected Tiles
	}{
		{
			name:     "push no tile",
			tiles:    Tiles{},
			expected: Tiles{},
		},
		{
			name: "push 1 tile",
			tiles: Tiles{
				{Red, 0},
			},
			expected: Tiles{
				{Red, 0},
			},
		},
		{
			name: "push 2 tiles",
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
			name: "push 3 tiles",
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
			name: "push 4 tiles",
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
			name: "push 5 tiles",
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
			name: "push 5 tiles again",
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
			tiles.Push(tc.tiles)
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
