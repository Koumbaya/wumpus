// Package labyrinth handles the labyrinth and the movements of its inhabitants.
package labyrinth

import (
	"math/rand"
	"strconv"
)

const (
	nbWumpus   = 1
	nbBats     = 2
	nbPits     = 2
	nbTermites = 1
	nbClues    = 3
	nbRepel    = 1
	nbKey      = 1
	nbDoor     = 1
	nbRope     = 1
)

type level struct {
	rooms  []room
	number int
	name   string
}

// Labyrinth is the collection of cLevel making up the dodecahedron.
type Labyrinth struct {
	r *rand.Rand
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
func NewLabyrinth(advanced, debug, wump3 bool, level int, seed int64) Labyrinth {
	l := Labyrinth{
		// there is probably a way to do this mathematically but is it worth it ?
		levels:   loadLevels(),
		advanced: advanced,
		debug:    debug,
		wump3:    wump3,
		r:        rand.New(rand.NewSource(seed)),
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
	randRooms := l.r.Perm(len(l.rooms))

	// use the randomization to give arbitrary numbers to the caves so that each play through is unique.
	for i := 0; i < len(l.rooms); i++ {
		l.rooms[i].fakeID = randRooms[i] + 1 // +1 so the player never sees a room 0.
		// reset entities on restart
		l.rooms[i].entities = entities{}
		// keep track of fakeIds for fast access
		l.fakeIDs[l.rooms[i].fakeID] = i
	}

	// place pits & bats in distinct locations
	for i := 0; i < nbPits; i++ {
		l.rooms[l.randomRoom(withoutHazard())].pit = true
	}

	for i := 0; i < nbBats; i++ {
		l.rooms[l.randomRoom(withoutHazard())].bat = true
	}

	if l.wump3 {
		for i := 0; i < nbTermites; i++ {
			l.rooms[l.randomRoom(withoutHazard())].termite = true
		}
	}

	if l.advanced {
		for i := 0; i < nbKey; i++ {
			l.rooms[l.randomRoom(withoutHazard(), withoutKeyItem())].key = true
		}
		for i := 0; i < nbDoor; i++ {
			l.rooms[l.randomRoom(withoutHazard(), withoutKeyItem())].door = true
		}

		// clues/rope/repel can be in the same room
		for i := 0; i < nbRepel; i++ {
			l.rooms[l.randomRoom(withoutHazard(), withoutKeyItem(), withoutEntity(Repel))].repel = true
		}

		for i := 0; i < nbRope; i++ {
			l.rooms[l.randomRoom(withoutHazard(), withoutKeyItem(), withoutEntity(Rope))].rope = true
		}

		for i := 0; i < nbClues; i++ {
			l.rooms[l.randomRoom(withoutHazard(), withoutKeyItem(), withoutEntity(Clue))].clue = true
		}
	}

	// place player
	l.playerLoc = l.randomRoom(withoutHazard(), withoutKeyItem(), withoutItem())
	l.rooms[l.playerLoc].player = true
	l.visited[l.playerLoc] = struct{}{}

	// place the Wumpus anywhere but where the player is
	l.rooms[l.randomRoom(withoutEntity(Player))].wumpus = true

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
	case Rope:
		return l.rooms[id].rope
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
	return l.rooms[location].edges[l.r.Intn(len(l.rooms[location].edges))]
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
