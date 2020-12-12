package main

/*
You can define a type with an underlying slice or array
For example, the sort package in the standard library declares a StringSlice type

type StringSlice []string

Attached to StringSlice is a Sort method:

func (p StringSlice) Sort()

*/
/*
func main() {
	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Nepune",
	}
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)

	// to make it even simpler the sort package has helper function that performs the type conversion and calls the sort method for you
	sort.Strings(planets)

}
*/
