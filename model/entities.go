package model

// Entity is the common type used by game and labyrinth to refer to the same objects.
type Entity string

const (
	Player  Entity = "player"
	Wumpus  Entity = "wumpus"
	Bat     Entity = "bat"
	Pit     Entity = "pit"
	Termite Entity = "termite"
	Clue    Entity = "clue"
	Repel   Entity = "repel"
	Key     Entity = "key"
	Door    Entity = "door"
	Rope    Entity = "rope"
	Shield  Entity = "shield"
	Arrow   Entity = "arrow"
)
