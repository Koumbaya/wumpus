package main

import "wumpus/labyrinth"

func main() {
	l := labyrinth.NewLabyrinth()
	g := NewGame(l)
	g.Loop()
}
