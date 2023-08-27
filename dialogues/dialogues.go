package dialogues

import (
	"fmt"
	"math/rand"
	"time"
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
