package main

/*
func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	// three index slicing introduces a capacity
	terrestrial := planets[0:4:4] // NOTE: capacity of 4
	// appending Ceres creates a new underlying array because the capacity is maxed out
	worlds := append(terrestrial, "Ceres")
	// the original array is still the same because a copy was taken
	fmt.Println(planets) // Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
	fmt.Println(worlds)  // Mercury Venus Earth Mars Ceres]
	// this time is will change the underlying array because it has capacity and therefore no copy is taken
	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	fmt.Println(worlds)  // [Mercury Venus Earth Mars Ceres]
	fmt.Println(planets) // [Mercury Venus Earth Mars Ceres Saturn Uranus Neptune]
}
*/
