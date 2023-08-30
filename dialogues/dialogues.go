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

func BatTeleport() string {
	r := []string{
		BatTeleport0,
		BatTeleport1,
		BatTeleport2,
		BatTeleport3,
	}

	return r[rand.Intn(len(r))]
}

func Exit() string {
	r := []string{
		Exit0,
		Exit1,
		Exit2,
		Exit3,
	}

	return r[rand.Intn(len(r))]
}

func DoorKeyDoor() string {
	r := []string{
		DoorKeyDoor0,
		DoorKeyDoor1,
		DoorKeyDoor2,
		DoorKeyDoor3,
		DoorKeyDoor4,
	}

	return r[rand.Intn(len(r))]
}

func BackAgainDoorNoKey() string {
	r := []string{
		BackAgainDoorNoKey0,
		BackAgainDoorNoKey1,
		BackAgainDoorNoKey2,
		BackAgainDoorNoKey3,
		BackAgainDoorNoKey4,
	}

	return r[rand.Intn(len(r))]
}

func WumpusStillAlive() string {
	r := []string{
		WumpusStillAlive0,
		WumpusStillAlive1,
		WumpusStillAlive2,
		WumpusStillAlive3,
		WumpusStillAlive4,
	}

	return r[rand.Intn(len(r))]
}

func DoorThenKey() string {
	r := []string{
		DoorThenKey0,
		DoorThenKey1,
		DoorThenKey2,
		DoorThenKey3,
		DoorThenKey4,
		DoorThenKey5,
	}

	return r[rand.Intn(len(r))]
}

func FirstKeyDiscoveryNoDoor() string {
	r := []string{
		FirstKeyDiscoveryNoDoor0,
		FirstKeyDiscoveryNoDoor1,
		FirstKeyDiscoveryNoDoor2,
		FirstKeyDiscoveryNoDoor3,
		FirstKeyDiscoveryNoDoor4,
		FirstKeyDiscoveryNoDoor5,
	}

	return r[rand.Intn(len(r))]
}

func FirstDoorDiscoveryNoKey() string {
	r := []string{
		FirstDoorDiscoveryNoKey0,
		FirstDoorDiscoveryNoKey1,
		FirstDoorDiscoveryNoKey2,
		FirstDoorDiscoveryNoKey3,
		FirstDoorDiscoveryNoKey4,
	}

	return r[rand.Intn(len(r))]
}

func KeyThenDoor() string {
	r := []string{
		KeyThenDoor0,
		KeyThenDoor1,
		KeyThenDoor2,
		KeyThenDoor3,
		KeyThenDoor4,
	}

	return r[rand.Intn(len(r))]
}

func ExitDoor() string {
	r := []string{
		ExitDoor0,
		ExitDoor1,
		ExitDoor2,
		ExitDoor3,
	}

	return r[rand.Intn(len(r))]
}
