package labyrinth

import (
	"fmt"
	"math/rand"
	"strconv"
)

// HasClue checks if a clue for a given location exists and has been found.
// If not mark it as found.
func (l *Labyrinth) HasClue(n int) bool {
	for loc, found := range l.clues {
		if n == loc {
			if !found {
				l.clues[n] = true // mark as found
				return true
			}
		}
	}

	return false
}

// GetClue returns a fresh clue about a random subject.
// Can give new clue if the wumpus moved.
// TODO : move dialogues subjects outside.
func (l *Labyrinth) GetClue() (loc int, subject string) {
	nbEntities := 5
	if l.wump3 {
		nbEntities++
	}
	for {
		n := rand.Intn(nbEntities)
		switch n {
		case 0: // pits
			loc = l.shuffled[l.pits[rand.Intn(len(l.pits))]] + 1
			sub := []string{"a pit", "a hole in the ground", "the abyss"}
			subject = sub[rand.Intn(len(sub))]
		case 1:
			loc = l.shuffled[l.bats[rand.Intn(len(l.bats))]] + 1
			sub := []string{"bats", "winged creatures", "gargoyles"}
			subject = sub[rand.Intn(len(sub))]
		case 2:
			loc = l.shuffled[l.wumpus] + 1
			subject = "the Wumpus"
		case 3:
			loc = l.shuffled[l.key] + 1
			subject = "a key"
		case 4:
			loc = l.shuffled[l.door] + 1
			subject = "a door"
		case 5:
			loc = l.shuffled[l.termites] + 1
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

func (l *Labyrinth) GetCluesLocFmt() string {
	var output string
	for i := range l.clues {
		loc := l.shuffled[i]
		output += strconv.Itoa(loc+1) + " "
	}
	return output
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
			l.shuffled[r]+1,
			l.shuffled[l.rooms[r].edges[rand.Intn(len(l.rooms[r].edges))]]+1,
		)
	}
	return output
}

// FoundRepel check if repel is at current location and mark as found, return true only the first time.
func (l *Labyrinth) FoundRepel() bool {
	if l.player == l.repel && !l.repelFound {
		l.repelFound = true
		return true
	}
	return false
}
