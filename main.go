package main

import (
	"flag"
	"time"

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
	seed     int64
)

func init() {
	flag.BoolVar(&nodelay, "nodelay", false, "Disable the fake delay when printing text")
	flag.BoolVar(&arrows, "arrows", false, "Gives infinite arrows")
	flag.BoolVar(&advanced, "advanced", true, "Experimental, expanded game")
	flag.BoolVar(&debug, "debug", false, "Print location of things for debug purpose")
	flag.BoolVar(&clean, "clean", false, "Remove symbols and colors from terminal output")
	flag.BoolVar(&wump3, "wump3", true, "Features from wumpus III")
	flag.IntVar(&level, "level", 1, "Start at a specific level")
	flag.Int64Var(&seed, "seed", 0, "Set a specific seed to the random functions for debug purpose")
}

func main() {
	flag.Parse()

	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	l := labyrinth.NewLabyrinth(advanced, debug, wump3, level, seed)
	g := game.NewGame(l, dialogues.NewPrinter(nodelay, clean, seed), game.Cfg{
		Seed:           seed,
		InfiniteArrows: arrows,
		Wump3:          wump3,
		Advanced:       advanced,
	})
	g.Loop()
}
