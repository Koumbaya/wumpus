package main

import (
	"flag"
	"time"

	"github.com/koumbaya/wumpus/dialogues"
	"github.com/koumbaya/wumpus/game"
	"github.com/koumbaya/wumpus/labyrinth"
)

var nodelay bool

func init() {
	flag.BoolVar(&nodelay, "nodelay", false, "Disable the fake delay when printing text")
}

func main() {
	flag.Parse()
	delay := 15 * time.Millisecond
	if nodelay {
		delay = 0
	}

	l := labyrinth.NewLabyrinth()
	g := game.NewGame(l, dialogues.NewPrinter(delay))
	g.Loop()
}
