package dialogues

const (
	resetColor = "\033[0m"
	dim        = "\033[2m"
	red        = "\033[31m"
	yellow     = "\033[33m"
)

const (
	Start           = "Your eyes adjust to the dim light, revealing a vast cavern..."
	Room            = "You are in cavern %d.\n"
	Tunnels         = "Tunnels lead to %s.\n"
	BatsNearby      = "The distant fluttering of bats reaches your ears!"
	PitsNearby      = "A chilling draft suggests the presence of a pit nearby!"
	WumpusNearby    = "An unmistakable scent hints at a Wumpus lurking nearby!"
	ChoiceShootMove = "Move or Shoot ? (M/S): "
	WhereTo         = "Where to ? (%d/%d/%d): "
	WhereToArrow    = "➵ Aim your arrow: (%d/%d/%d): "
	NotNumber       = dim + "This is not a number." + resetColor
	NotValidDest    = dim + "That path is blocked. Choose another direction!" + resetColor
	StumbledWumpus  = "You accidentally stumbled upon the Wumpus!"
	StartledWumpus  = "Your presence spooked the Wumpus, causing it to flee!"
	KilledByWumpus  = red + "☠It ate you!☠" + resetColor
	BatTeleport     = yellow + "A giant bat took you away and dropped you into cavern %d\n!" + resetColor
	FellIntoPit     = red + "☠ You lost your footing and plummeted into a bottomless pit!! ☠" + resetColor
	MovedTo         = "You took the tunnel and arrive in cavern %d\n"
	PlayAGain       = "Would you like to venture again into the unknown? (Y/N):"
	DontUnderstand  = dim + "I'm sorry I couldn't quite catch that..." + resetColor
	FireArrow       = "You fired a curved arrow ➶"
	ArrowTravel     = "➵ The curved arrow flew through the tunnel and arrive in cavern %d.\n"
	KilledWumpus    = yellow + "➵ With a triumphant strike, your arrow fells the Wumpus! Victory is yours!" + resetColor
	ArrowStartle    = yellow + "➴ The curved arrow struck the ground, startling the Wumpus !" + resetColor
	ArrowFell       = "➴ The arrow loses its momentum, falling harmlessly to the cavern floor."
	ArrowPlayer     = red + "➴ A miscalculation! Your own arrow returns to strike you down! ☠" + resetColor
	WumpusTrample   = red + "In its panic, the Wumpus rampages through, trampling you in the process! ☠" + resetColor
	Exit            = "As you retreat from the echoing depths of the caverns, a serene silence envelops you. Thank you for venturing into the unknown with us. Until our paths cross again in the shadows... Farewell, brave adventurer."
)
