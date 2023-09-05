package labyrinth

import (
	"bytes"
	"embed"
	"encoding/json"
)

const folder = "levels"

//go:embed levels/*.json
var levelsDir embed.FS

// internal representation of a level.
type level struct {
	rooms  []room
	number int
	name   string
	setup  setup
}

func loadLevels() map[int]level {
	// this could be avoided by directly putting json tags on the room and level struct,
	// but that way it's separated cleanly, allowing future changes.

	entries, err := levelsDir.ReadDir(folder)
	if err != nil {
		panic(err) //todo: return errors
	}

	levels := make(map[int]level, len(entries))

	for _, entry := range entries {
		var lvl struct {
			Name   string `json:"name"`
			Number int    `json:"level"`
			Setup  Setup  `json:"setup"`
			Rooms  []struct {
				Id    int   `json:"id"`
				Edges []int `json:"edges"`
			} `json:"rooms"`
		}
		// Read the embedded JSON data
		content, err := levelsDir.ReadFile(folder + "/" + entry.Name())
		if err != nil {
			panic(err) //todo: return errors
		}

		decoder := json.NewDecoder(bytes.NewReader(content))
		if err := decoder.Decode(&lvl); err != nil {
			panic(err)
		}

		l := level{
			number: lvl.Number,
			name:   lvl.Name,
			setup:  lvl.Setup.parse(),
			rooms:  make([]room, len(lvl.Rooms)),
		}

		for i := 0; i < len(lvl.Rooms); i++ {
			l.rooms[lvl.Rooms[i].Id] = room{edges: lvl.Rooms[i].Edges}
		}

		levels[l.number] = l
	}

	return levels
}

func (l *Labyrinth) CurrentLevel() int {
	// todo : name ?
	return l.curLevel
}

func (l *Labyrinth) HasNextLevel() bool {
	_, exist := l.levels[l.curLevel+1]
	return exist
}

// internal representation of the settings per level.
type setup struct {
	nbClue         int
	nbTermite      int
	nbPit          int
	nbBat          int
	nbShield       int
	nbRope         int
	nbRepel        int
	migrations     bool
	wumpusStartPos *int
	playerStartPos *int
	keyPos         *int
	doorPos        *int
	pitsPos        []int
	batsPos        []int
	termitePos     []int
	cluePos        []int
	shieldPos      []int
	ropePos        []int
	repelPos       []int
}

// Setup is the optional json definitions for # of various entities.
type Setup struct {
	Bat            *int   `json:"bat,omitempty"`
	Pit            *int   `json:"pit,omitempty"`
	Termite        *int   `json:"termite,omitempty"`
	Clue           *int   `json:"clue,omitempty"`
	Repel          *int   `json:"repel,omitempty"`
	Rope           *int   `json:"rope,omitempty"`
	Shield         *int   `json:"shield,omitempty"`
	Migrations     *bool  `json:"migrations,omitempty"`
	WumpusStartPos *int   `json:"wumpus_start_pos,omitempty"`
	PlayerStartPos *int   `json:"player_start_pos,omitempty"`
	KeyPos         *int   `json:"key_pos,omitempty"`
	DoorPos        *int   `json:"door_pos,omitempty"`
	PitsPos        *[]int `json:"pits_pos,omitempty"`
	BatsPos        *[]int `json:"bats_pos,omitempty"`
	TermitePos     *[]int `json:"termite_pos,omitempty"`
	CluePos        *[]int `json:"clue_pos,omitempty"`
	ShieldPos      *[]int `json:"shield_pos,omitempty"`
	RopePos        *[]int `json:"rope_pos,omitempty"`
	RepelPos       *[]int `json:"repel_pos,omitempty"`
}

// parse loads the default values for various # of entities, then overload with the json value if defined.
func (s *Setup) parse() setup {
	p := setup{
		nbClue:     nbClues,
		nbTermite:  nbTermites,
		nbPit:      nbPits,
		nbBat:      nbBats,
		nbShield:   nbShield,
		nbRope:     nbRope,
		nbRepel:    nbRepel,
		migrations: true,
	}
	if s.Bat != nil {
		p.nbBat = *s.Bat
	}
	if s.Pit != nil {
		p.nbPit = *s.Pit
	}
	if s.Termite != nil {
		p.nbTermite = *s.Termite
	}
	if s.Clue != nil {
		p.nbClue = *s.Clue
	}
	if s.Repel != nil {
		p.nbRepel = *s.Repel
	}
	if s.Rope != nil {
		p.nbRope = *s.Rope
	}
	if s.Shield != nil {
		p.nbShield = *s.Shield
	}
	if s.Migrations != nil {
		p.migrations = *s.Migrations
	}
	if s.PitsPos != nil {
		p.pitsPos = *s.PitsPos
	}
	if s.BatsPos != nil {
		p.batsPos = *s.BatsPos
	}
	if s.TermitePos != nil {
		p.termitePos = *s.TermitePos
	}
	if s.CluePos != nil {
		p.cluePos = *s.CluePos
	}
	if s.ShieldPos != nil {
		p.shieldPos = *s.ShieldPos
	}
	if s.RopePos != nil {
		p.ropePos = *s.RopePos
	}
	if s.RepelPos != nil {
		p.repelPos = *s.RepelPos
	}
	p.wumpusStartPos = s.WumpusStartPos
	p.playerStartPos = s.PlayerStartPos
	p.keyPos = s.KeyPos
	p.doorPos = s.DoorPos
	return p
}
