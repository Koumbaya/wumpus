package labyrinth

import (
	"fmt"
	"math/rand"
)

// Earthquake changes the pits' location.
func (l *Labyrinth) Earthquake() {
	// save current location
	existing := make([]int, 0, nbPits)
	for i := 0; i < len(l.rooms); i++ {
		if l.rooms[i].pit {
			existing = append(existing, i)
		}
		if len(existing) == nbPits {
			break
		}
	}

	for i := 0; i < nbPits; i++ {
		for {
			n := rand.Intn(len(l.rooms))
			if !l.rooms[n].occupied() { // okay if there is a wumpus or a clue
				l.rooms[n].pit = true
				if l.debug {
					fmt.Printf("pit moved to %d \n", l.rooms[n].fakeID)
				}
				break
			}
		}
	}

	// erase previous locations
	for _, i := range existing {
		l.rooms[i].pit = false
	}
}
