package main

/* it can be tempting to adopt nil whenever a value can be nothing.
e.g. a pointer to an integer (*int)can represent both zero and nil.
This isn't necessarily the best option.
Instead of using a pointer, one alternative is to declare a small structure with
a fe methods.  It requires a bit more code, but don't require a pointer or nil
*/
/*
type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}

func main() {
	n := newNumber(42)
	fmt.Println(n)
	e := number{}
	fmt.Println(e)
}
*/
