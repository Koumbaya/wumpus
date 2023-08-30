package dialogues

const (
	reset  = "\033[0m"
	dim    = "\033[2m"
	red    = "\033[31m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	bold   = "\033[1m"
)

func mapColors(s string) string {
	switch s {
	case "reset":
		return reset
	case "dim":
		return dim
	case "red":
		return red
	case "yellow":
		return yellow
	case "cyan":
		return cyan
	case "bold":
		return bold
	}
	return ""
}

const (
	Start           = "Start"
	Room            = "Room"
	Tunnels         = "Tunnels"
	BatsNearby      = "BatsNearby"
	PitsNearby      = "PitsNearby"
	WumpusNearby    = "WumpusNearby"
	ChoiceShootMove = "ChoiceShootMove"
	WhereTo         = "WhereTo"
	WhereToArrow    = "WhereToArrow"
	NotNumber       = "NotNumber"
	NotValidDest    = "NotValidDest"
	StumbledWumpus  = "StumbledWumpus"
	StartledWumpus  = "StartledWumpus"
	KilledByWumpus  = "KilledByWumpus"
	BatTeleport     = "BatTeleport"
	FellIntoPit     = "FellIntoPit"
	MovedTo         = "MovedTo"
	PlayAGain       = "PlayAGain"
	DontUnderstand  = "DontUnderstand"
	FireArrow       = "FireArrow"
	SelectPower     = "SelectPower"
	ArrowTravel     = "ArrowTravel"
	KilledWumpus    = "KilledWumpus"
	ArrowStartle    = "ArrowStartle"
	ArrowFell       = "ArrowFell"
	ArrowPlayer     = "ArrowPlayer"
	WumpusTrample   = "WumpusTrample"
	Turns           = "Turns"
	Exit            = "Exit"
	ExitWumpus      = "ExitWumpus"
	NoMoreArrows    = "NoMoreArrows"
	RemainingArrows = "RemainingArrows"
)

// advanced features.
const (
	FirstDoorDiscoveryNoKey = "FirstDoorDiscoveryNoKey"
	BackAgainDoorNoKey      = "BackAgainDoorNoKey"
	FirstKeyDiscoveryNoDoor = "FirstKeyDiscoveryNoDoor"
	DoorThenKey             = "DoorThenKey"
	KeyThenDoor             = "KeyThenDoor"
	WumpusStillAlive        = "WumpusStillAlive"
	DoorKeyDoor             = "DoorKeyDoor"
	ExitDoor                = "ExitDoor"
	NowExit                 = "NowExit"
	MaybeDoor               = "MaybeDoor"
	MaybeKey                = "MaybeKey"
	CertainKeyDoor          = "CertainKeyDoor"
)
