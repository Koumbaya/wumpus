package labyrinth

import (
	"fmt"
)

// Arrow current location of the arrow.
func (l *Labyrinth) Arrow() int {
	return l.arrow
}

// FireArrow sets the arrow position to that of the player.
func (l *Labyrinth) FireArrow() {
	l.arrow = l.playerLoc
}

// ArrowPOV return the shuffled arrow location.
func (l *Labyrinth) ArrowPOV() int {
	return l.rooms[l.arrow].fakeID
}

// MoveArrow handle the location and travel of the arrow, reducing its capacity by one.
func (l *Labyrinth) MoveArrow(fakeTarget int) {
	target := l.fakeIDs[fakeTarget]
	if l.validDestination(l.arrow, target) {
		l.arrow = target
	} else {
		// invalid destination, we move the arrow at random between the edges.
		if l.debug {
			fmt.Printf("fakeTarget %d, real %d\n", fakeTarget, l.fakeIDs[fakeTarget])
		}
		l.arrow = l.randomDest(l.arrow)
	}
}
