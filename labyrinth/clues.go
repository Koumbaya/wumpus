package labyrinth

import (
	"fmt"
	"strconv"
)

// GetClue returns a fresh clue about a random subject.
// Can give new clue if the wumpus/pits/bats moved.
// TODO : move dialogues subjects outside.
func (l *Labyrinth) GetClue(clueLoc int) (loc int, subject string) {
	l.rooms[clueLoc].removeEntity(Clue) //remove the clue from the room
	nbEntities := 5
	if l.wump3 {
		nbEntities++ // for termites
	}
	for {
		n := l.r.Intn(nbEntities)
		switch n {
		case 0: // pits
			loc = l.rooms[l.randomRoom(withEntity(Pit))].fakeID
			sub := []string{"a pit", "a hole in the ground", "the abyss"}
			subject = sub[l.r.Intn(len(sub))]
		case 1:
			loc = l.rooms[l.randomRoom(withEntity(Bat))].fakeID
			sub := []string{"bats", "winged creatures", "gargoyles"}
			subject = sub[l.r.Intn(len(sub))]
		case 2:
			loc = l.rooms[l.randomRoom(withEntity(Wumpus))].fakeID
			subject = "the Wumpus"
		case 3:
			loc = l.rooms[l.randomRoom(withEntity(Key))].fakeID
			subject = "a key"
		case 4:
			loc = l.rooms[l.randomRoom(withEntity(Door))].fakeID
			subject = "a door"
		case 5:
			loc = l.rooms[l.randomRoom(withEntity(Termite))].fakeID
			sub := []string{"insects that eat wood", "termites", "a colony of wood eater"}
			subject = sub[l.r.Intn(len(sub))]
		}
		key := subject + strconv.Itoa(loc)
		if _, found := l.cluesGiven[key]; found {
			continue // this specific clue was already given
		}
		l.cluesGiven[key] = struct{}{} // store this specific clue as given
		return loc, subject
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
		output += fmt.Sprintf("%d --> %d\n",
			l.rooms[r].fakeID,
			l.rooms[l.rooms[r].edges[l.r.Intn(len(l.rooms[r].edges))]].fakeID,
		)
	}
	return output
}

// FoundRepel check if repel is at current location and mark as found, return true only the first time.
func (l *Labyrinth) FoundRepel() bool {
	if l.rooms[l.playerLoc].hasEntity(Repel) {
		l.rooms[l.playerLoc].removeEntity(Repel)
		return true
	}
	return false
}

func (l *Labyrinth) FoundRope() bool {
	if l.rooms[l.playerLoc].hasEntity(Rope) {
		l.rooms[l.playerLoc].removeEntity(Rope)
		return true
	}
	return false
}

func (l *Labyrinth) FoundShield() bool {
	if l.rooms[l.playerLoc].hasEntity(Shield) {
		l.rooms[l.playerLoc].removeEntity(Shield)
		return true
	}
	return false
}
