package main

/*
Both worlds and planets are slices, and though worlds is a copy, they both point to the same underlying array
changes by the worlds slice is visit by other slices of the array
unlike arrays the length of an array isn't part of the type, therefore a slice of any size can be passed to function
*/
/*
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
func main() {
	planets := []string{"     Venus     ", "      Earth", "Mars      "}
	hyperspace(planets)
	fmt.Println(strings.Join(planets, "")) // VenusEarthMars
}
*/
