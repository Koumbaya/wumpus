package labyrinth

func (l *Labyrinth) Key() int {
	return l.key
}

func (l *Labyrinth) HasKey(n int) bool {
	return n == l.key
}

func (l *Labyrinth) Door() int {
	return l.door
}

func (l *Labyrinth) HasDoor(n int) bool {
	return n == l.door
}
