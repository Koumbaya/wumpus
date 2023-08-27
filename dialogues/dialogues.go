package dialogues

const (
	resetColor = "\033[0m"
	dim        = "\033[2m"
	red        = "\033[31m"
	yellow     = "\033[33m"
)

const (
	Start           = "You wake up in a giant cavern..."
	Room            = "You are in cavern %d.\n"
	Tunnels         = "Tunnels lead to %s.\n"
	BatsNearby      = "You hear bats nearby!"
	PitsNearby      = "You feel a strong draft!"
	WumpusNearby    = "You smell a Wumpus!"
	ChoiceShootMove = "Move or Shoot ? (M/S): "
	WhereTo         = "Where to ? (%d/%d/%d): "
	WhereToArrow    = "➵ Where to ? (%d/%d/%d): "
	NotNumber       = dim + "This is not a number." + resetColor
	NotValidDest    = dim + "You can't go there from here!" + resetColor
	StumbledWumpus  = "You stumbled upon the Wumpus!"
	StartledWumpus  = "You startled it, it ran away!"
	KilledByWumpus  = red + "☠It ate you!☠" + resetColor
	BatTeleport     = yellow + "A giant bat took you away and dropped you into cavern %d\n!" + resetColor
	FellIntoPit     = red + "☠You fell into a bottomless pit!☠" + resetColor
	MovedTo         = "You took the tunnel and arrive in cavern %d\n"
	PlayAGain       = "Do you want to play again ? (Y/N)"
	DontUnderstand  = dim + "I'm sorry I couldn't quite catch that..." + resetColor
	FireArrow       = "You fired a curved arrow ➶"
	ArrowTravel     = "➵ The curved arrow flew through the tunnel and arrive in cavern %d.\n"
	KilledWumpus    = yellow + "➵ The curved arrow struck the Wumpus! It died ☠! You won!" + resetColor
	ArrowStartle    = yellow + "➴ The curved arrow struck the ground, startling the Wumpus !" + resetColor
	ArrowFell       = "➴ The curved arrow fell to the ground silently."
	ArrowPlayer     = red + "➴ The curved arrow struck you ! You died ☠!" + resetColor
	WumpusTrample   = red + "The Wumpus trampled you in its escape! You died ☠!" + resetColor
)
