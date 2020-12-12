package main

/*
func main() {
	var administrator *string
	scolese := "Christopher J. Scolese"
	administrator = &scolese
	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)

	// changes to the value of bolden can be made in one place, because the administrator variable
	// points to Bolden rather than storing a copy

	bolden = "Charles Frank Bolden Jr"
	fmt.Println(*administrator)

	// assigning major to administrator results in a new pointer that also points at Bolden string

	major := administrator
	*major = "Major General Charles Frank Bolden Jr"
	fmt.Println(bolden)

	// the major and administrator pointers both hold the same memory address and are therefore equal
	fmt.Println(administrator == major) // true

	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot

	fmt.Println(administrator == major) // false

	//	Assigning the dereferenced value of major to another variable makes a copy of the string.
	charles := *major
	//	After the clone is mad, direct and indirect modifications to bolden have no effect on the
	//	values of charles, or vise versa
	*major = "Charles Bolden"

	fmt.Println(charles)
	fmt.Println(bolden)

	// if two vars contai the same string, the are considered equal
	// this is the case even though they have different memory addresses
	charles = "Charles Bolden"
	fmt.Println(charles == bolden)   // true
	fmt.Println(&charles == &bolden) // false

}
*/
