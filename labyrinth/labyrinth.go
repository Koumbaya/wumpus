// Package labyrinth handles the labyrinth and the movements of its inhabitants.
package labyrinth

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	nbRooms = 20
)

// room is a vertex of the dodecahedron.
type room struct {
	neighbors []int
}

// Labyrinth is the collection of rooms making up the dodecahedron.
type Labyrinth struct {
	rooms []room
	// visited keep track of the # of explored rooms.
	visited map[int]struct{}
	// shuffled is a nbRooms length slice with values 0-randRoom randomized.
	// those values are what are shown to the player (so that on each play through, the cavern numbers on the map are different).
	shuffled []int
	// ordered is the reverse of shuffled. It is used when taking player input to find the "real" room.
	ordered []int
	// arrowTravel keep track of how many rooms the arrow can travel.
	arrowTravel int
	advanced    bool // experimental
	debug       bool

	// locations
	player int
	arrow  int
	wumpus int
	bats   []int
	pits   []int
	key    int // advanced
	door   int // advanced
}

// NewLabyrinth returns an initialized dodecahedron Labyrinth and game elements in their starting positions.
func NewLabyrinth(advanced, debug bool) Labyrinth {
	l := Labyrinth{
		// there is probably a way to do this mathematically but is it worth it ?
		rooms: []room{
			{neighbors: []int{1, 5, 4}},
			{neighbors: []int{0, 7, 2}},
			{neighbors: []int{1, 9, 3}},
			{neighbors: []int{2, 11, 4}},
			{neighbors: []int{3, 13, 0}},
			{neighbors: []int{0, 14, 6}},
			{neighbors: []int{5, 16, 7}},
			{neighbors: []int{1, 6, 8}},
			{neighbors: []int{7, 9, 17}},
			{neighbors: []int{2, 8, 10}},
			{neighbors: []int{9, 11, 18}},
			{neighbors: []int{10, 3, 12}},
			{neighbors: []int{19, 11, 13}},
			{neighbors: []int{14, 12, 4}},
			{neighbors: []int{13, 5, 15}},
			{neighbors: []int{14, 19, 16}},
			{neighbors: []int{6, 15, 17}},
			{neighbors: []int{16, 8, 18}},
			{neighbors: []int{10, 17, 19}},
			{neighbors: []int{12, 15, 18}},
		},
		advanced: advanced,
		debug:    debug,
	}

	l.Init()
	return l
}

// Init randomly places the player, wumpus, pits and bats.
func (l *Labyrinth) Init() {
	l.visited = make(map[int]struct{}, 20)
	l.ordered = make([]int, nbRooms)
	randRooms := make([]int, nbRooms)
	for i := range randRooms {
		randRooms[i] = i
	}

	rand.Shuffle(len(randRooms), func(i, j int) {
		randRooms[i], randRooms[j] = randRooms[j], randRooms[i]
	})

	// use the randomization to give arbitrary numbers to rooms so that each play through is unique.
	l.shuffled = randRooms // k: true value, v : rand
	for i, r := range randRooms {
		l.ordered[r] = i
	} // k: rand, v : true value

	// place pits & bats in distinct locations
	l.pits = randRooms[0:2]
	l.bats = randRooms[2:4]

	offset := 5
	if l.advanced {
		l.key = randRooms[5]
		l.door = randRooms[6]
		offset += 2
	}

	// place the Wumpus anywhere
	rWumpus := rand.Intn(nbRooms)
	l.wumpus = randRooms[rWumpus]

	// place the player in a location distinct from hazards
	l.player = randRooms[randNotEqual(offset, nbRooms, rWumpus)]

	l.visited[l.player] = struct{}{}

	if l.debug {
		l.printDebug()
	}
}

// Player return the player location.
func (l *Labyrinth) Player() int {
	return l.player
}

// PlayerPOV return the shuffled player location.
func (l *Labyrinth) PlayerPOV() int {
	return l.shuffled[l.player] + 1
}

func (l *Labyrinth) HasBat(n int) bool {
	return n == l.bats[0] || n == l.bats[1]
}

func (l *Labyrinth) BatsNearby() bool {
	for _, i := range l.rooms[l.player].neighbors {
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
	for _, i := range l.rooms[l.player].neighbors {
		if l.HasPit(i) {
			return true
		}
	}
	return false
}

// Wumpus returns the shuffled location of the wumpus.
func (l *Labyrinth) Wumpus() int {
	return l.shuffled[l.wumpus] + 1
}

func (l *Labyrinth) HasWumpus(n int) bool {
	return n == l.wumpus
}

func (l *Labyrinth) WumpusNearby() bool {
	for _, i := range l.rooms[l.player].neighbors {
		if i == l.wumpus {
			return true
		}
	}
	return false
}

// ActivateBat teleports the player to a different room.
func (l *Labyrinth) ActivateBat() int {
	l.player = randNotEqual(0, nbRooms, l.player)
	return l.player
}

// FoundWumpus has a 1/2 chance of killing the player.
// In any case the Wumpus will relocate.
func (l *Labyrinth) FoundWumpus() (killed bool) {
	// move the wumpus to another room
	l.wumpus = randNotEqual(0, nbRooms, l.wumpus)

	return rand.Intn(2) == 1
}

// StartleWumpus has a 1/2 chance of making the Wumpus move.
func (l *Labyrinth) StartleWumpus() bool {
	if rand.Intn(2) == 1 {
		l.wumpus = randNotEqual(0, nbRooms, l.wumpus)
		return true
	}

	return false
}

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
	if target == l.rooms[l.arrow].neighbors[0] ||
		target == l.rooms[l.arrow].neighbors[1] ||
		target == l.rooms[l.arrow].neighbors[2] {
		l.arrow = target
	} else {
		// invalid destination, we move the arrow at random between the neighbors.
		l.arrow = l.rooms[l.arrow].neighbors[rand.Intn(3)]
	}

	l.arrowTravel--
}

// TryMovePlayer moves the player if the position is valid.
func (l *Labyrinth) TryMovePlayer(target int) bool {
	target = l.ordered[target]
	if target == l.rooms[l.player].neighbors[0] ||
		target == l.rooms[l.player].neighbors[1] ||
		target == l.rooms[l.player].neighbors[2] ||
		l.debug /*allow teleport in debug mode*/ {
		l.player = target
		l.visited[target] = struct{}{}
		return true
	}

	return false
}

func (l *Labyrinth) Visited() int {
	return len(l.visited)
}

func (l *Labyrinth) Key() int {
	return l.key
}

func (l *Labyrinth) HasKey(n int) bool {
	return n == l.key
}

func (l *Labyrinth) Door() int {
	return l.door
}

func (l *Labyrinth) HasDoor(n int) bool {
	return n == l.door
}

func (l *Labyrinth) GetFmtNeighbors(n int) string {
	return fmt.Sprintf("%d, %d, %d",
		l.shuffled[l.rooms[n].neighbors[0]]+1,
		l.shuffled[l.rooms[n].neighbors[1]]+1,
		l.shuffled[l.rooms[n].neighbors[2]]+1,
	)
}

func (l *Labyrinth) printDebug() {
	fmt.Printf("player %d\n", l.shuffled[l.player]+1)
	fmt.Printf("pits %d %d\n", l.shuffled[l.pits[0]]+1, l.shuffled[l.pits[1]]+1)
	fmt.Printf("bats %d %d\n", l.shuffled[l.bats[0]]+1, l.shuffled[l.bats[1]]+1)
	fmt.Printf("wumpus %d\n", l.shuffled[l.wumpus]+1)
	fmt.Printf("wumpus neighboring caves %s\n", l.GetFmtNeighbors(l.wumpus))
	fmt.Printf("key %d\n", l.shuffled[l.key]+1)
	fmt.Printf("door %d\n", l.shuffled[l.door]+1)
}

func randNotEqual(min, max, different int) (x int) {
	for {
		x = rand.Intn((max)-min) + min
		if x != different {
			return x
		}
	}
}
