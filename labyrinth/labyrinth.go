// Package labyrinth handles the labyrinth and the movements of its inhabitants.
package labyrinth

import (
	"fmt"
	"math/rand"
	"strconv"
)

// room is a vertex of the graph.
type room struct {
	edges []int
}

type level struct {
	rooms  []room
	number int
	name   string
}

// Labyrinth is the collection of cLevel making up the dodecahedron.
type Labyrinth struct {
	// levels store all the jsons levels
	levels   map[int]level
	curLevel int
	rooms    []room // current level topology
	// visited keep track of the # of explored rooms.
	visited map[int]struct{}
	// shuffled is a nbRooms length slice with values 0-randRoom randomized.
	// those values are what are shown to the player (so that on each play through, the cavern numbers on the map are different).
	shuffled []int
	// ordered is the reverse of shuffled. It is used when taking player input to find the "real" room.
	ordered []int
	// arrowTravel keep track of how many cLevel the arrow can travel.
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
		levels:   loadLevels(),
		advanced: advanced,
		debug:    debug,
	}

	l.Init(1)
	return l
}

// Init randomly places the player, wumpus, pits and bats.
func (l *Labyrinth) Init(targetLvl int) {
	l.curLevel = targetLvl
	l.rooms = l.levels[targetLvl].rooms
	l.visited = make(map[int]struct{}, len(l.rooms))
	l.ordered = make([]int, len(l.rooms))
	randRooms := make([]int, len(l.rooms))
	for i := range randRooms {
		randRooms[i] = i
	}

	rand.Shuffle(len(randRooms), func(i, j int) {
		randRooms[i], randRooms[j] = randRooms[j], randRooms[i]
	})

	// use the randomization to give arbitrary numbers to cLevel so that each play through is unique.
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
	rWumpus := rand.Intn(len(l.rooms))
	l.wumpus = randRooms[rWumpus]

	// place the player in a location distinct from hazards
	l.player = randRooms[randNotEqual(offset, len(l.rooms), rWumpus)]

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
	for _, i := range l.rooms[l.player].edges {
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
	for _, i := range l.rooms[l.player].edges {
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
	for _, i := range l.rooms[l.player].edges {
		if i == l.wumpus {
			return true
		}
	}
	return false
}

// ActivateBat teleports the player to a different room.
func (l *Labyrinth) ActivateBat() int {
	l.player = randNotEqual(0, len(l.rooms), l.player)
	return l.player
}

// FoundWumpus has a 1/2 chance of killing the player.
// In any case the Wumpus will relocate.
func (l *Labyrinth) FoundWumpus() (killed bool) {
	// move the wumpus to another room
	l.wumpus = randNotEqual(0, len(l.rooms), l.wumpus)

	return rand.Intn(2) == 1
}

// StartleWumpus usually makes the Wumpus relocate.
func (l *Labyrinth) StartleWumpus() bool {
	if rand.Intn(4) != 0 { // 3 times out of 4 the wumpus will relocate
		l.wumpus = randNotEqual(0, len(l.rooms), l.wumpus)
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
	if l.validDestination(l.arrow, target) {
		l.arrow = target
	} else {
		// invalid destination, we move the arrow at random between the edges.
		l.arrow = l.randomDest(l.arrow)
	}

	l.arrowTravel--
}

func (l *Labyrinth) validDestination(location, target int) bool {
	for _, edge := range l.rooms[location].edges {
		if edge == target {
			return true
		}
	}
	return false
}

// randomDest chooses an edge at random for a given location
func (l *Labyrinth) randomDest(location int) int {
	return l.rooms[location].edges[rand.Intn(len(l.rooms[location].edges))]
}

// TryMovePlayer moves the player if the position is valid.
func (l *Labyrinth) TryMovePlayer(target int) bool {
	target = l.ordered[target]
	if l.validDestination(l.player, target) || l.debug /*allow teleport in debug mode*/ {
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

// GetFmtNeighbors returns the shuffled (player POV) and formatted list of outgoing edges (tunnels) for a given room.
func (l *Labyrinth) GetFmtNeighbors(room int) string {
	var output string
	for idxEdge := 0; idxEdge < len(l.rooms[room].edges); idxEdge++ {
		output += strconv.Itoa(l.shuffled[l.rooms[room].edges[idxEdge]] + 1)
		if !(idxEdge == len(l.rooms[room].edges)-1) {
			output += ", "
		}
	}
	return output
}

func (l *Labyrinth) CurrentLevel() int {
	// todo : name ?
	return l.curLevel
}

func (l *Labyrinth) HasNextLevel() bool {
	_, exist := l.levels[l.curLevel+1]
	return exist
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
