package labyrinth

import (
	"fmt"
	"math/rand"
)

// ActivateBat teleports the player to a different room.
func (l *Labyrinth) ActivateBat() int {
	l.rooms[l.playerLoc].player = false
	n := rand.Intn(len(l.rooms))
	l.rooms[n].player = true
	l.playerLoc = n
	return l.rooms[n].fakeID
}

// BatMigration changes the bats' location.
func (l *Labyrinth) BatMigration() {
	// save current location
	existing := make([]int, 0, nbBats)
	for i := 0; i < len(l.rooms); i++ {
		if l.rooms[i].bat {
			existing = append(existing, i)
		}
		if len(existing) == nbBats {
			break
		}
	}

	for i := 0; i < nbBats; i++ {
		r := l.randomRoom(withoutKeyItem(), withoutHazard())
		l.rooms[r].bat = true
		if l.debug {
			fmt.Printf("bat relocated to %d\n", l.rooms[r].fakeID)
		}
	}

	// erase previous locations
	for _, i := range existing {
		l.rooms[i].bat = false
	}
}
