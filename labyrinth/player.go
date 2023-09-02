package labyrinth

// Player return the player location.
func (l *Labyrinth) Player() int {
	return l.player
}

// PlayerPOV return the shuffled player location.
func (l *Labyrinth) PlayerPOV() int {
	return l.shuffled[l.player] + 1
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
