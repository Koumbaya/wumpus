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

type Entity int

const (
	Player Entity = iota
	Wumpus
	Bat
	Pit
	Termite
	Clue
	Repel
	Key
	Door
)

type entities struct {
	pit     bool
	bat     bool
	termite bool
	wumpus  bool
	player  bool
	clue    bool
	repel   bool
	key     bool
	door    bool
}

// room is a vertex of the graph.
type room struct {
	edges  []int
	fakeID int
	entities
}

// return true if the room already has one entity that can't cohabit with others.
// technically player can cohabit, but we don't want to add another entity to the player's location (usually).
func (r *room) occupied() bool {
	return r.termite || r.pit || r.player || r.clue || r.bat || r.door || r.key || r.repel
}

func (r *room) printEntities() {
	if r.pit {
		fmt.Printf("pit %d\n", r.fakeID)
	}
	if r.bat {
		fmt.Printf("bat %d\n", r.fakeID)
	}
	if r.termite {
		fmt.Printf("termite %d\n", r.fakeID)
	}
	if r.wumpus {
		fmt.Printf("wumpus %d\n", r.fakeID)
	}
	if r.player {
		fmt.Printf("player %d\n", r.fakeID)
	}
	if r.clue {
		fmt.Printf("clue %d\n", r.fakeID)
	}
	if r.repel {
		fmt.Printf("repel %d\n", r.fakeID)
	}
	if r.key {
		fmt.Printf("key %d\n", r.fakeID)
	}
	if r.door {
		fmt.Printf("door %d\n", r.fakeID)
	}
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
	// arrowTravel keep track of how many cLevel the arrow can travel.
	arrowTravel int
	advanced    bool // experimental
	wump3       bool
	debug       bool

	// locations
	playerLoc int         // keep a reference as to the player location to avoid looping at each move
	fakeIDs   map[int]int // keep a mapping of fakeID / real ids
	arrow     int
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
	l.fakeIDs = make(map[int]int)
	l.visited = make(map[int]struct{}, len(l.rooms))
	l.cluesGiven = make(map[string]struct{}, nbClues)
	randRooms := make([]int, len(l.rooms))
	for i := range randRooms {
		randRooms[i] = i
	}

	rand.Shuffle(len(randRooms), func(i, j int) {
		randRooms[i], randRooms[j] = randRooms[j], randRooms[i]
	})

	// use the randomization to give arbitrary numbers to the caves so that each play through is unique.
	for i := 0; i < len(l.rooms); i++ {
		l.rooms[i].fakeID = randRooms[i] + 1
		// reset entities on restart
		l.rooms[i].entities = entities{}
		// keep track of fakeIds for fast access
		l.fakeIDs[l.rooms[i].fakeID] = i
	}

	// place pits & bats in distinct locations
	// we use randRooms indexes so that for example both pits aren't placed on the same vertices from the graph definition each time
	offset := 0
	for i := 0; i < nbPits; i++ {
		l.rooms[randRooms[offset]].pit = true
		offset++
	}

	for i := 0; i < nbBats; i++ {
		l.rooms[randRooms[offset]].bat = true
		offset++
	}

	if l.wump3 {
		for i := 0; i < nbTermites; i++ {
			l.rooms[randRooms[offset]].termite = true
			offset++
		}
	}

	if l.advanced {
		for i := 0; i < nbKey; i++ {
			l.rooms[randRooms[offset]].key = true
			offset++
		}
		for i := 0; i < nbDoor; i++ {
			l.rooms[randRooms[offset]].door = true
			offset++
		}
		for i := 0; i < nbRepel; i++ {
			l.rooms[randRooms[offset]].repel = true
			offset++
		}
		// clues could be on the same room as key or door technically, but not the rest
		for i := 0; i < nbClues; i++ {
			l.rooms[randRooms[offset]].clue = true
			offset++
		}
	}

	// place player
	l.rooms[randRooms[offset]].player = true
	l.playerLoc = randRooms[offset]
	l.visited[randRooms[offset]] = struct{}{}

	// place the Wumpus anywhere but where the player is
	for {
		n := rand.Intn(len(l.rooms))
		if !l.rooms[n].player {
			l.rooms[n].wumpus = true
			break
		}
	}

	if l.debug {
		l.PrintDebug()
	}
}

func (l *Labyrinth) Has(id int, e Entity) bool {
	switch e {
	case Player:
		return l.rooms[id].player
	case Wumpus:
		return l.rooms[id].wumpus
	case Bat:
		return l.rooms[id].bat
	case Pit:
		return l.rooms[id].pit
	case Termite:
		return l.rooms[id].termite
	case Clue:
		return l.rooms[id].clue
	case Repel:
		return l.rooms[id].repel
	case Key:
		return l.rooms[id].key
	case Door:
		return l.rooms[id].door
	}
	return false
}

func (l *Labyrinth) Nearby(e Entity) bool {
	for _, i := range l.rooms[l.playerLoc].edges {
		if l.Has(i, e) {
			return true
		}
	}
	return false
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
		output += strconv.Itoa(l.rooms[l.rooms[room].edges[idxEdge]].fakeID)
		if !(idxEdge == len(l.rooms[room].edges)-1) {
			output += ", "
		}
	}
	return output
}

func (l *Labyrinth) PrintDebug() {
	for _, r := range l.rooms {
		r.printEntities()
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
