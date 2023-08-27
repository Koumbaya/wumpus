package main

import (
	"github.com/koumbaya/wumpus/game"
	"github.com/koumbaya/wumpus/labyrinth"
)

func main() {
	l := labyrinth.NewLabyrinth()
	g := game.NewGame(l)
	g.Loop()
}
