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

// internal representation of the settings per level.
type setup struct {
	nbClue    int
	nbTermite int
	nbPit     int
	nbBat     int
	nbShield  int
	nbRope    int
	nbRepel   int
}

// Setup is the optional json definitions for # of various entities.
type Setup struct {
	Bat     *int `json:"bat,omitempty"`
	Pit     *int `json:"pit,omitempty"`
	Termite *int `json:"termite,omitempty"`
	Clue    *int `json:"clue,omitempty"`
	Repel   *int `json:"repel,omitempty"`
	Rope    *int `json:"rope,omitempty"`
	Shield  *int `json:"shield,omitempty"`
}

// parse loads the default values for various # of entities, then overload with the json value if defined.
func (s *Setup) parse() setup {
	p := setup{
		nbClue:    nbClues,
		nbTermite: nbTermites,
		nbPit:     nbPits,
		nbBat:     nbBats,
		nbShield:  nbShield,
		nbRope:    nbRope,
		nbRepel:   nbRepel,
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
	return p
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
