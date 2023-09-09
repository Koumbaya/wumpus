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

// FoundWumpus has a 1/2 chance of attacking the player if he's already awake.
// In any case the Wumpus will relocate to a random room, avoiding the player.
func (l *Labyrinth) FoundWumpus(awake bool) (killed bool) {
	// move the wumpus to another room
	l.relocateWumpus(true)

	if awake {
		return l.r.Intn(2) == 0
	}

	return false
}

// StartleWumpus will relocate the Wumpus to any random room.
// Might result in the Wumpus going into the player room.
func (l *Labyrinth) StartleWumpus() {
	l.relocateWumpus(false)
}

// MigrateWumpus move the wumpus 1 adjacent cavern randomly.
func (l *Labyrinth) MigrateWumpus() {
	var loc int
	for i, r := range l.rooms {
		if r.hasEntity(Wumpus) {
			loc = i
		}
	}

	if len(l.rooms[loc].edges) == 0 {
		return // edge case for custom labyrinth, doesn't make much sense, means dead end.
	}

	n := l.r.Intn(len(l.rooms[loc].edges))
	l.rooms[loc].removeEntity(Wumpus)
	l.rooms[n].addEntity(Wumpus)

	if l.debug {
		fmt.Printf("wumpus relocated to %d\n", l.rooms[n].fakeID)
	}
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
