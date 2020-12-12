package main

/*
func main() {
	// when there isn't enough capacity for append go copies the contents into a new array with sufficient capacity
	// to prevent this a slice capacity can be preallocated with make.
	planets := make([]string, 0, 10) // start with capacity of 10
	planets = []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	planets = append(planets, "Planet X")
	fmt.Println(planets)
}
*/
