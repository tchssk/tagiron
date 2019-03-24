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
