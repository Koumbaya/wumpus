package labyrinth

import (
	"fmt"
	"math/rand"
)

// Wumpus returns the shuffled location of the wumpus.
func (l *Labyrinth) Wumpus() int {
	for _, r := range l.rooms {
		if r.wumpus {
			return r.fakeID
		}
	}
	return 0
}

// FoundWumpus has a 1/2 chance of killing the player.
// In any case the Wumpus will relocate.
func (l *Labyrinth) FoundWumpus() (killed bool) {
	// move the wumpus to another room
	l.relocateWumpus(true)

	return rand.Intn(2) == 0
}

// StartleWumpus usually makes the Wumpus relocate.
func (l *Labyrinth) StartleWumpus() bool {
	if rand.Intn(4) != 0 { // 3 times out of 4 the wumpus will relocate
		l.relocateWumpus(false)
		return true
	}

	return false
}

func (l *Labyrinth) SleepwalkWumpus() {
	l.relocateWumpus(true)
}

func (l *Labyrinth) relocateWumpus(avoidPlayer bool) {
	old := 0
	for i := 0; i < len(l.rooms); i++ {
		if l.rooms[i].wumpus {
			l.rooms[i].wumpus = false
			old = i
		}
	}

	for {
		n := rand.Intn(len(l.rooms))
		if n == old || avoidPlayer && l.rooms[n].player {
			continue
		}
		l.rooms[n].wumpus = true
		if l.debug {
			fmt.Printf("wumpus %d\n", l.rooms[n].fakeID)
			for i := 0; i < len(l.rooms[n].edges); i++ {
				fmt.Printf("adjacent caves %d\n", l.rooms[l.rooms[n].edges[i]].fakeID)
			}
		}
		break
	}
}
