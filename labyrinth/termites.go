package labyrinth

import "fmt"

func (l *Labyrinth) HasTermites(n int) bool {
	return n == l.termites
}

func (l *Labyrinth) TermitesNearby() bool {
	for _, i := range l.rooms[l.player].edges {
		if l.HasTermites(i) {
			return true
		}
	}
	return false
}

func (l *Labyrinth) TermitesMigration() {
	l.termites = randNotEqual(0, len(l.rooms), l.player, l.termites, l.door, l.key, l.bats[0], l.bats[1], l.pits[0], l.pits[1])
	if l.debug {
		fmt.Printf("termites %d\n", l.shuffled[l.termites]+1)
	}
}
