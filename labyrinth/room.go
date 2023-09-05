package labyrinth

import (
	"fmt"
	"math/rand"
)

type entity string

const (
	Player  entity = "player"
	Wumpus  entity = "wumpus"
	Bat     entity = "bat"
	Pit     entity = "pit"
	Termite entity = "termite"
	Clue    entity = "clue"
	Repel   entity = "repel"
	Key     entity = "key"
	Door    entity = "door"
	Rope    entity = "rope"
	Shield  entity = "shield"
)

// room is a vertex of the graph.
type room struct {
	edges    []int
	fakeID   int
	name     string
	entities map[entity]struct{}
}

func (r *room) addEntity(e entity) {
	r.entities[e] = struct{}{}
}

func (r *room) removeEntity(e entity) {
	delete(r.entities, e)
}

func (r *room) hasEntity(e entity) bool {
	_, present := r.entities[e]
	return present
}

type filterFunc func(r *room) bool

// return true if the room doesn't contain items that are needed to win the game (avoid pit on a door for example).
// used when migrating dangers
func withoutKeyItem() filterFunc {
	return func(r *room) bool {
		return !r.hasEntity(Door) && !r.hasEntity(Key)
	}
}

// return true if the room contains no consumable item.
func withoutItem() filterFunc {
	return func(r *room) bool {
		return !r.hasEntity(Rope) && !r.hasEntity(Repel) && !r.hasEntity(Shield) && !r.hasEntity(Clue)
	}
}

// return false if the room contain something dangerous (or the player).
// used at init to allow cohabitation of clues/repel/rope.
func withoutHazard() filterFunc {
	return func(r *room) bool {
		return !r.hasEntity(Pit) && !r.hasEntity(Bat) && !r.hasEntity(Player) && !r.hasEntity(Wumpus) && !r.hasEntity(Termite)
	}
}

func withEntity(e entity) filterFunc {
	return func(r *room) bool {
		return r.hasEntity(e)
	}
}

func withoutEntity(e entity) filterFunc {
	return func(r *room) bool {
		return !r.hasEntity(e)
	}
}

// randomRoom return a random room matching filters.
func (l *Labyrinth) randomRoom(filters ...filterFunc) int {
	perm := l.r.Perm(len(l.rooms))

	for _, index := range perm {
		candidate := l.rooms[index]
		match := true
		for _, f := range filters {
			if !f(&candidate) {
				match = false
				break
			}
		}
		if match {
			return index
		}
	}
	panic("no room matching, shouldn't happen")
}

func (r *room) printEntities() {
	for e := range r.entities {
		fmt.Printf("%s %d\n", e, r.fakeID)
	}
}

var (
	prefixes = []string{
		"Dra", "Bel", "Tor", "Mol", "Ven", "Aer", "Rha", "Gor",
		"Kal", "Thra", "Nar", "Hel", "For", "Lin", "Ser",
	}

	middlefixes = []string{
		"zar", "mir", "lun", "vorn", "ther", "ran", "gar",
		"kin", "val", "tel", "fyr", "mor", "sil", "din", "rex",
	}

	suffixes = []string{
		"gorn", "delve", "depth", "more", "lyn", "stone", "light",
		"shade", "rift", "vein", "mire", "fall", "peak", "dell", "spire",
	}
)

func generateCavernName() string {
	prefix := prefixes[rand.Intn(len(prefixes))]
	middle := middlefixes[rand.Intn(len(middlefixes))]
	suffix := suffixes[rand.Intn(len(suffixes))]
	return prefix + middle + suffix
}
