package labyrinth

import (
	"fmt"
	"math/rand"
	"strconv"
)

// GetClue returns a fresh clue about a random subject.
// Can give new clue if the wumpus/pits/bats moved.
// TODO : move dialogues subjects outside.
func (l *Labyrinth) GetClue(clueLoc int) (loc int, subject string) {
	l.rooms[clueLoc].clue = false //remove the clue from the room
	nbEntities := 5
	if l.wump3 {
		nbEntities++ // for termites
	}
	for {
		n := rand.Intn(nbEntities)
		switch n {
		case 0: // pits
			loc = l.rooms[l.getRoom(withEntity(Pit))].fakeID
			sub := []string{"a pit", "a hole in the ground", "the abyss"}
			subject = sub[rand.Intn(len(sub))]
		case 1:
			loc = l.rooms[l.getRoom(withEntity(Bat))].fakeID
			sub := []string{"bats", "winged creatures", "gargoyles"}
			subject = sub[rand.Intn(len(sub))]
		case 2:
			loc = l.rooms[l.getRoom(withEntity(Wumpus))].fakeID
			subject = "the Wumpus"
		case 3:
			loc = l.rooms[l.getRoom(withEntity(Key))].fakeID
			subject = "a key"
		case 4:
			loc = l.rooms[l.getRoom(withEntity(Door))].fakeID
			subject = "a door"
		case 5:
			loc = l.rooms[l.getRoom(withEntity(Termite))].fakeID
			sub := []string{"insects that eat wood", "termites", "a colony of wood eater"}
			subject = sub[rand.Intn(len(sub))]
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
	n := rand.Intn(3) //how many connections to display
	n++               // at least 1
	output += "\n"
	for i := 0; i < n; i++ {
		r := rand.Intn(len(l.rooms))
		output += fmt.Sprintf("%d --> %d\n",
			l.rooms[r].fakeID,
			l.rooms[l.rooms[r].edges[rand.Intn(len(l.rooms[r].edges))]].fakeID,
		)
	}
	return output
}

// FoundRepel check if repel is at current location and mark as found, return true only the first time.
func (l *Labyrinth) FoundRepel() bool {
	if l.rooms[l.playerLoc].repel {
		l.rooms[l.playerLoc].repel = false
		return true
	}
	return false
}

func (l *Labyrinth) FoundRope() bool {
	if l.rooms[l.playerLoc].rope {
		l.rooms[l.playerLoc].rope = false
		return true
	}
	return false
}
