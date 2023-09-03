package game

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	dia "github.com/koumbaya/wumpus/dialogues"
	"github.com/koumbaya/wumpus/labyrinth"
)

type state int

const (
	waitShootMove state = iota
	waitWhereTo
	waitPlayAgain
	waitArrowWhereTo
	waitArrowPower
)

const (
	maxArrows  = 4
	randMaps   = 10 // todo : adjust
	randEvent  = 20 // original probability 1/12
	randWumpus = 20
)

type Printer interface {
	Printf(f string, a ...any)
	Print(s string)
	Println(s string)
}

type Game struct {
	l labyrinth.Labyrinth
	p Printer
	state
	turns          int
	arrowsFired    int
	timer          time.Time
	infiniteArrows bool
	wump3          bool
	// advanced features
	advanced     bool
	foundKey     bool
	foundDoor    bool
	foundRepel   bool // todo: refactor repel
	usedRepel    bool
	foundRope    bool
	usedRope     bool
	killedWumpus bool
}

func NewGame(labyrinth labyrinth.Labyrinth, printer Printer, arrows, advanced, wump3 bool) Game {
	return Game{
		l:              labyrinth,
		p:              printer,
		state:          waitShootMove,
		infiniteArrows: arrows,
		advanced:       advanced,
		wump3:          wump3,
	}
}

func (g *Game) Loop() {
	var move bool
	reader := bufio.NewReader(os.Stdin)
	g.start()

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = clean(input)

		// pre-parsing for "meta" commands & shortcut to bypass waitShootMove state and move directly.
		switch input {
		case "exit":
			g.p.Println(dia.Exit)
			g.p.Printf(dia.ExitWumpus, g.l.Wumpus())
			return
		case "reset":
			g.l.Init(1)
			g.start()
			continue
		case "debug":
			g.l.PrintDebug()
			continue
		}

		if g.state == waitShootMove { // todo : ugly way of doing things, refactor
			move, input = isMoveCommand(input)
			if move {
				g.state = waitWhereTo
			}
		}

		if g.playerState(input) {
			return
		}
		reader.Reset(reader)
	}
}

// playerState is the main state machine.
func (g *Game) playerState(input string) bool {
	switch g.state {
	case waitShootMove:
		if strings.EqualFold(input, "S") {
			if !g.infiniteArrows && g.arrowsFired >= maxArrows {
				g.p.Println(dia.NoMoreArrows)
				g.p.Print(dia.ChoiceShootMove)
				break
			}
			g.p.Print(dia.SelectPower)
			g.state = waitArrowPower
		} else if strings.EqualFold(input, "M") {
			g.whereTo()
			g.state = waitWhereTo
		} else {
			g.p.Println(dia.DontUnderstand)
			g.p.Print(dia.ChoiceShootMove)
		}
	case waitWhereTo:
		if !g.tryMove(input) {
			break // error parsing
		}
		g.turns++
		g.events()
		if g.explore() { //dead
			g.p.Print(dia.PlayAGain)
			g.state = waitPlayAgain
			break
		}
		g.describe()
		if g.keyDoor() { // won this level
			if !g.l.HasNextLevel() {
				g.p.Print(dia.PlayAGain)
				g.state = waitPlayAgain
			} else {
				g.l.Init(g.l.CurrentLevel() + 1)
				g.start()
				g.state = waitShootMove
			}
			break
		}
		g.clues()
		g.maps()
		g.p.Print(dia.ChoiceShootMove)
		g.state = waitShootMove
	case waitArrowPower:
		g.l.FireArrow(input)
		g.p.Println(dia.FireArrow)
		if !g.infiniteArrows {
			g.p.Printf(dia.RemainingArrows, maxArrows-g.arrowsFired)
		}
		g.whereToArrow()
		g.arrowsFired++
		g.state = waitArrowWhereTo
	case waitArrowWhereTo:
		if !g.tryArrow(input) {
			break
		}
		g.state = g.handleArrow()
	case waitPlayAgain:
		if strings.EqualFold(input, "Y") {
			g.l.Init(g.l.CurrentLevel())
			g.start()
			g.state = waitShootMove
		} else {
			g.p.Println(dia.Exit)
			return true
		}
	}

	return false
}

func (g *Game) start() {
	g.turns = 0
	g.arrowsFired = 0
	g.foundKey = false
	g.foundDoor = false
	g.foundRepel = false
	g.usedRepel = false
	g.foundRope = false
	g.usedRope = false
	g.killedWumpus = false
	g.timer = time.Now()
	g.p.Println(dia.Start)
	g.cavern()
	g.describe()
	g.p.Print(dia.ChoiceShootMove)
}

func (g *Game) tryArrow(input string) bool {
	d, err := strconv.Atoi(input)
	if err != nil {
		g.p.Println(dia.NotNumber)
		g.whereToArrow()
		return false
	}

	g.l.MoveArrow(d)
	return true
}

func (g *Game) handleArrow() state {
	g.p.Printf(dia.ArrowTravel, g.l.ArrowPOV())
	if g.l.Has(g.l.Arrow(), labyrinth.Wumpus) && !g.killedWumpus {
		g.p.Println(dia.KilledWumpus)
		g.killedWumpus = true
		if !g.advanced || g.keyDoor() { // check the edge case that player is already standing in the room with the door and has the key.
			g.p.Printf(dia.Turns, g.l.CurrentLevel(), g.turns, time.Since(g.timer).String(), g.arrowsFired, g.l.Visited())
			g.p.Print(dia.PlayAGain)
			return waitPlayAgain
		}

		g.mustExit()
		g.p.Print(dia.ChoiceShootMove)
		return waitShootMove
	}

	if g.l.Player() == g.l.Arrow() {
		g.p.Println(dia.ArrowPlayer)
		g.p.Printf(dia.ExitWumpus, g.l.Wumpus())
		g.p.Print(dia.PlayAGain)
		return waitPlayAgain
	}

	if g.l.PowerRemaining() == 0 {
		if g.l.StartleWumpus() && !g.killedWumpus {
			g.p.Println(dia.ArrowStartle)
			// check 1/20 odds that the wumpus moved to player's cavern
			if g.l.Has(g.l.Player(), labyrinth.Wumpus) {
				g.p.Println(dia.WumpusTrample)
				g.p.Print(dia.PlayAGain)
				return waitPlayAgain
			}
		} else {
			g.p.Println(dia.ArrowFell)
		}
		g.p.Print(dia.ChoiceShootMove)
		return waitShootMove
	}
	g.whereToArrow()
	return waitArrowWhereTo
}

func (g *Game) tryMove(input string) bool {
	d, err := strconv.Atoi(input)
	if err != nil {
		g.p.Println(dia.NotNumber)
		g.whereTo()
		return false
	}
	moved := g.l.TryMovePlayer(d)
	if !moved {
		g.p.Println(dia.NotValidDest)
		g.whereTo()
		return false
	}

	return true
}

func (g *Game) cavern() {
	g.p.Printf(dia.Room, g.l.PlayerPOV())
}

func (g *Game) describe() {
	g.p.Printf(dia.Tunnels, g.l.GetFmtNeighbors(g.l.Player()))
	if g.l.Nearby(labyrinth.Bat) {
		g.p.Println(dia.BatsNearby)
	}
	if g.l.Nearby(labyrinth.Pit) {
		g.p.Println(dia.PitsNearby)
	}
	if g.l.Nearby(labyrinth.Wumpus) && !g.killedWumpus {
		g.p.Println(dia.WumpusNearby)
	}

	if g.wump3 && g.l.Nearby(labyrinth.Termite) {
		g.p.Println(dia.TermitesNearby)
	}
}

func (g *Game) whereTo() {
	g.p.Printf(dia.WhereTo,
		g.l.GetFmtNeighbors(g.l.Player()),
	)
}

func (g *Game) whereToArrow() {
	g.p.Printf(dia.WhereToArrow,
		g.l.GetFmtNeighbors(g.l.Arrow()),
	)
}

func (g *Game) events() {
	if !g.wump3 {
		return
	}

	if rand.Intn(randEvent) == 0 {
		g.l.Earthquake()
		g.p.Println(dia.Earthquake)
	}

	if rand.Intn(randEvent) == 0 {
		g.l.BatMigration()
		g.p.Println(dia.BatMigration)
	}

	if rand.Intn(randEvent) == 0 {
		g.l.TermitesMigration()
		g.p.Println(dia.TermiteMigration)
	}

	if rand.Intn(randWumpus) == 0 { // lower probability
		g.l.SleepwalkWumpus()
		g.p.Println(dia.SleepWalkWumpus)
	}
}

func (g *Game) explore() bool {
	g.p.Printf(dia.MovedTo, g.l.PlayerPOV())
	return g.hazards()
}

func (g *Game) clues() {
	if !g.advanced {
		return
	}

	if g.l.Has(g.l.Player(), labyrinth.Clue) {
		loc, subject := g.l.GetClue(g.l.Player())
		g.p.Printf(dia.FoundClue, subject, loc)
	}

	if g.l.FoundRepel() {
		g.foundRepel = true
		g.p.Println(dia.FoundRepel)
	}

	if g.l.FoundRope() {
		g.foundRope = true
		g.p.Println(dia.FoundRope)
	}
}

// maps randomly gives partial maps tips.
func (g *Game) maps() {
	if !g.advanced {
		return
	}

	if rand.Intn(randMaps) == 0 {
		g.p.Printf(dia.PartialMap, g.l.GetFmtMap())
	}
}

// hazards checks for wumpus/bats/pits when entering a new room.
// Return true if a hazard killed the player.
// If a bat moves the player, call recursively.
func (g *Game) hazards() bool {
	// the wumpus is immune to hazards, so we check for it first
	if g.l.Has(g.l.Player(), labyrinth.Wumpus) && !g.killedWumpus {
		g.p.Println(dia.StumbledWumpus)
		if dead := g.l.FoundWumpus(); dead {
			g.p.Println(dia.KilledByWumpus)
			return true
		}
		g.p.Println(dia.StartledWumpus)
	}

	// the bat may teleport to a pit or the wumpus, so we check it second
	if g.l.Has(g.l.Player(), labyrinth.Bat) {
		if g.foundRepel && !g.usedRepel {
			g.usedRepel = true
			g.p.Println(dia.UseRepel)
		} else {
			g.p.Printf(dia.BatTeleport, g.l.ActivateBat())
			return g.hazards()
		}
	}

	if g.l.Has(g.l.Player(), labyrinth.Pit) {
		if g.foundRope && !g.usedRope {
			g.usedRope = true
			g.p.Println(dia.UseRope)
		} else {
			g.p.Println(dia.FellIntoPit)
			g.p.Printf(dia.ExitWumpus, g.l.Wumpus())
			return true
		}
	}

	if g.wump3 && g.arrowsFired < maxArrows && g.l.Has(g.l.Player(), labyrinth.Termite) {
		g.p.Println(dia.TermiteEatArrow)
		g.p.Printf(dia.RemainingArrows, maxArrows-g.arrowsFired)
		g.arrowsFired++
	}

	return false
}

// keyDoor resolve dialogues & handle logic for key & door depending on the order of discovery.
// return true if all winning conditions are met
func (g *Game) keyDoor() bool {
	if !g.advanced {
		return false
	}

	door := g.l.Has(g.l.Player(), labyrinth.Door)
	key := g.l.Has(g.l.Player(), labyrinth.Key)

	if !door && !key {
		return false
	}

	canUnlock := false // in the door room with the key.
	switch {
	case door && g.foundKey && g.foundDoor:
		// found the door, then the key, and are back to the room with the door
		g.p.Println(dia.DoorKeyDoor)
		canUnlock = true
	case door && g.foundKey:
		// found the key first then the door (first time seeing it)
		g.p.Println(dia.KeyThenDoor)
		g.foundDoor = true
		canUnlock = true
	case door && !g.foundDoor:
		// first time seeing the door, no key
		g.p.Println(dia.FirstDoorDiscoveryNoKey)
		g.foundDoor = true
	case door:
		// back in the cavern with the door again
		g.p.Println(dia.BackAgainDoorNoKey)
	case key && g.foundDoor && !g.foundKey:
		// found the door first, then this key
		g.p.Println(dia.DoorThenKey)
		g.foundKey = true
	case key && !g.foundDoor && !g.foundKey:
		// found the key first
		g.p.Println(dia.FirstKeyDiscoveryNoDoor)
		g.foundKey = true
	}

	if canUnlock && !g.killedWumpus {
		g.p.Println(dia.WumpusStillAlive)
	} else if canUnlock {
		if !g.l.HasNextLevel() {
			g.p.Println(dia.ExitDoor)
		} else {
			g.p.Println(dia.GoNextLevel)
		}
		g.p.Printf(dia.Turns, g.l.CurrentLevel(), g.turns, time.Since(g.timer).String(), g.arrowsFired, g.l.Visited())
		return true
	}

	return false
}

// mustExit resolve the dialogues for the next step of the game in advanced mode.
func (g *Game) mustExit() {
	switch {
	case g.foundKey && g.foundDoor:
		g.p.Println(dia.CertainKeyDoor)
	case g.foundKey && !g.foundDoor:
		g.p.Println(dia.MaybeKey)
	case !g.foundKey && g.foundDoor:
		g.p.Println(dia.MaybeDoor)
	default:
		g.p.Println(dia.NowExit)
	}
}

func clean(input string) string {
	input = strings.TrimRight(input, "\n")
	input = strings.TrimRight(input, "\r\n")
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToLower(input)
	var stack []rune
	for _, r := range input {
		if r == '\b' || r == '\ufffd' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, r)
		}
	}

	return string(stack)
}

func isMoveCommand(input string) (bool, string) {
	output := strings.ReplaceAll(input, "m", "")
	_, err := strconv.Atoi(output)
	if err != nil {
		return false, input
	}
	return true, output

}
