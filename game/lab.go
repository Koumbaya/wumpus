package game

import (
	. "github.com/koumbaya/wumpus/model"
)

type Labyrinth interface {
	GetClue(clueLoc int) (loc int, subject Entity)
	GetFmtMap() (output string)
	FoundObject(e Entity) bool
	Player() int
	PlayerPOV() int
	TryMovePlayer(fakeTarget int) bool
	Visited() int
	Arrow() int
	ArrowPOV() int
	FireArrow()
	MoveArrow(fakeTarget int)
	Wumpus() int
	FoundWumpus(awake bool) (killed bool)
	StartleWumpus()
	MigrateWumpus()
	ActivateBat() int
	Migration(e Entity)
	CurrentLevel() int
	HasNextLevel() bool
	Init(targetLvl int)
	Name(id int) string
	Has(id int, e Entity) bool
	Nearby(e Entity) bool
	GetFmtNeighbors(room int) string
	PrintDebug()
}
