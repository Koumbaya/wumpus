package labyrinth

import (
	"fmt"
	"strconv"

	. "github.com/koumbaya/wumpus/model"
)

// GetClue returns a fresh clue about a random subject.
// Can give new clue if the wumpus/pits/bats moved.
// TODO : move dialogues subjects outside.
func (l *Labyrinth) GetClue(clueLoc int) (loc int, e Entity) {
	l.rooms[clueLoc].removeEntity(Clue) //remove the clue from the room
	possibleClues := []Entity{Pit, Bat, Wumpus, Key, Door, Termite}
	for {
		n := l.r.Intn(len(possibleClues))
		loc = l.rooms[l.randomRoom(withEntity(possibleClues[n]))].fakeID
		if loc == -1 {
			continue
		}

		key := string(possibleClues[n]) + strconv.Itoa(loc)
		if _, found := l.cluesGiven[key]; found {
			continue // this specific clue was already given
		}

		l.cluesGiven[key] = struct{}{} // store this specific clue as given
		return loc, possibleClues[n]
	}
}

// GetFmtMap returns a random (formatted) partial map.
// "maps" don't have locations and are not unique.
func (l *Labyrinth) GetFmtMap() (output string) {
	n := l.r.Intn(3) //how many connections to display
	n++              // at least 1
	output += "\n"
	for i := 0; i < n; i++ {
		r := l.r.Intn(len(l.rooms))
		output += fmt.Sprintf("%d ⟶ %d\n",
			l.rooms[r].fakeID,
			l.rooms[l.rooms[r].edges[l.r.Intn(len(l.rooms[r].edges))]].fakeID,
		)
	}
	return output
}

// FoundObject check if Entity is at current location and mark as found, return true only the first time.
func (l *Labyrinth) FoundObject(e Entity) bool {
	if l.rooms[l.playerLoc].hasEntity(e) {
		l.rooms[l.playerLoc].removeEntity(e)
		return true
	}
	return false
}
