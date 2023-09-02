package labyrinth

import "fmt"

func (l *Labyrinth) HasBat(n int) bool {
	return n == l.bats[0] || n == l.bats[1]
}

func (l *Labyrinth) BatsNearby() bool {
	for _, i := range l.rooms[l.player].edges {
		if l.HasBat(i) {
			return true
		}
	}
	return false
}

// ActivateBat teleports the player to a different room.
func (l *Labyrinth) ActivateBat() int {
	l.player = randNotEqual(0, len(l.rooms), l.player)
	return l.shuffled[l.player] + 1
}

// BatMigration changes the bats' location.
func (l *Labyrinth) BatMigration() {
	l.bats[0] = randNotEqual(0, len(l.rooms), l.player, l.door, l.key, l.bats[0], l.bats[1], l.pits[0], l.pits[1])
	l.bats[1] = randNotEqual(0, len(l.rooms), l.player, l.door, l.key, l.bats[0], l.bats[1], l.pits[0], l.pits[1])
	if l.debug {
		fmt.Printf("bats %d %d\n", l.shuffled[l.bats[0]]+1, l.shuffled[l.bats[1]]+1)
	}
}
