package main

/*
func main() {
	// when a variable is declared to be an interface type without an assignment, the zero value is nil
	// below both the interface and value type are nil.
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil) // nil nil true
	// when a variable with an interface type is assigned a value, the interface internally points to the type and value of that variable.
	// this surprisingly leads to behaviour where a nil value doesn't compare as equal to nil.
	// Both the interface type and value need to be nil for the variable to equal nil
	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil)
	// the #v format verb is shorthand to see both type and value, also revealing the variable contains (*int)(nil) rather than just <nil>
	fmt.Printf("%#v\n", v)
}
*/
