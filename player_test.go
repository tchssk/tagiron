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

func TestPlayersFindByID(t *testing.T) {
	var (
		playerFoo = NewPlayer("foo", "FOO")
		playerBar = NewPlayer("bar", "BAR")
		playerBaz = NewPlayer("baz", "BAR")
	)
	cases := []struct {
		name     string
		id       string
		expected Player
	}{
		{
			name:     "found",
			id:       "FOO",
			expected: playerFoo,
		},
		{
			name:     "not found",
			id:       "QUX",
			expected: Player{},
		},
	}
	players := Players{
		playerFoo,
		playerBar,
		playerBaz,
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			player := players.FindByID(tc.id)
			if actual := player; actual.Name != tc.expected.Name || actual.ID != tc.expected.ID {
				t.Errorf("got %#v, expected %#v", actual, tc.expected)
			}
		})
	}
}
