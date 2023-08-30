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

type Line struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
	Color  string   `json:"color"`
}
type DialoguesJSON struct {
	Data []Line `json:"data"`
}

type dialogueInternal struct {
	values []string
	color  string
}

type Printer struct {
	delay     time.Duration
	dialogues map[string]dialogueInternal
}

func NewPrinter(t time.Duration) *Printer {
	content, err := dialogueJSON.ReadFile("dialogues.json")
	if err != nil {
		panic(err)
	}

	var dialogues DialoguesJSON
	decoder := json.NewDecoder(bytes.NewReader(content))
	if err := decoder.Decode(&dialogues); err != nil {
		panic(err)
	}

	p := &Printer{
		delay:     t,
		dialogues: make(map[string]dialogueInternal, len(dialogues.Data)),
	}

	// todo : disable color & special chars at runtime for windows (maybe?)
	for i := 0; i < len(dialogues.Data); i++ {
		p.dialogues[dialogues.Data[i].Key] = dialogueInternal{
			values: dialogues.Data[i].Values,
			color:  dialogues.Data[i].Color,
		}
	}

	return p
}

func (p *Printer) getRandomValColored(key string) string {
	return color(p.dialogues[key].values[rand.Intn(len(p.dialogues[key].values))], p.dialogues[key].color)
}

func color(s, color string) string {
	return mapColors(color) + s + reset
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
