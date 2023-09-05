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

// Migration triggers the relocation of all instance of entity.
func (l *Labyrinth) Migration(e entity) {
	existing := make([]int, 0)
	for i := 0; i < len(l.rooms); i++ {
		if l.rooms[i].hasEntity(e) {
			existing = append(existing, i)
		}
		if len(existing) == l.levels[l.curLevel].setup.nbTermite {
			break
		}
	}

	for i := 0; i < len(existing); i++ {
		r := l.randomRoom(withoutKeyItem(), withoutHazard())
		l.rooms[r].addEntity(e)
		if l.debug {
			fmt.Printf("%s relocated to %d\n", e, l.rooms[r].fakeID)
		}
	}

	// erase previous locations
	for _, i := range existing {
		l.rooms[i].removeEntity(e)
	}
}
