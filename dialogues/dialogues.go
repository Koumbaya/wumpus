package dialogues

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	reset  = "\033[0m"
	dim    = "\033[2m"
	red    = "\033[31m"
	yellow = "\033[33m"
	bold   = "\033[1m"
)

const (
	Start           = "Your eyes adjust to the dim light, revealing a vast cavern..."
	Room            = "You are in cavern %d.\n"
	Tunnels         = "Tunnels lead to %s.\n"
	BatsNearby      = "The distant fluttering of bats reaches your ears!"
	PitsNearby      = "A chilling draft suggests the presence of a pit nearby!"
	WumpusNearby    = "An unmistakable scent hints at a Wumpus lurking nearby!"
	ChoiceShootMove = bold + "Move or Shoot ? ( M / S ): " + reset
	WhereTo         = bold + "Where to ? ( %d / %d / %d ): " + reset
	WhereToArrow    = bold + "➵ Aim your arrow: ( %d / %d / %d ): " + reset
	NotNumber       = dim + "This is not a number." + reset
	NotValidDest    = dim + "That path is blocked. Choose another direction!" + reset
	StumbledWumpus  = "You accidentally stumbled upon the Wumpus!"
	StartledWumpus  = "Your presence spooked the Wumpus, causing it to flee!"
	KilledByWumpus0 = red + "☠It ate you!☠" + reset
	KilledByWumpus1 = red + "☠ The Wumpus, emerging from the shadows, made you its prey! ☠" + reset
	KilledByWumpus2 = red + "☠ The lurking Wumpus caught you off-guard! Your adventure ends here ☠" + reset
	KilledByWumpus3 = red + "☠ In a fateful encounter, the Wumpus proved to be your doom! ☠" + reset
	KilledByWumpus4 = red + "☠ The Wumpus, swift and silent, ended your journey abruptly! ☠" + reset
	BatTeleport     = yellow + "A giant bat took you away and dropped you into cavern %d\n!" + reset
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
	Exit            = dim + "As you retreat from the echoing depths of the caverns, a serene silence envelops you. Thank you for venturing into the unknown with us. Until our paths cross again in the shadows... Farewell, brave adventurer." + reset
	NoMoreArrows    = yellow + "➴ You don't have any arrows left !" + reset
	RemainingArrows = "➵ You have %d remaining arrows.\n"
)

type Printer struct {
	delay time.Duration
}

func NewPrinter(t time.Duration) *Printer {
	return &Printer{delay: t}
}

func (p *Printer) Printf(f string, a ...any) {
	if p.delay == 0 {
		fmt.Printf(f, a...)
		return
	}

	r := fmt.Sprintf(f, a...)
	for _, c := range r {
		time.Sleep(p.delay)
		fmt.Print(string(c))
	}
}

func (p *Printer) Print(s string) {
	if p.delay == 0 {
		fmt.Print(s)
		return
	}
	for _, c := range s {
		time.Sleep(p.delay)
		fmt.Print(string(c))
	}
}

func (p *Printer) Println(s string) {
	if p.delay == 0 {
		fmt.Println(s)
		return
	}
	for _, c := range s {
		time.Sleep(p.delay)
		fmt.Print(string(c))
	}
	fmt.Println()
}

func KilledWumpus() string {
	r := []string{
		KilledWumpus0,
		KilledWumpus1,
		KilledWumpus2,
		KilledWumpus3,
		KilledWumpus4,
		KilledWumpus5,
	}

	return r[rand.Intn(len(r))]
}

func FellIntoPit() string {
	r := []string{
		FellIntoPit0,
		FellIntoPit1,
		FellIntoPit2,
		FellIntoPit3,
	}

	return r[rand.Intn(len(r))]
}

func KilledByWumpus() string {
	r := []string{
		KilledByWumpus0,
		KilledByWumpus1,
		KilledByWumpus2,
		KilledByWumpus3,
		KilledByWumpus4,
	}

	return r[rand.Intn(len(r))]
}

func MovedTo() string {
	r := []string{
		MovedTo0,
		MovedTo1,
		MovedTo2,
		MovedTo3,
		MovedTo4,
		MovedTo5,
		MovedTo6,
		MovedTo7,
		MovedTo8,
		MovedTo9,
		MovedTo10,
	}
	return r[rand.Intn(len(r))]
}
