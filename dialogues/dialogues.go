package dialogues

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

//go:embed dialogues.json
var dialogueJSON embed.FS

type dialogueVariations struct {
	values []string
	color  string
}

type Printer struct {
	delay     time.Duration
	dialogues map[string]dialogueVariations
}

func NewPrinter(t time.Duration) *Printer {
	return &Printer{
		delay:     t,
		dialogues: loadDialogues(),
	}
}

func (p *Printer) Printf(key string, a ...any) {
	value := p.getRandomValColored(key)
	if p.delay == 0 {
		fmt.Printf(value, a...)
		return
	}

	r := fmt.Sprintf(value, a...)
	for _, c := range r {
		time.Sleep(p.delay)
		fmt.Print(string(c))
	}
}

func (p *Printer) Print(key string) {
	value := p.getRandomValColored(key)
	if p.delay == 0 {
		fmt.Print(value)
		return
	}
	for _, c := range value {
		time.Sleep(p.delay)
		fmt.Print(string(c))
	}
}

func (p *Printer) Println(key string) {
	value := p.getRandomValColored(key)
	if p.delay == 0 {
		fmt.Println(value)
		return
	}
	for _, c := range value {
		time.Sleep(p.delay)
		fmt.Print(string(c))
	}
	fmt.Println()
}

func loadDialogues() map[string]dialogueVariations {
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
		res[dialogues.Data[i].Key] = dialogueVariations{
			values: dialogues.Data[i].Values,
			color:  dialogues.Data[i].Color,
		}
	}

	return res
}

func (p *Printer) getRandomValColored(key string) string {
	return color(p.dialogues[key].values[rand.Intn(len(p.dialogues[key].values))], p.dialogues[key].color)
}

func color(s, color string) string {
	return mapColors(color) + s + reset
}

func mapColors(s string) string {
	switch s {
	case "reset":
		return reset
	case "dim":
		return dim
	case "red":
		return red
	case "yellow":
		return yellow
	case "cyan":
		return cyan
	case "bold":
		return bold
	}
	return ""
}
