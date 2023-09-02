package labyrinth

import (
	"bytes"
	"embed"
	"encoding/json"
)

const folder = "levels"

//go:embed levels/*.json
var levelsDir embed.FS

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
