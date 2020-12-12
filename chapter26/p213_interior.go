package main

/*
type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func main() {
	// go provides a feature called interior pointers, used to determine the memory address of a field inside a structure
	player := character{name: "Matthias"}
	levelUp(&player.stats)            // this provides a pointer to the interior of the structure
	fmt.Printf("%+v\n", player.stats) // {level:1 endurance:56 health:280}
}
*/
