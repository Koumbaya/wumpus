package labyrinth

import "strconv"

// Arrow current location of the arrow.
func (l *Labyrinth) Arrow() int {
	return l.arrow
}

// ArrowPOV return the shuffled arrow location.
func (l *Labyrinth) ArrowPOV() int {
	return l.shuffled[l.arrow] + 1
}

// FireArrow sets the arrow position to that of the player and reset its travel capacity.
func (l *Labyrinth) FireArrow(input string) {
	p, err := strconv.Atoi(input)
	if err != nil || p > 5 || p == 0 {
		p = 5
	}

	l.arrow = l.player
	l.arrowTravel = p
}

func (l *Labyrinth) PowerRemaining() int {
	return l.arrowTravel
}

// MoveArrow handle the location and travel of the arrow, reducing its capacity by one.
func (l *Labyrinth) MoveArrow(target int) {
	target = l.ordered[target]
	if l.validDestination(l.arrow, target) {
		l.arrow = target
	} else {
		// invalid destination, we move the arrow at random between the edges.
		l.arrow = l.randomDest(l.arrow)
	}

	l.arrowTravel--
}
