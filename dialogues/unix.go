//go:build !windows

package dialogues

const (
	reset  = "\033[0m"
	dim    = "\033[2m"
	red    = "\033[31m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	bold   = "\033[1m"
)

const (
	Start           = "Your eyes adjust to the dim light, revealing a vast cavern..."
	Room            = "You are in cavern %d.\n"
	Tunnels         = "Tunnels lead to %s.\n"
	BatsNearby      = cyan + "The distant fluttering of bats reaches your ears!" + reset
	PitsNearby      = cyan + "A chilling draft suggests the presence of a pit nearby!" + reset
	WumpusNearby    = cyan + "An unmistakable scent hints at a Wumpus lurking nearby!" + reset
	ChoiceShootMove = bold + "Move or Shoot ? ( M / S ): " + reset
	WhereTo         = bold + "Where to ? ( %s ): " + reset
	WhereToArrow    = bold + "➵ Aim your arrow: ( %s ): " + reset
	NotNumber       = dim + "This is not a number." + reset
	NotValidDest    = dim + "That path is blocked. Choose another direction!" + reset
	StumbledWumpus  = "You accidentally stumbled upon the Wumpus!"
	StartledWumpus  = "Your presence spooked the Wumpus, causing it to flee!"
	KilledByWumpus0 = red + "☠It ate you!☠" + reset
	KilledByWumpus1 = red + "☠ The Wumpus, emerging from the shadows, made you its prey! ☠" + reset
	KilledByWumpus2 = red + "☠ The lurking Wumpus caught you off-guard! Your adventure ends here ☠" + reset
	KilledByWumpus3 = red + "☠ In a fateful encounter, the Wumpus proved to be your doom! ☠" + reset
	KilledByWumpus4 = red + "☠ The Wumpus, swift and silent, ended your journey abruptly! ☠" + reset
	BatTeleport0    = yellow + "A giant bat took you away and dropped you into cavern %d!\n" + reset
	BatTeleport1    = yellow + "With a flurry of wings, a giant bat snatches you up, depositing you in a new cavern %d!\n" + reset
	BatTeleport2    = yellow + "Suddenly, talons grip you! A bat swiftly carries you aloft, dropping you in unfamiliar cavern %d.\n" + reset
	BatTeleport3    = yellow + "The shadows move, and before you realize, a bat has whisked you away, setting you down in cavern %d. Stay alert!\n" + reset
	FellIntoPit0    = red + "☠ You lost your footing and plummeted into a bottomless pit!! ☠" + reset
	FellIntoPit1    = red + "☠ Misstepping, you plummet into the abyss of a bottomless pit! ☠" + reset
	FellIntoPit2    = red + "☠ The ground betrays you, sending you falling into a dark void! ☠" + reset
	FellIntoPit3    = red + "☠ You lost your balance and met your end in the depths below! ☠" + reset
	MovedTo0        = "Venturing forth, you emerge in cavern %d\n"
	MovedTo1        = "Navigating the dimly lit tunnels, you find yourself in cavern %d.\n"
	MovedTo2        = "With cautious steps, you transition into the embrace of cavern %d.\n"
	MovedTo3        = "The winding paths lead you onward to the heart of cavern %d.\n"
	MovedTo4        = "A gentle echo announces your arrival in cavern %d.\n"
	MovedTo5        = "The path unfolds before you, revealing the secrets of cavern %d.\n"
	MovedTo6        = "Pushing forth, the mysteries of cavern %d lay before you.\n"
	MovedTo7        = "The tunnel's end opens up to the expanse of cavern %d.\n"
	MovedTo8        = "With each step, the ambiance changes, signaling your entry into cavern %d.\n"
	MovedTo9        = "Leaving the familiar behind, you're greeted by the sights and sounds of cavern %d.\n"
	MovedTo10       = "As the previous chamber fades, the allure of cavern %d beckons.\n"
	PlayAGain       = bold + "Would you like to venture again into the unknown? (Y/N):" + reset
	DontUnderstand  = dim + "I'm sorry I couldn't quite catch that..." + reset
	FireArrow       = "You fired a curved arrow ➶"
	SelectPower     = bold + "How far into the caverns will you let the arrow fly? (1-5):" + reset
	ArrowTravel     = "➵ The arrow arcs gracefully, eventually reaching cavern %d.\n"
	KilledWumpus0   = yellow + "➵ With a triumphant strike, your arrow fells the Wumpus! Victory is yours!" + reset
	KilledWumpus1   = yellow + "➵ With a resounding impact, your arrow finds its mark! The Wumpus falls defeated! Victory is yours! ☠" + reset
	KilledWumpus2   = yellow + "➵ As the echo of the arrow's flight fades, the Wumpus lets out a final cry! You've vanquished the beast! ☠" + reset
	KilledWumpus3   = yellow + "➵ Your aim is true, and the Wumpus is no more! Triumph awaits you! ☠" + reset
	KilledWumpus4   = yellow + "➵ Against all odds, you've bested the Wumpus with a single shot! Celebrate your prowess! ☠" + reset
	KilledWumpus5   = yellow + "➵ The caverns fall silent as the Wumpus meets its end by your arrow! Well done, adventurer! ☠" + reset
	ArrowStartle    = yellow + "➴ The curved arrow struck the ground, startling the Wumpus !" + reset
	ArrowFell       = "➴ The arrow loses its momentum, falling harmlessly to the cavern floor."
	ArrowPlayer     = red + "➴ A miscalculation! Your own arrow returns to strike you down! ☠" + reset
	WumpusTrample   = red + "In its panic, the Wumpus rampages through, trampling you in the process! ☠" + reset
	Turns           = dim + "You won in %d turns, firing %d arrows and visiting %d caverns.\n" + reset
	Exit0           = dim + "As you retreat from the echoing depths of the caverns, a serene silence envelops you. Thank you for venturing into the unknown with us. Until our paths cross again in the shadows... Farewell, brave adventurer." + reset
	Exit1           = dim + "Your hunt for the Wumpus concludes, for now. Till we meet in the caverns again." + reset
	Exit2           = dim + "The Wumpus remains elusive as you depart. Until another quest calls you back." + reset
	Exit3           = dim + "You choose to leave the caverns, but the Wumpus awaits your return. Farewell for now." + reset
	ExitWumpus      = dim + "The Wumpus was hidden in cavern %d\n" + reset
	NoMoreArrows    = yellow + "➴ You don't have any arrows left !" + reset
	RemainingArrows = "➵ You have %d remaining arrows.\n"
)

// advanced features.
const (
	FirstDoorDiscoveryNoKey0 = yellow + "In a dimly lit alcove, you stumble upon a door unlike any other, its very presence an enigma. Its locked nature piques your curiosity, urging you to search for a way to access its other side." + reset
	FirstDoorDiscoveryNoKey1 = yellow + "As you navigate the winding paths of the cavern, an imposing door with intricate engravings stands before you, sealing away its secrets. Without a key, you wonder how to unlock its mysteries." + reset
	FirstDoorDiscoveryNoKey2 = yellow + "Amidst the shadows, the silhouette of a grand door looms. Adorned with symbols from an age long past, it poses a silent challenge: Find the key, unveil its secrets." + reset
	FirstDoorDiscoveryNoKey3 = yellow + "The rhythmic echo of your footsteps halts abruptly as you face an ornate door, an obvious relic from ancient times. You're met with a locked barrier, urging you to find its corresponding key." + reset
	FirstDoorDiscoveryNoKey4 = yellow + "As you traverse deeper into the labyrinth, a mysterious door stands sentinel, its purpose unknown. You ponder the whereabouts of the key that might unlock its story." + reset
	BackAgainDoorNoKey0      = yellow + "Wandering through the labyrinth, you find yourself back at the mysterious door that had captivated your curiosity earlier. Its silent presence taunts you, and the absence of a key keeps its secrets just out of reach." + reset
	BackAgainDoorNoKey1      = yellow + "The familiar patterns of the grand door greet you as you navigate the winding paths. A pang of frustration hits; despite your exploration, the key remains a missing piece in this puzzle." + reset
	BackAgainDoorNoKey2      = yellow + "Once again, the cavern leads you to the ornate door you'd encountered before. It stands stoically, waiting for its key, a testament to the mysteries yet undiscovered in this maze." + reset
	BackAgainDoorNoKey3      = yellow + "The twists and turns of the cavern converge back to the imposing door you'd stumbled upon earlier. The weight of its locked state is palpable, and the elusive key continues to be a beacon of challenge." + reset
	BackAgainDoorNoKey4      = yellow + "You circle back to the enigmatic door that once held your gaze. Its secrets remain locked behind its grand facade, and without the key, it stands as a monument to the adventure yet to be unraveled." + reset
	FirstKeyDiscoveryNoDoor0 = yellow + "Tucked away in a crevice, a mysterious key gleams faintly. Intrigued, you reach out, securing the key and pondering its potential use." + reset
	FirstKeyDiscoveryNoDoor1 = yellow + "Hidden among the stones, an intricate key captures your gaze. Curiosity piqued, you carefully pick it up, wondering about its origins." + reset
	FirstKeyDiscoveryNoDoor2 = yellow + "As you traverse the cavern, you stumble upon an enigmatic key with elaborate engravings. Without hesitation, you take it, sensing its significance." + reset
	FirstKeyDiscoveryNoDoor3 = yellow + "In a dimly lit nook, a peculiar key with a lustrous sheen awaits. Compelled by its mystique, you decide to claim it for your journey ahead." + reset
	FirstKeyDiscoveryNoDoor4 = yellow + "Amongst the relics of the cavern, an ornate key stands out, beckoning to be taken. You grasp it, feeling its cool weight and unspoken promise." + reset
	FirstKeyDiscoveryNoDoor5 = yellow + "You come across an arcane key, its craftsmanship suggesting age and purpose. Drawn to its mystery, you take it, hoping it'll unlock secrets later on." + reset
	DoorThenKey0             = yellow + "You're instantly reminded of the door's enigmatic aura when you discover a key bearing similar motifs. Convinced it's more than mere coincidence, you take the key, eager to see if it unlocks the previous mystery." + reset
	DoorThenKey1             = yellow + "Tucked away in a corner, a key of intricate design beckons. The mysterious door from earlier springs to mind. Could this be its counterpart? With a mix of hope and curiosity, you decide to claim the key." + reset
	DoorThenKey2             = yellow + "As you explore, a peculiar key catches your eye, and your mind drifts to the locked door you had come across. Perhaps this key could be the gateway? With anticipation, you take it with you." + reset
	DoorThenKey3             = yellow + "The memory of that sealed door flashes back as you stumble upon a key with ancient symbols. Intrigued and optimistic, you secure the key, wondering if it might be the answer to the door's riddle." + reset
	DoorThenKey4             = yellow + "Earlier, the sight of that imposing, locked door had puzzled you. Now, as you hold a mysterious, ornate key that you've just found, hope flickers. Could this be its match? You take the key, eager to find out." + reset
	DoorThenKey5             = yellow + "Recalling the enigmatic door you encountered earlier, a sense of purpose fills you as you discover a beautifully crafted key hidden in the shadows. Without a second thought, you pocket it, hoping it might reveal the door's secrets." + reset
	KeyThenDoor0             = yellow + "The ornate key you found earlier feels heavy in your pocket as you face a magnificent door, its designs eerily similar. An air of destiny surrounds you, hinting that the key might unveil the door's concealed world." + reset
	KeyThenDoor1             = yellow + "Your fingers instinctively clutch the mysterious key you discovered earlier as you come upon an imposing door. The designs mirror each other, suggesting they're intrinsically linked. It's time to see if the key fits." + reset
	KeyThenDoor2             = yellow + "As a grand door with ancient motifs presents itself, memories of the key you picked up flood back. The connection seems unmistakable. With anticipation, you approach, ready to test your earlier find." + reset
	KeyThenDoor3             = yellow + "Before you stands a door, cloaked in mystery and echoing tales of old. The key you'd stumbled upon earlier now feels purposeful, as if it's the missing piece of this puzzle." + reset
	KeyThenDoor4             = yellow + "The sight of the enigmatic door reminds you of the key you found. With symbols and engravings that seem to tell of a shared history, you can't help but wonder if the key might unlock the door's secrets." + reset
	WumpusStillAlive0        = yellow + "However, a sense of duty holds you back. The echoing growls of the Wumpus are a reminder that your quest is not yet complete." + reset
	WumpusStillAlive1        = yellow + "Yet, as you stand at the threshold, you realize your mission remains unfinished. The Wumpus still lurks, and you cannot leave it unchecked." + reset
	WumpusStillAlive2        = yellow + "But a lingering unease prevents your exit. The Wumpus, that menacing beast, still roams the labyrinth. Your job isn't done." + reset
	WumpusStillAlive3        = yellow + "Although freedom beckons, a nagging thought anchors you. The Wumpus is still at large, and you're not one to leave a task half-done." + reset
	WumpusStillAlive4        = yellow + "But you resist the call of the outside world. Somewhere in the shadows, the Wumpus waits, and you're determined to confront it." + reset
	DoorKeyDoor0             = yellow + "Returning to the cavern that houses the mysterious door, you feel the weight of the key you've since discovered. Anticipation grows as you approach, ready to unlock the door's long-held secrets." + reset
	DoorKeyDoor1             = yellow + "The winding paths of the labyrinth lead you back to the grand door. With the newly acquired key in hand, the moment feels ripe to finally reveal what lies beyond its imposing facade." + reset
	DoorKeyDoor2             = yellow + "With the ornate key you found, you trace your steps back to the enigmatic door that once stood as an inscrutable barrier. A sense of destiny fills the air as you inch closer to unveiling its mysteries." + reset
	DoorKeyDoor3             = yellow + "The familiar designs of the door beckon as you navigate the cavern once more. The key, now in your possession, promises to shed light on the door's hidden tales." + reset
	DoorKeyDoor4             = yellow + "Back in the chamber where the intricate door first captured your wonder, the key you've discovered feels like the final piece of this puzzle. The anticipation is palpable as you prepare to turn it, bridging the gap between mystery and revelation." + reset
	ExitDoor0                = yellow + "With a deep breath, you turn the key and the door gradually gives way, leading you to the world outside the maze." + reset
	ExitDoor1                = yellow + "Holding the key, a moment of clarity strikes. With a swift motion, you unlock the door, and it swings open to display the vast expanse outside the maze." + reset
	ExitDoor2                = yellow + "The intricate design of the key glints under the dim light. Turning it, the door eases open, showing the boundless horizon beyond the labyrinth." + reset
	ExitDoor3                = yellow + "With the key's cold metal pressed against your fingers, anticipation mounts. Unlocking the door, it opens up to the freedom of the world outside." + reset
	NowExit                  = yellow + "Trapped within these winding walls, you must seek an exit." + reset
	MaybeDoor                = yellow + "Could the enigmatic door you stumbled upon be your passage out?" + reset
	MaybeKey                 = yellow + "The ornate key you discovered might hold answers to this labyrinth's riddles." + reset
	CertainKeyDoor           = yellow + "The door for which you have the key calls to you. It's time to retrace your steps and make your way to the surface through that passage!" + reset
)
