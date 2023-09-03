package labyrinth

import (
	"fmt"
	"math/rand"
)

func (l *Labyrinth) TermitesMigration() {
	// save current location
	existing := make([]int, 0, nbTermites)
	for i := 0; i < len(l.rooms); i++ {
		if l.rooms[i].termite {
			existing = append(existing, i)
		}
		if len(existing) == nbTermites {
			break
		}
	}

	for i := 0; i < nbTermites; i++ {
		for {
			n := rand.Intn(len(l.rooms))
			if !l.rooms[n].occupied() { // okay if there is a wumpus or a clue
				l.rooms[n].termite = true
				if l.debug {
					fmt.Printf("termite moved to %d \n", l.rooms[n].fakeID)
				}
				break
			}
		}
	}

	// erase previous locations
	for _, i := range existing {
		l.rooms[i].termite = false
	}
}
