package game

type item int

const (
	repel item = iota
	key
	door
	rope
	shield
	arrow
	wumpusHide
)

// inventory keeps tracks of the player's inventory or actions done/memories (wumpus killed, key or door found...)
type inventory map[item]int

func (i *inventory) init() {
	*i = make(map[item]int)
}

func (i *inventory) add(it item) {
	(*i)[it]++
}

func (i *inventory) addn(it item, n int) {
	(*i)[it] += n
}

func (i *inventory) use(it item) {
	(*i)[it]--
}

func (i *inventory) tryUse(it item) bool {
	n, exist := (*i)[it]
	if !exist || n == 0 {
		return false
	}
	(*i)[it]--
	return true
}

func (i *inventory) has(it item) bool {
	return (*i)[it] > 0
}

func (i *inventory) count(it item) int {
	return (*i)[it]
}
