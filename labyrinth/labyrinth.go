// Package labyrinth handles the labyrinth and the movements of its inhabitants.
package labyrinth

import (
	"fmt"
	"math/rand"
)

const (
	NbRooms  = 20
	RandRoom = NbRooms - 1
)

// Room is a vertex of the dodecahedron.
type Room struct {
	ID        int
	Neighbors []int
}

// Labyrinth is the collection of Rooms making up the dodecahedron.
type Labyrinth struct {
	Rooms []Room

	// locations
	player      int
	arrow       int
	arrowTravel int
	wumpus      int
	bats        []int
	pits        []int
}

// NewLabyrinth returns an initialized dodecahedron Labyrinth.
func NewLabyrinth() Labyrinth {
	l := Labyrinth{
		bats: make([]int, 0, 2),
		pits: make([]int, 0, 2),
		// there is probably a way to do this mathematically but is it worth it ?
		Rooms: []Room{
			{ID: 0, Neighbors: []int{1, 5, 4}},
			{ID: 1, Neighbors: []int{0, 7, 2}},
			{ID: 2, Neighbors: []int{1, 9, 3}},
			{ID: 3, Neighbors: []int{2, 11, 4}},
			{ID: 4, Neighbors: []int{3, 13, 0}},
			{ID: 5, Neighbors: []int{0, 14, 6}},
			{ID: 6, Neighbors: []int{5, 16, 7}},
			{ID: 7, Neighbors: []int{1, 6, 8}},
			{ID: 8, Neighbors: []int{7, 9, 17}},
			{ID: 9, Neighbors: []int{2, 8, 10}},
			{ID: 10, Neighbors: []int{9, 11, 18}},
			{ID: 11, Neighbors: []int{10, 3, 12}},
			{ID: 12, Neighbors: []int{19, 11, 13}},
			{ID: 13, Neighbors: []int{14, 12, 4}},
			{ID: 14, Neighbors: []int{13, 5, 15}},
			{ID: 15, Neighbors: []int{14, 19, 16}},
			{ID: 16, Neighbors: []int{6, 15, 17}},
			{ID: 17, Neighbors: []int{16, 8, 18}},
			{ID: 18, Neighbors: []int{10, 17, 19}},
			{ID: 19, Neighbors: []int{12, 15, 18}},
		},
	}

	l.Init()
	return l
}

// Init randomly places the player, wumpus, pits and bats.
func (l *Labyrinth) Init() {
	randRooms := make([]int, NbRooms)
	for i := range randRooms {
		randRooms[i] = i
	}

	rand.Shuffle(len(randRooms), func(i, j int) {
		randRooms[i], randRooms[j] = randRooms[j], randRooms[i]
	})

	// place pits & bats in distinct locations
	l.pits = randRooms[0:2]
	l.bats = randRooms[2:4]

	// place the Wumpus anywhere
	l.wumpus = rand.Intn(RandRoom)

	// place the player in a location distinct from hazards
	for l.player = randRooms[rand.Intn((RandRoom)-4)+4]; l.player == l.wumpus; {
	}
}

func (l *Labyrinth) Player() int {
	return l.player
}

func (l *Labyrinth) HasPlayer(n int) bool {
	return n == l.player
}

func (l *Labyrinth) HasBat(n int) bool {
	return n == l.bats[0] || n == l.bats[1]
}

func (l *Labyrinth) BatsNearby() bool {
	for _, i := range l.Rooms[l.player].Neighbors {
		if l.HasBat(i) {
			return true
		}
	}
	return false
}

func (l *Labyrinth) HasPit(n int) bool {
	return n == l.pits[0] || n == l.pits[1]
}

func (l *Labyrinth) PitNearby() bool {
	for _, i := range l.Rooms[l.player].Neighbors {
		if l.HasPit(i) {
			return true
		}
	}
	return false
}

func (l *Labyrinth) HasWumpus(n int) bool {
	return n == l.wumpus
}

func (l *Labyrinth) WumpusNearby() bool {
	for _, i := range l.Rooms[l.player].Neighbors {
		if i == l.wumpus {
			return true
		}
	}
	return false
}

// ActivateBat teleports the player to a different room.
func (l *Labyrinth) ActivateBat() int {
	var move int
	for move = rand.Intn(RandRoom); move == l.player; {
	}

	l.player = move
	return l.player
}

// FoundWumpus has a 1/2 chance of killing the player.
// In any case the Wumpus will relocate.
func (l *Labyrinth) FoundWumpus() (killed bool) {
	// move the wumpus to another room
	var move int
	for move = rand.Intn(RandRoom); move == l.wumpus; {
	}

	l.wumpus = move

	return rand.Intn(2) == 1
}

// StartleWumpus has a 1/2 chance of making the Wumpus move.
func (l *Labyrinth) StartleWumpus() bool {
	if rand.Intn(2) == 1 {
		var move int
		for move = rand.Intn(RandRoom); move == l.wumpus; {
		}

		l.wumpus = move

		return true
	}

	return false
}

// Arrow current location of the arrow.
func (l *Labyrinth) Arrow() int {
	return l.arrow
}

// FireArrow sets the arrow position to that of the player and reset its travel capacity.
func (l *Labyrinth) FireArrow() {
	l.arrow = l.player
	l.arrowTravel = 5
}

func (l *Labyrinth) PowerRemaining() int {
	return l.arrowTravel
}

// MoveArrow handle the location and travel of the arrow, reducing its capacity by one.
func (l *Labyrinth) MoveArrow(target int) {
	if target == l.Rooms[l.arrow].Neighbors[0] ||
		target == l.Rooms[l.arrow].Neighbors[1] ||
		target == l.Rooms[l.arrow].Neighbors[2] {
		l.arrow = target
	} else {
		// invalid destination, we move the arrow at random.
		l.arrow = l.Rooms[l.arrow].Neighbors[rand.Intn(3)]
	}

	l.arrowTravel--
}

// TryMovePlayer moves the player if the position is valid.
func (l *Labyrinth) TryMovePlayer(target int) bool {
	if target == l.Rooms[l.player].Neighbors[0] ||
		target == l.Rooms[l.player].Neighbors[1] ||
		target == l.Rooms[l.player].Neighbors[2] {
		l.player = target

		return true
	}

	return false
}

func (l *Labyrinth) GetFmtNeighbors(n int) string {
	return fmt.Sprintf("%d, %d, %d",
		l.Rooms[n].Neighbors[0]+1,
		l.Rooms[n].Neighbors[1]+1,
		l.Rooms[n].Neighbors[2]+1,
	)
}
