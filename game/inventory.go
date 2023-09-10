package game

import (
	. "github.com/koumbaya/wumpus/model"
)

// inventory keeps tracks of the player's inventory or actions done/memories (wumpus killed, key or door found...)
type inventory map[Entity]int

func (i *inventory) init() {
	*i = make(map[Entity]int)
}

func (i *inventory) add(e Entity) {
	(*i)[e]++
}

func (i *inventory) addn(e Entity, n int) {
	(*i)[e] += n
}

func (i *inventory) use(e Entity) {
	(*i)[e]--
}

func (i *inventory) tryUse(e Entity) bool {
	n, exist := (*i)[e]
	if !exist || n == 0 {
		return false
	}
	(*i)[e]--
	return true
}

func (i *inventory) has(e Entity) bool {
	return (*i)[e] > 0
}

func (i *inventory) count(e Entity) int {
	return (*i)[e]
}
