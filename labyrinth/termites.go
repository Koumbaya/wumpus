package labyrinth

import "fmt"

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
		r := l.randomRoom(withoutKeyItem(), withoutHazard())
		l.rooms[r].termite = true
		if l.debug {
			fmt.Printf("termite relocated to %d\n", l.rooms[r].fakeID)
		}
	}

	// erase previous locations
	for _, i := range existing {
		l.rooms[i].termite = false
	}
}
