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
	clean    bool
	wump3    bool
	level    int
)

func init() {
	flag.BoolVar(&nodelay, "nodelay", false, "Disable the fake delay when printing text")
	flag.BoolVar(&arrows, "arrows", false, "Gives infinite arrows")
	flag.BoolVar(&advanced, "advanced", true, "Experimental, expanded game")
	flag.BoolVar(&debug, "debug", false, "Print location of things for debug purpose")
	flag.BoolVar(&clean, "clean", false, "Remove symbols and colors from terminal output")
	flag.BoolVar(&wump3, "wump3", true, "Features from wumpus III")
	flag.IntVar(&level, "level", 1, "Start at a specific level")
}

func main() {
	flag.Parse()

	l := labyrinth.NewLabyrinth(advanced, debug, level)
	g := game.NewGame(l, dialogues.NewPrinter(nodelay, clean), arrows, advanced, wump3)
	g.Loop()
}
