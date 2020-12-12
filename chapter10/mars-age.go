package main

/*
trying to join a number to a string in go will report an error i.e. it won't try cooerce the types

type conversions are required between signed and unisgned integer types, and between types of different sizes.
it's always safe to convert to a type with a larger range e.g. int8 to int32.  Other conversions come with risk.
*/
/*
func main() {
	age := 41
	marsAge := float64(age)
	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("I am ", marsAge, "years old on Mars.")
	// you can convert from a floating point to an integer, though digits after the decimal point will be truncated without rounding
	fmt.Println(int((earthDays)))
}
*/
