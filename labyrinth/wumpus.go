package labyrinth

import (
	"fmt"
	"math/rand"
)

// Wumpus returns the shuffled location of the wumpus.
func (l *Labyrinth) Wumpus() int {
	return l.shuffled[l.wumpus] + 1
}

func (l *Labyrinth) HasWumpus(n int) bool {
	return n == l.wumpus
}

func (l *Labyrinth) WumpusNearby() bool {
	for _, i := range l.rooms[l.player].edges {
		if i == l.wumpus {
			return true
		}
	}
	return false
}

// FoundWumpus has a 1/2 chance of killing the player.
// In any case the Wumpus will relocate.
func (l *Labyrinth) FoundWumpus() (killed bool) {
	// move the wumpus to another room
	l.wumpus = randNotEqual(0, len(l.rooms), l.wumpus)

	return rand.Intn(2) == 0
}

// StartleWumpus usually makes the Wumpus relocate.
func (l *Labyrinth) StartleWumpus() bool {
	if rand.Intn(4) != 0 { // 3 times out of 4 the wumpus will relocate
		l.wumpus = randNotEqual(0, len(l.rooms), l.wumpus)
		return true
	}

	return false
}

func (l *Labyrinth) SleepwalkWumpus() {
	l.wumpus = randNotEqual(0, len(l.rooms), l.wumpus, l.player)
	if l.debug {
		fmt.Printf("wumpus %d\n", l.shuffled[l.wumpus]+1)
	}
}
