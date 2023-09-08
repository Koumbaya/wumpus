package dialogues

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const textDelay = 20 * time.Millisecond

//go:embed dialogues.json
var dialogueJSON embed.FS

// dialogueVariations holds the different possibles values for a particular piece of dialogue.
type dialogueVariations struct {
	values []string
}

type Printer struct {
	noDelay   bool
	dialogues map[string]dialogueVariations
	r         *rand.Rand
}

func NewPrinter(noDelay, clean bool, seed int64) *Printer {
	return &Printer{
		noDelay:   noDelay,
		dialogues: loadDialogues(clean),
		r:         rand.New(rand.NewSource(seed)),
	}
}

func (p *Printer) Printf(key string, a ...any) {
	value := p.getRandomVal(key)
	if p.noDelay {
		fmt.Printf(value, a...)
		return
	}

	r := fmt.Sprintf(value, a...)
	for _, c := range r {
		time.Sleep(textDelay)
		fmt.Print(string(c))
	}
}

func (p *Printer) Print(key string) {
	value := p.getRandomVal(key)
	if p.noDelay {
		fmt.Print(value)
		return
	}

	for _, c := range value {
		time.Sleep(textDelay)
		fmt.Print(string(c))
	}
}

func (p *Printer) Println(key string) {
	value := p.getRandomVal(key)
	if p.noDelay {
		fmt.Println(value)
		return
	}

	for _, c := range value {
		time.Sleep(textDelay)
		fmt.Print(string(c))
	}

	fmt.Println()
}

// loadDialogues parse the json values and put them in a map for instant access.
func loadDialogues(clean bool) map[string]dialogueVariations {
	content, err := dialogueJSON.ReadFile("dialogues.json")
	if err != nil {
		panic(err)
	}

	var dialogues struct {
		Data []struct {
			Key    string   `json:"key"`
			Values []string `json:"values"`
			Color  string   `json:"color"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(bytes.NewReader(content))
	if err := decoder.Decode(&dialogues); err != nil {
		panic(err)
	}

	res := make(map[string]dialogueVariations, len(dialogues.Data))
	for i := 0; i < len(dialogues.Data); i++ {
		if clean {
			// disable all formatting on windows. allow cross-compile without build flags or duplicated files.
			res[dialogues.Data[i].Key] = dialogueVariations{
				values: removeSpecialChars(dialogues.Data[i].Values),
			}
		} else {
			res[dialogues.Data[i].Key] = dialogueVariations{
				values: color(formatSpecialChars(dialogues.Data[i].Values), dialogues.Data[i].Color),
			}
		}
	}

	return res
}

func formatSpecialChars(s []string) []string {
	for i, s2 := range s {
		s2 = strings.ReplaceAll(s2, SkullPlcHolder, SkullChar)
		s2 = strings.ReplaceAll(s2, ArrowDownPlcHolder, ArrowDownChar)
		s2 = strings.ReplaceAll(s2, ArrowStraightPlcHolder, ArrowStraightChar)
		s2 = strings.ReplaceAll(s2, ArrowUpPlcHolder, ArrowUpChar)
		s[i] = s2
	}

	return s
}

func removeSpecialChars(s []string) []string {
	for i, s2 := range s {
		s2 = strings.ReplaceAll(s2, SkullPlcHolder, "*")
		s2 = strings.ReplaceAll(s2, ArrowDownPlcHolder, "->")
		s2 = strings.ReplaceAll(s2, ArrowStraightPlcHolder, "->")
		s2 = strings.ReplaceAll(s2, ArrowUpPlcHolder, "->")
		s[i] = s2
	}

	return s
}

// getRandomVal returns one of the dialogue at random for a given key.
func (p *Printer) getRandomVal(key string) string {
	k := len(p.dialogues[key].values)
	if k == 0 {
		return "missing values in dialogue.json for key:" + key
	}
	return p.dialogues[key].values[p.r.Intn(k)]
}

func color(s []string, color string) []string {
	if color == "" {
		return s
	}

	for i := range s {
		s[i] = mapColors(color) + s[i] + Reset
	}

	return s
}

func mapColors(s string) string {
	switch s {
	case "Reset":
		return Reset
	case "dim":
		return Dim
	case "red":
		return red
	case "yellow":
		return yellow
	case "cyan":
		return cyan
	case "bold":
		return bold
	case "cyan+dim":
		return cyan + Dim
	case "yellow+dim":
		return yellow + Dim
	case "red+bold":
		return red + bold
	}
	return ""
}
