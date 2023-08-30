package main

import (
	"flag"
	"time"

	"github.com/koumbaya/wumpus/dialogues"
	"github.com/koumbaya/wumpus/game"
	"github.com/koumbaya/wumpus/labyrinth"
)

const textDelay = 15 * time.Millisecond

var (
	nodelay  bool
	arrows   bool
	advanced bool
	debug    bool
)

func init() {
	flag.BoolVar(&nodelay, "nodelay", false, "Disable the fake delay when printing text")
	flag.BoolVar(&arrows, "arrows", false, "Gives infinite arrows")
	flag.BoolVar(&advanced, "advanced", false, "Experimental, expanded game")
	flag.BoolVar(&debug, "debug", false, "Print location of things for debug purpose")
}

func main() {
	flag.Parse()
	delay := textDelay
	if nodelay {
		delay = 0
	}

	l := labyrinth.NewLabyrinth(advanced, debug)
	g := game.NewGame(l, dialogues.NewPrinter(delay), arrows, advanced)
	g.Loop()
}
