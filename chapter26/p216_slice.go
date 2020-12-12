package main

/*
func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func main() {
	// slices point to arrays
	// to point at an element of the array, wlices use a pointer
	// a slice represents internally a structure with three element: a pointer to an array, the capacity of the slice and length
	// the internal pointer allows the underlying data to be mutated when a slice is passed directly to a function or method
	// an explicit pointer to a slice is only useful when modifying the slice itself i.e. length, capacity or starting offset.
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets) // remove Pluto from planet classification
	fmt.Println(planets) // [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
}
*/
