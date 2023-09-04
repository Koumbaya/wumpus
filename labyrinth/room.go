package labyrinth

import (
	"fmt"
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
	Rope
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
	rope    bool
}

// room is a vertex of the graph.
type room struct {
	edges  []int
	fakeID int
	entities
}

type filterFunc func(room) bool

// return false if the room already has one entity that can't cohabit with others or is needed to win the game (ex. no pit on a door).
// technically player can cohabit, but we don't want to add another entity to the player's location (usually).
// used when migrating dangers
func withoutKeyItem() filterFunc {
	return func(r room) bool {
		return !r.door && !r.key
	}
}

// return false if the room contains a consumable item.
func withoutItem() filterFunc {
	return func(r room) bool {
		return !r.rope && !r.repel && !r.clue
	}
}

// return false if the room contain something dangerous (or the player).
// used at init to allow cohabitation of clues/repel/rope.
func withoutHazard() filterFunc {
	return func(r room) bool {
		return !r.pit && !r.bat && !r.player && !r.wumpus && !r.termite
	}
}

func withEntity(e Entity) filterFunc {
	return func(r room) bool {
		switch e {
		case Player:
			return r.player
		case Wumpus:
			return r.wumpus
		case Bat:
			return r.bat
		case Pit:
			return r.pit
		case Termite:
			return r.termite
		case Clue:
			return r.clue
		case Repel:
			return r.repel
		case Key:
			return r.key
		case Door:
			return r.door
		case Rope:
			return r.rope
		}
		return true
	}
}

func withoutEntity(e Entity) filterFunc {
	return func(r room) bool {
		switch e {
		case Player:
			return !r.player
		case Wumpus:
			return !r.wumpus
		case Bat:
			return !r.bat
		case Pit:
			return !r.pit
		case Termite:
			return !r.termite
		case Clue:
			return !r.clue
		case Repel:
			return !r.repel
		case Key:
			return !r.key
		case Door:
			return !r.door
		case Rope:
			return !r.rope
		}
		return false
	}
}

// randomRoom return a random room matching filters.
func (l *Labyrinth) randomRoom(filters ...filterFunc) int {
	perm := l.r.Perm(len(l.rooms))

	for _, index := range perm {
		candidate := l.rooms[index]
		match := true
		for _, f := range filters {
			if !f(candidate) {
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
	if r.rope {
		fmt.Printf("rope %d\n", r.fakeID)
	}
}
