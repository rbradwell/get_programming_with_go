package main

/*
type laser int

type startship struct {
	laser
}

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	// interfaces can be used with struct embedding
	s := startship{laser(3)}
	// embedding laser gives the startship a talk method that forwards to the laser
	fmt.Println(s.talk())
	// the starship also satisfies the talker interface allowing it to be used with shoud
	shout(s)
}
*/
