package main

type Player struct {
	Name      string
	ID        string
	Tiles     Tiles
	Questions Questions
}

type Players []Player
