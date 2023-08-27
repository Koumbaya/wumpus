package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	dia "github.com/koumbaya/wumpus/dialogues"
	"github.com/koumbaya/wumpus/labyrinth"
)

type state int

const (
	waitShootMove state = iota
	waitWhereTo
	waitPlayAgain
	waitArrowWhereTo
)

type Game struct {
	l labyrinth.Labyrinth
	state
}

func NewGame(l labyrinth.Labyrinth) Game {
	return Game{
		l:     l,
		state: waitShootMove,
	}
}

func (g *Game) Loop() {
	reader := bufio.NewReader(os.Stdin)
	g.start()

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = clean(input)

		if strings.EqualFold(input, "exit") {
			return
		}

		g.playerState(input)
	}
}

// playerState is the main state machine.
func (g *Game) playerState(input string) {
	switch g.state {
	case waitShootMove:
		if strings.EqualFold(input, "S") {
			g.l.FireArrow()
			fmt.Println(dia.FireArrow)
			g.whereToArrow()
			g.state = waitArrowWhereTo
		} else if strings.EqualFold(input, "M") {
			g.whereTo()
			g.state = waitWhereTo
		} else {
			fmt.Println(dia.DontUnderstand)
			g.choiceSM()
		}
	case waitWhereTo:
		if !g.tryMove(input) {
			break // error parsing
		}
		if g.explore() { //dead
			fmt.Println(dia.PlayAGain)
			g.state = waitPlayAgain
			break
		}
		g.describe()
		g.choiceSM()
		g.state = waitShootMove
	case waitArrowWhereTo:
		if !g.tryArrow(input) {
			break
		}
		g.state = g.handleArrow()
	case waitPlayAgain:
		if strings.EqualFold(input, "Y") {
			g.l.Init()
			g.start()
			g.state = waitShootMove
		} else {
			fmt.Println("Goodbye")
			os.Exit(0)
		}
	}
}

func (g *Game) tryArrow(input string) bool {
	d, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(dia.NotNumber)
		g.whereToArrow()
		return false
	}

	g.l.MoveArrow(d - 1)
	return true
}

func (g *Game) handleArrow() state {
	fmt.Printf(dia.ArrowTravel, g.l.Arrow()+1)
	if g.l.HasWumpus(g.l.Arrow()) {
		fmt.Println(dia.KilledWumpus)
		fmt.Println(dia.PlayAGain)
		return waitPlayAgain
	}

	if g.l.Player() == g.l.Arrow() {
		fmt.Println(dia.ArrowPlayer)
		fmt.Println(dia.PlayAGain)
		return waitPlayAgain
	}

	if g.l.PowerRemaining() == 0 {
		if g.l.StartleWumpus() {
			fmt.Println(dia.ArrowStartle)
			// check 1/20 odds that the wumpus moved to player's cavern
			if g.l.HasWumpus(g.l.Player()) {
				fmt.Println(dia.WumpusTrample)
				fmt.Println(dia.PlayAGain)
				return waitPlayAgain
			}
		} else {
			fmt.Println(dia.ArrowFell)
		}
		g.choiceSM()
		return waitShootMove
	}
	g.whereToArrow()
	return waitArrowWhereTo
}

func (g *Game) start() {
	fmt.Println(dia.Start)
	g.cavern()
	g.describe()
	g.choiceSM()
}

func (g *Game) tryMove(input string) bool {
	d, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(dia.NotNumber)
		g.whereTo()
		return false
	}
	moved := g.l.TryMovePlayer(d - 1)
	if !moved {
		fmt.Println(dia.NotValidDest)
		g.whereTo()
		return false
	}

	return true
}

func (g *Game) cavern() {
	fmt.Printf(dia.Room, g.l.Player()+1)
}

func (g *Game) describe() {
	fmt.Printf(dia.Tunnels, g.l.GetFmtNeighbors(g.l.Player()))
	if g.l.BatsNearby() {
		fmt.Println(dia.BatsNearby)
	}
	if g.l.PitNearby() {
		fmt.Println(dia.PitsNearby)
	}
	if g.l.WumpusNearby() {
		fmt.Println(dia.WumpusNearby)
	}
}

func (g *Game) choiceSM() {
	fmt.Print(dia.ChoiceShootMove)
}

func (g *Game) whereTo() {
	fmt.Printf(dia.WhereTo,
		g.l.Rooms[g.l.Player()].Neighbors[0]+1,
		g.l.Rooms[g.l.Player()].Neighbors[1]+1,
		g.l.Rooms[g.l.Player()].Neighbors[2]+1,
	)
}

func (g *Game) whereToArrow() {
	fmt.Printf(dia.WhereToArrow,
		g.l.Rooms[g.l.Arrow()].Neighbors[0]+1,
		g.l.Rooms[g.l.Arrow()].Neighbors[1]+1,
		g.l.Rooms[g.l.Arrow()].Neighbors[2]+1,
	)
}

func (g *Game) explore() bool {
	fmt.Printf(dia.MovedTo, g.l.Player()+1)
	return g.hazards()
}

// hazards checks for wumpus/pits/bats when entering a new room.
// Return true if a hazard killed the player.
// If a bat moves the player, call recursively.
func (g *Game) hazards() bool {
	// the wumpus is immune to hazards, so we check for it first
	if g.l.HasWumpus(g.l.Player()) {
		fmt.Println(dia.StumbledWumpus)
		if dead := g.l.FoundWumpus(); dead {
			fmt.Println(dia.KilledByWumpus)
			return true
		}
		fmt.Println(dia.StartledWumpus)
	}

	// the bat may teleport to a pit or the wumpus, so we check it second
	if g.l.HasBat(g.l.Player()) {
		fmt.Printf(dia.BatTeleport, g.l.ActivateBat())
		return g.hazards()
	}

	if g.l.HasPit(g.l.Player()) {
		fmt.Println(dia.FellIntoPit)
		return true
	}

	return false
}

func clean(input string) string {
	input = strings.TrimRight(input, "\n")
	input = strings.ReplaceAll(input, " ", "")
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
