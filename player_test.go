package main

import (
	"testing"
)

func TestPlayersAdd(t *testing.T) {
	var (
		playerFoo = NewPlayer("foo", "FOO")
		playerBar = NewPlayer("bar", "BAR")
	)
	cases := []struct {
		name     string
		player   Player
		expected Players
	}{
		{
			name:   "add 1 player",
			player: playerFoo,
			expected: Players{
				playerFoo,
			},
		},
		{
			name:   "add 1 player again",
			player: playerBar,
			expected: Players{
				playerFoo,
				playerBar,
			},
		},
	}
	players := Players{}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			players.Add(tc.player)
			if actual := players; len(actual) != len(tc.expected) {
				t.Errorf("expected the number of all required values to match %d got %d ", len(tc.expected), len(actual))
			} else {
				for i, v := range actual {
					if v.Name != tc.expected[i].Name || v.ID != tc.expected[i].ID {
						t.Errorf("got %#v, expected %#v at index %d", v, tc.expected[i], i)
					}
				}
			}
		})
	}
}
