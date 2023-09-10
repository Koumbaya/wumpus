package labyrinth

import (
	. "github.com/koumbaya/wumpus/model"
)

// Player return the player location.
func (l *Labyrinth) Player() int {
	return l.playerLoc
}

// PlayerPOV return the shuffled player location.
func (l *Labyrinth) PlayerPOV() int {
	return l.rooms[l.playerLoc].fakeID
}

// TryMovePlayer moves the player if the position is valid.
func (l *Labyrinth) TryMovePlayer(fakeTarget int) bool {
	target := l.fakeIDs[fakeTarget]
	if l.validDestination(l.playerLoc, target) || l.debug /*allow teleport in debug mode*/ {
		l.rooms[l.playerLoc].removeEntity(Player)
		l.rooms[target].addEntity(Player)
		l.playerLoc = target
		l.visited[target] = struct{}{}
		return true
	}

	return false
}

func (l *Labyrinth) Visited() int {
	return len(l.visited)
}
