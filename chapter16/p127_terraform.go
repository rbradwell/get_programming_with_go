package main

/*
// terraform does nothing because arrays are values and functions pass by value
// i.e. terraform is operating on a copy of the planets
// NOTE: the array length is part of the type and must be equal to the number of elements
// Therefore arrays aren't nearly as useful as function parameters as slices
func terraform(planets [3]string) {
	for i := range planets {
		planets[i] = "New" + planets[i]
	}
}

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
	}
	terraform(planets)
	fmt.Println(planets)
}
*/
