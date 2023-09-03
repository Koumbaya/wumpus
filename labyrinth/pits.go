package labyrinth

import "fmt"

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
		r := l.randomRoom(withoutKeyItem(), withoutHazard())
		l.rooms[r].pit = true
		if l.debug {
			fmt.Printf("pit relocated to %d\n", l.rooms[r].fakeID)
		}
	}

	// erase previous locations
	for _, i := range existing {
		l.rooms[i].pit = false
	}
}
