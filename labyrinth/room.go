package labyrinth

import (
	"fmt"
	"math/rand"

	. "github.com/koumbaya/wumpus/model"
)

// room is a vertex of the graph.
type room struct {
	edges    []int
	fakeID   int
	name     string
	entities map[Entity]struct{}
}

func (r *room) addEntity(e Entity) {
	r.entities[e] = struct{}{}
}

func (r *room) removeEntity(e Entity) {
	delete(r.entities, e)
}

func (r *room) hasEntity(e Entity) bool {
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

func withEntity(e Entity) filterFunc {
	return func(r *room) bool {
		return r.hasEntity(e)
	}
}

func withoutEntity(e Entity) filterFunc {
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
	return -1
}

func (r *room) printEntities() {
	for e := range r.entities {
		fmt.Printf("%s ", e)
	}
}

func getCavernNames(n int) []string {
	filter := make(map[string]struct{}, n)
	out := make([]string, 0, n)
	for {
		if len(out) == n {
			return out
		}
		name := generateCavernName()
		if _, exist := filter[name]; !exist {
			filter[name] = struct{}{}
			out = append(out, name)
		}
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
