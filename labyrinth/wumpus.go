package labyrinth

import (
	"fmt"
)

// Wumpus returns the shuffled location of the wumpus.
func (l *Labyrinth) Wumpus() int {
	for _, r := range l.rooms {
		if r.hasEntity(Wumpus) {
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

	return l.r.Intn(2) == 0
}

// StartleWumpus usually makes the Wumpus relocate.
func (l *Labyrinth) StartleWumpus() bool {
	if l.r.Intn(4) != 0 { // 3 times out of 4 the wumpus will relocate
		l.relocateWumpus(false)
		return true
	}

	return false
}

func (l *Labyrinth) SleepwalkWumpus() {
	l.relocateWumpus(true)
}

func (l *Labyrinth) relocateWumpus(avoidPlayer bool) {
	// save current location
	existing := make([]int, 0, nbWumpus)
	for i := 0; i < len(l.rooms); i++ {
		if l.rooms[i].hasEntity(Wumpus) {
			existing = append(existing, i)
		}
		if len(existing) == nbWumpus {
			break
		}
	}
	var cond []filterFunc
	if avoidPlayer {
		cond = []filterFunc{withoutEntity(Wumpus), withoutEntity(Player)}
	} else {
		cond = []filterFunc{withoutEntity(Wumpus)}
	}
	for i := 0; i < nbWumpus; i++ {
		r := l.randomRoom(cond...)
		l.rooms[r].addEntity(Wumpus)
		if l.debug {
			fmt.Printf("wumpus relocated to %d\n", l.rooms[r].fakeID)
		}
	}

	// erase previous locations
	for _, i := range existing {
		l.rooms[i].removeEntity(Wumpus)
	}
}
