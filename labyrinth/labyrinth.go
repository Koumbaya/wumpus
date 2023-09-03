// Package labyrinth handles the labyrinth and the movements of its inhabitants.
package labyrinth

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	nbClues    = 3
	nbPits     = 2
	nbBats     = 2
	nbKey      = 1
	nbDoor     = 1
	nbTermites = 1
	nbRepel    = 1
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
	wump3       bool
	debug       bool

	// locations
	player     int
	arrow      int
	wumpus     int
	bats       []int
	pits       []int
	key        int
	door       int
	termites   int
	repel      int
	repelFound bool
	// clues location and found status.
	clues map[int]bool
	// to keep track of already given clues per level.
	cluesGiven map[string]struct{}
}

// NewLabyrinth returns an initialized dodecahedron Labyrinth and game elements in their starting positions.
func NewLabyrinth(advanced, debug, wump3 bool, level int) Labyrinth {
	l := Labyrinth{
		// there is probably a way to do this mathematically but is it worth it ?
		levels:   loadLevels(),
		advanced: advanced,
		debug:    debug,
		wump3:    wump3,
	}

	if _, exist := l.levels[level]; !exist {
		level = 1
	}

	l.Init(level)
	return l
}

// Init randomly places the player, wumpus, clues, pits and bats.
func (l *Labyrinth) Init(targetLvl int) {
	l.curLevel = targetLvl
	l.rooms = l.levels[targetLvl].rooms
	l.pits = make([]int, nbPits)
	l.bats = make([]int, nbBats)
	l.visited = make(map[int]struct{}, len(l.rooms))
	l.clues = make(map[int]bool, nbClues)
	l.cluesGiven = make(map[string]struct{}, nbClues)
	l.ordered = make([]int, len(l.rooms))
	randRooms := make([]int, len(l.rooms))
	for i := range randRooms {
		randRooms[i] = i
	}

	rand.Shuffle(len(randRooms), func(i, j int) {
		randRooms[i], randRooms[j] = randRooms[j], randRooms[i]
	})

	// use the randomization to give arbitrary numbers to the caves so that each play through is unique.
	l.shuffled = randRooms // k: true value, v : rand
	for i, r := range randRooms {
		l.ordered[r] = i
	} // k: rand, v : true value

	// place pits & bats in distinct locations
	offset := 0
	copy(l.pits, randRooms[offset:offset+nbPits])
	offset += nbPits
	copy(l.bats, randRooms[offset:offset+nbBats])
	offset += nbBats
	if l.wump3 {
		l.termites = randRooms[offset]
		offset += nbTermites
	}

	if l.advanced {
		// place key/door/clues in distinct locations
		l.key = randRooms[offset]
		offset += nbKey
		l.door = randRooms[offset]
		offset += nbDoor
		for i := 0; i < nbClues; i++ {
			l.clues[randRooms[i+offset]] = false
		}

		offset += nbClues
	}

	// place the Wumpus anywhere
	idxWumpus := rand.Intn(len(l.rooms))
	l.wumpus = randRooms[idxWumpus]

	// place the player in a location distinct from hazards/clues
	l.player = randRooms[randNotEqual(offset, len(l.rooms), idxWumpus)]

	// place repel anywhere but pits or player location
	l.repel = randNotEqual(0, len(l.rooms), l.player, l.pits[0], l.pits[1])
	l.repelFound = false

	l.visited[l.player] = struct{}{}

	if l.debug {
		l.PrintDebug()
	}
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

func (l *Labyrinth) PrintDebug() {
	fmt.Printf("player %d\n", l.shuffled[l.player]+1)
	fmt.Printf("pits %d %d\n", l.shuffled[l.pits[0]]+1, l.shuffled[l.pits[1]]+1)
	fmt.Printf("bats %d %d\n", l.shuffled[l.bats[0]]+1, l.shuffled[l.bats[1]]+1)
	fmt.Printf("wumpus %d\n", l.shuffled[l.wumpus]+1)
	fmt.Printf("wumpus neighboring caves %s\n", l.GetFmtNeighbors(l.wumpus))
	fmt.Printf("key %d\n", l.shuffled[l.key]+1)
	fmt.Printf("door %d\n", l.shuffled[l.door]+1)
	fmt.Printf("clues %s\n", l.GetCluesLocFmt())
	fmt.Printf("repel %d\n", l.shuffled[l.repel]+1)
}

func randNotEqual(min, max int, exclude ...int) (x int) {
	if (max - min + 1) <= len(exclude) {
		return 0 // shouldn't happen
	}
	for {
		x = rand.Intn((max)-min) + min
		if !contains(x, exclude) {
			return x
		}
	}
}

func contains(val int, slice []int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
