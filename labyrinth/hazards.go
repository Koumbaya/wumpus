package labyrinth

import (
	"fmt"
)

// ActivateBat teleports the player to a different room.
func (l *Labyrinth) ActivateBat() int {
	l.rooms[l.playerLoc].removeEntity(Player)
	n := l.r.Intn(len(l.rooms))
	l.rooms[n].addEntity(Player)
	l.playerLoc = n
	return l.rooms[n].fakeID
}

// BatMigration changes the bats' location.
func (l *Labyrinth) BatMigration() {
	// save current location
	existing := make([]int, 0, nbBats)
	for i := 0; i < len(l.rooms); i++ {
		if l.rooms[i].hasEntity(Bat) {
			existing = append(existing, i)
		}
		if len(existing) == nbBats {
			break
		}
	}

	for i := 0; i < nbBats; i++ {
		r := l.randomRoom(withoutKeyItem(), withoutHazard())
		l.rooms[r].addEntity(Bat)
		if l.debug {
			fmt.Printf("bat relocated to %d\n", l.rooms[r].fakeID)
		}
	}

	// erase previous locations
	for _, i := range existing {
		l.rooms[i].removeEntity(Bat)
	}
}

// Earthquake changes the pits' location.
func (l *Labyrinth) Earthquake() {
	// save current location
	existing := make([]int, 0, nbPits)
	for i := 0; i < len(l.rooms); i++ {
		if l.rooms[i].hasEntity(Pit) {
			existing = append(existing, i)
		}
		if len(existing) == nbPits {
			break
		}
	}

	for i := 0; i < nbPits; i++ {
		r := l.randomRoom(withoutKeyItem(), withoutHazard())
		l.rooms[r].addEntity(Pit)
		if l.debug {
			fmt.Printf("pit relocated to %d\n", l.rooms[r].fakeID)
		}
	}

	// erase previous locations
	for _, i := range existing {
		l.rooms[i].removeEntity(Pit)
	}
}

func (l *Labyrinth) TermitesMigration() {
	// save current location
	existing := make([]int, 0, nbTermites)
	for i := 0; i < len(l.rooms); i++ {
		if l.rooms[i].hasEntity(Termite) {
			existing = append(existing, i)
		}
		if len(existing) == nbTermites {
			break
		}
	}

	for i := 0; i < nbTermites; i++ {
		r := l.randomRoom(withoutKeyItem(), withoutHazard())
		l.rooms[r].addEntity(Termite)
		if l.debug {
			fmt.Printf("termite relocated to %d\n", l.rooms[r].fakeID)
		}
	}

	// erase previous locations
	for _, i := range existing {
		l.rooms[i].removeEntity(Termite)
	}
}
