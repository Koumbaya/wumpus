// Package labyrinth handles the labyrinth and the movements of its inhabitants.
package labyrinth

import (
	"fmt"
	"math/rand"
	"strconv"

	. "github.com/koumbaya/wumpus/model"
)

const (
	nbWumpus   = 1
	nbBats     = 2
	nbPits     = 2
	nbTermites = 2
	nbClues    = 3
	nbRepel    = 1
	nbKey      = 1
	nbDoor     = 1
	nbRope     = 1
	nbShield   = 1
)

// Labyrinth is the collection of cLevel making up the dodecahedron.
type Labyrinth struct {
	r *rand.Rand
	// levels store all the jsons levels
	levels   map[int]level
	curLevel int
	rooms    []room // current level topology // todo : do we need to copy this anyway ?
	// visited keep track of the # of explored rooms.
	visited  map[int]struct{}
	advanced bool // experimental
	wump3    bool
	debug    bool

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
	l.cluesGiven = make(map[string]struct{}, l.levels[l.curLevel].setup.nbClue)
	randRooms := l.r.Perm(len(l.rooms))

	// use the randomization to give arbitrary numbers to the caves so that each play through is unique.
	names := getCavernNames(len(l.rooms))
	for i := 0; i < len(l.rooms); i++ {
		l.rooms[i].fakeID = randRooms[i] + 1 // +1 so the player never sees a room 0.
		l.rooms[i].name = names[i]
		// reset entities on restart
		l.rooms[i].entities = make(map[Entity]struct{}, 5) // 10 different entities but few can coexist anyway.
		// keep track of fakeIds for fast access
		l.fakeIDs[l.rooms[i].fakeID] = i
	}

	// place pits & bats in distinct locations
	for i := 0; i < l.levels[l.curLevel].setup.nbPit; i++ {
		if len(l.levels[l.curLevel].setup.pitsPos) == 0 {
			l.rooms[l.randomRoom(withoutHazard())].addEntity(Pit)
		} else {
			l.rooms[rollIndex(l.levels[l.curLevel].setup.pitsPos, i)].addEntity(Pit)
		}
	}

	for i := 0; i < l.levels[l.curLevel].setup.nbBat; i++ {
		if len(l.levels[l.curLevel].setup.batsPos) == 0 {
			l.rooms[l.randomRoom(withoutHazard())].addEntity(Bat)
		} else {
			l.rooms[rollIndex(l.levels[l.curLevel].setup.batsPos, i)].addEntity(Bat)
		}
	}

	if l.wump3 {
		for i := 0; i < l.levels[l.curLevel].setup.nbTermite; i++ {
			if len(l.levels[l.curLevel].setup.termitePos) == 0 {
				l.rooms[l.randomRoom(withoutHazard())].addEntity(Termite)
			} else {
				l.rooms[rollIndex(l.levels[l.curLevel].setup.termitePos, i)].addEntity(Termite)
			}
		}
	}

	if l.advanced {
		for i := 0; i < nbKey; i++ {
			if l.levels[l.curLevel].setup.keyPos == nil {
				l.rooms[l.randomRoom(withoutHazard(), withoutKeyItem())].addEntity(Key)
			} else {
				l.rooms[*l.levels[l.curLevel].setup.keyPos].addEntity(Key)
			}
		}
		for i := 0; i < nbDoor; i++ {
			if l.levels[l.curLevel].setup.doorPos == nil {
				l.rooms[l.randomRoom(withoutHazard(), withoutKeyItem())].addEntity(Door)
			} else {
				l.rooms[*l.levels[l.curLevel].setup.doorPos].addEntity(Door)

			}
		}

		// clues/rope/repel/shield can be in the same room as each other, but we avoid clues on door/key
		for i := 0; i < l.levels[l.curLevel].setup.nbClue; i++ {
			if len(l.levels[l.curLevel].setup.cluePos) == 0 {
				l.rooms[l.randomRoom(withoutHazard(), withoutKeyItem(), withoutEntity(Clue))].addEntity(Clue)
			} else {
				l.rooms[rollIndex(l.levels[l.curLevel].setup.cluePos, i)].addEntity(Clue)
			}
		}

		for i := 0; i < l.levels[l.curLevel].setup.nbRepel; i++ {
			if len(l.levels[l.curLevel].setup.repelPos) == 0 {
				l.rooms[l.randomRoom(withoutHazard(), withoutEntity(Repel))].addEntity(Repel)
			} else {
				l.rooms[rollIndex(l.levels[l.curLevel].setup.repelPos, i)].addEntity(Repel)
			}
		}

		for i := 0; i < l.levels[l.curLevel].setup.nbRope; i++ {
			if len(l.levels[l.curLevel].setup.ropePos) == 0 {
				l.rooms[l.randomRoom(withoutHazard(), withoutEntity(Rope))].addEntity(Rope)
			} else {
				l.rooms[rollIndex(l.levels[l.curLevel].setup.ropePos, i)].addEntity(Rope)
			}
		}

		for i := 0; i < l.levels[l.curLevel].setup.nbShield; i++ {
			if len(l.levels[l.curLevel].setup.shieldPos) == 0 {
				l.rooms[l.randomRoom(withoutHazard(), withoutEntity(Shield))].addEntity(Shield)
			} else {
				l.rooms[rollIndex(l.levels[l.curLevel].setup.shieldPos, i)].addEntity(Shield)
			}
		}
	}

	// place player
	if l.levels[l.curLevel].setup.playerStartPos == nil {
		l.playerLoc = l.randomRoom(withoutHazard(), withoutKeyItem(), withoutItem())
	} else {
		l.playerLoc = *l.levels[l.curLevel].setup.playerStartPos
	}
	l.rooms[l.playerLoc].addEntity(Player)
	l.visited[l.playerLoc] = struct{}{}

	// place the Wumpus anywhere but where the player is
	if l.levels[l.curLevel].setup.wumpusStartPos == nil {
		l.rooms[l.randomRoom(withoutEntity(Player))].addEntity(Wumpus)
	} else {
		l.rooms[*l.levels[l.curLevel].setup.wumpusStartPos].addEntity(Wumpus)
	}

	if l.debug {
		l.PrintDebug()
	}
}

func (l *Labyrinth) Name(id int) string {
	return l.rooms[id].name
}

func (l *Labyrinth) Has(id int, e Entity) bool {
	return l.rooms[id].hasEntity(e)
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
	for i, r := range l.rooms {
		fmt.Printf("cave %d (%d): ", r.fakeID, i)
		r.printEntities()
		fmt.Println()
		if r.hasEntity(Wumpus) {
			fmt.Printf("wumpus neighboring caves: %s\n", l.GetFmtNeighbors(i))
		}
	}
}

// rollIndex allow us to never hit out of bounds on user-defined positions
func rollIndex[T any](s []T, i int) T {
	return s[i%len(s)]
}
