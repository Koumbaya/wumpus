package game

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	dia "github.com/koumbaya/wumpus/dialogues"
	. "github.com/koumbaya/wumpus/model"
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
	nbArrows   = 5
	randMaps   = 12 // todo : adjust
	randArrows = 16 // todo : adjust
	randEvent  = 20 // original probability 1/12
	randWumpus = 20
)

type Printer interface {
	Printf(f string, a ...any)
	Print(s string)
	Println(s string)
	Get(s string) string
}

type Cfg struct {
	Seed           int64
	InfiniteArrows bool
	Wump3          bool
	Advanced       bool
}

type Game struct {
	l Labyrinth
	p Printer
	state
	cfg         Cfg
	r           *rand.Rand
	turns       int
	arrowsFired int
	wumpusAwake bool
	timer       time.Time
	inventory   inventory
	// keep track of the arrow remaining distance
	arrowTravel int
}

func NewGame(labyrinth Labyrinth, printer Printer, cfg Cfg) Game {
	return Game{
		l:     labyrinth,
		p:     printer,
		state: waitShootMove,
		cfg:   cfg,
		r:     rand.New(rand.NewSource(cfg.Seed)),
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
		case "seed":
			g.p.Printf(dia.Seed, g.cfg.Seed)
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
		g.state = g.handleChoice(input)
	case waitWhereTo:
		g.state = g.handleMovement(input)
	case waitArrowPower:
		g.l.FireArrow()
		g.arrowTravel = parsePower(input)
		g.p.Println(dia.FireArrow)
		if !g.cfg.InfiniteArrows {
			g.inventory.use(Arrow)
			g.p.Printf(dia.RemainingArrows, g.inventory.count(Arrow))
		}
		g.arrowsFired++
		g.whereToArrow()
		g.state = waitArrowWhereTo
	case waitArrowWhereTo:
		g.state = g.handleArrow(input)
	case waitPlayAgain:
		if strings.EqualFold(input, "Y") {
			g.l.Init(g.l.CurrentLevel())
			g.start()
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
	g.wumpusAwake = false
	g.inventory.init()
	g.inventory.addn(Arrow, nbArrows)
	g.timer = time.Now()
	g.p.Println(dia.Start)
	g.cavern()
	g.describe()
	g.p.Print(dia.ChoiceShootMove)
	g.state = waitShootMove
}

func (g *Game) handleChoice(input string) state {
	if strings.EqualFold(input, "S") {
		if !g.cfg.InfiniteArrows && !g.inventory.has(Arrow) {
			g.p.Println(dia.NoMoreArrows)
			g.p.Print(dia.ChoiceShootMove)
			return waitShootMove
		}
		g.p.Print(dia.SelectPower)
		return waitArrowPower
	}

	if strings.EqualFold(input, "M") {
		g.whereTo()
		return waitWhereTo
	}

	g.p.Println(dia.DontUnderstand)
	g.p.Print(dia.ChoiceShootMove)

	return waitShootMove
}

func (g *Game) handleMovement(input string) state {
	if !g.tryMove(input) {
		// invalid input
		return waitWhereTo
	}

	g.turns++
	g.events() // check random events.

	if g.explore() { //dead
		g.p.Print(dia.PlayAGain)
		return waitPlayAgain
	}

	g.describe()

	if g.keyDoor() { // won this level
		if !g.l.HasNextLevel() {
			g.p.Print(dia.PlayAGain)
			return waitPlayAgain
		}
		// load next level
		g.l.Init(g.l.CurrentLevel() + 1)
		g.start()
		return waitShootMove
	}

	g.items()
	g.p.Print(dia.ChoiceShootMove)
	return waitShootMove
}

func (g *Game) handleArrow(input string) state {
	d, _ := strconv.Atoi(input) // we don't care for invalid input, it'll be moved at random.
	g.l.MoveArrow(d)

	g.arrowTravel--
	g.p.Printf(dia.ArrowTravel, g.l.ArrowPOV())

	if g.l.Has(g.l.Arrow(), Wumpus) && !g.inventory.has(Wumpus) {
		g.p.Println(dia.KilledWumpus)
		g.inventory.add(Wumpus)
		if !g.cfg.Advanced || g.keyDoor() { // check the edge case that player is already standing in the room with the door and has the key.
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

	// once we checked that no wumpus or player is in the room, we check for termites (to avoid protection from arrows)
	if g.l.Has(g.l.Arrow(), Termite) && g.cfg.Wump3 {
		g.p.Println(dia.EatenArrow)
		g.p.Print(dia.ChoiceShootMove)
		return waitShootMove
	}

	if g.arrowTravel == 0 {
		// 3 out of 4 time an arrow falling will have an effect on the wumpus.
		if g.r.Intn(4) != 0 && !g.inventory.has(Wumpus) {
			if g.wumpusAwake {
				g.p.Println(dia.ArrowStartle)
				g.l.StartleWumpus() // make the wumpus move randomly
				// check 1/20 odds that the wumpus moved to player's cavern
				if g.l.Has(g.l.Player(), Wumpus) {
					g.p.Println(dia.WumpusTrample)
					g.p.Print(dia.PlayAGain)
					return waitPlayAgain
				}
			} else {
				g.p.Println(dia.ArrowWakeup)
				g.wumpusAwake = true
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
	g.p.Printf(dia.Room, g.l.Name(g.l.Player()), g.l.PlayerPOV())
}

func (g *Game) describe() {
	g.p.Printf(dia.Tunnels, g.l.GetFmtNeighbors(g.l.Player()))
	if g.l.Nearby(Bat) {
		g.p.Println(dia.BatsNearby)
	}
	if g.l.Nearby(Pit) {
		g.p.Println(dia.PitsNearby)
	}
	if g.l.Nearby(Wumpus) && !g.inventory.has(Wumpus) {
		g.p.Println(dia.WumpusNearby)
	}

	if g.cfg.Wump3 && g.l.Nearby(Termite) {
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
	if !g.cfg.Wump3 {
		return
	}

	if g.r.Intn(randEvent) == 0 {
		g.l.Migration(Pit)
		g.p.Println(dia.Earthquake)
	}

	if g.r.Intn(randEvent) == 0 {
		g.l.Migration(Bat)
		g.p.Println(dia.BatMigration)
	}

	if g.r.Intn(randEvent) == 0 {
		g.l.Migration(Termite)
		g.p.Println(dia.TermiteMigration)
	}

	if g.r.Intn(randWumpus) == 0 && g.wumpusAwake { // lower probability
		g.l.MigrateWumpus() // move 1 cave over
		g.p.Println(dia.SleepWalkWumpus)
	}
}

func (g *Game) explore() bool {
	g.p.Printf(dia.MovedTo, g.l.Name(g.l.Player()), g.l.PlayerPOV())
	return g.hazards()
}

func (g *Game) items() {
	if !g.cfg.Advanced {
		return
	}

	if g.l.Has(g.l.Player(), Clue) {
		loc, sub := g.l.GetClue(g.l.Player())
		g.p.Printf(dia.FoundClue, g.p.Get(subject(sub)), loc)
	}

	if g.l.FoundObject(Repel) {
		g.inventory.add(Repel)
		g.p.Println(dia.FoundRepel)
	}

	if g.l.FoundObject(Rope) {
		g.inventory.add(Rope)
		g.p.Println(dia.FoundRope)
	}

	if g.l.FoundObject(Shield) {
		g.inventory.add(Shield)
		g.p.Println(dia.FoundShield)
	}

	if g.r.Intn(randMaps) == 0 {
		g.p.Printf(dia.PartialMap, g.l.GetFmtMap())
	}

	if g.r.Intn(randArrows) == 0 {
		g.inventory.add(Arrow)
		g.p.Println(dia.FoundArrow)
	}
}

// hazards checks for wumpus/bats/pits when entering a new room.
// Return true if a hazard killed the player.
// If a bat moves the player, call recursively.
func (g *Game) hazards() bool {
	// the wumpus is immune to hazards, so we check for it first
	if g.l.Has(g.l.Player(), Wumpus) && !g.inventory.has(Wumpus) {
		g.p.Println(dia.StumbledWumpus)
		if attack := g.l.FoundWumpus(g.wumpusAwake); attack {
			if g.inventory.tryUse(Shield) {
				g.p.Println(dia.UseShield) // the wumpus is relocated in any case
			} else {
				g.p.Println(dia.KilledByWumpus)
				return true
			}
		} else {
			g.p.Println(dia.StartledWumpus)
		}
		g.wumpusAwake = true
	}

	// the bat may teleport to a pit or the wumpus, so we check it second
	if g.l.Has(g.l.Player(), Bat) {
		if g.inventory.tryUse(Repel) {
			g.p.Println(dia.UseRepel)
		} else {
			newLoc := g.l.ActivateBat()
			g.p.Printf(dia.BatTeleport, g.l.Name(g.l.Player()), newLoc)
			return g.hazards()
		}
	}

	if g.l.Has(g.l.Player(), Pit) {
		if g.inventory.tryUse(Rope) {
			g.p.Println(dia.UseRope)
		} else {
			g.p.Println(dia.FellIntoPit)
			g.p.Printf(dia.ExitWumpus, g.l.Wumpus())
			return true
		}
	}

	if g.cfg.Wump3 && g.l.Has(g.l.Player(), Termite) && g.inventory.tryUse(Arrow) {
		g.p.Println(dia.TermiteEatArrow)
		g.p.Printf(dia.RemainingArrows, g.inventory.count(Arrow))
	}

	return false
}

// keyDoor resolve dialogues & handle logic for key & door depending on the order of discovery.
// return true if all winning conditions are met
func (g *Game) keyDoor() bool {
	if !g.cfg.Advanced {
		return false
	}

	doorRoom := g.l.Has(g.l.Player(), Door)
	keyRoom := g.l.Has(g.l.Player(), Key)

	if !doorRoom && !keyRoom {
		return false
	}

	canUnlock := false // in the door room with the key.
	switch {
	case doorRoom && g.inventory.has(Key) && g.inventory.has(Door):
		// found the door, then the key, and are back to the room with the door
		g.p.Println(dia.DoorKeyDoor)
		canUnlock = true
	case doorRoom && g.inventory.has(Key):
		// found the key first then the door (first time seeing it)
		g.p.Println(dia.KeyThenDoor)
		g.inventory.add(Door)
		canUnlock = true
	case doorRoom && !g.inventory.has(Door):
		// first time seeing the door, no key
		g.p.Println(dia.FirstDoorDiscoveryNoKey)
		g.inventory.add(Door)
	case doorRoom:
		// back in the cavern with the door again
		g.p.Println(dia.BackAgainDoorNoKey)
	case keyRoom && g.inventory.has(Door) && !g.inventory.has(Key):
		// found the door first, then this key
		g.p.Println(dia.DoorThenKey)
		g.inventory.add(Key)
	case keyRoom && !g.inventory.has(Door) && !g.inventory.has(Key):
		// found the key first
		g.p.Println(dia.FirstKeyDiscoveryNoDoor)
		g.inventory.add(Key)
	}

	if canUnlock && !g.inventory.has(Wumpus) {
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
	case g.inventory.has(Key) && g.inventory.has(Door):
		g.p.Println(dia.CertainKeyDoor)
	case g.inventory.has(Key) && !g.inventory.has(Door):
		g.p.Println(dia.MaybeKey)
	case !g.inventory.has(Key) && g.inventory.has(Door):
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

func parsePower(input string) int {
	p, err := strconv.Atoi(input)
	if err != nil || p > 5 || p == 0 {
		p = 5
	}
	return p
}

func subject(e Entity) string {
	switch e {
	case Pit:
		return dia.SubjectPit
	case Bat:
		return dia.SubjectBat
	case Wumpus:
		return dia.SubjectWumpus
	case Key:
		return dia.SubjectKey
	case Door:
		return dia.SubjectDoor
	case Termite:
		return dia.SubjectTermite
	}
	return ""
}
