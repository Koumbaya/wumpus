package main

import (
	"flag"

	"github.com/koumbaya/wumpus/dialogues"
	"github.com/koumbaya/wumpus/game"
	"github.com/koumbaya/wumpus/labyrinth"
)

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

	l := labyrinth.NewLabyrinth(advanced, debug)
	g := game.NewGame(l, dialogues.NewPrinter(nodelay), arrows, advanced)
	g.Loop()
}
