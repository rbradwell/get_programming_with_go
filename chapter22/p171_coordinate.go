package main

/*
type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println(lat.decimal(), long.decimal()) // -4.5895 137.4417

	curiosity := location{lat.decimal(), long.decimal()}
	fmt.Println(curiosity)

	// could also use a constructor function to build location e.g.
	// NOTE: there is nothing special about the constructor function
	curiosity = newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
}
*/
