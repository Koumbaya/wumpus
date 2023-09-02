package labyrinth

import "fmt"

func (l *Labyrinth) HasPit(n int) bool {
	return n == l.pits[0] || n == l.pits[1]
}

func (l *Labyrinth) PitNearby() bool {
	for _, i := range l.rooms[l.player].edges {
		if l.HasPit(i) {
			return true
		}
	}
	return false
}

// Earthquake changes the pits' location.
func (l *Labyrinth) Earthquake() {
	l.pits[0] = randNotEqual(0, len(l.rooms), l.player, l.door, l.key, l.pits[0], l.pits[1], l.bats[0], l.bats[1])
	l.pits[1] = randNotEqual(0, len(l.rooms), l.player, l.door, l.key, l.pits[0], l.pits[1], l.bats[0], l.bats[1])
	if l.debug {
		fmt.Printf("pits %d %d\n", l.shuffled[l.pits[0]]+1, l.shuffled[l.pits[1]]+1)
	}
}
