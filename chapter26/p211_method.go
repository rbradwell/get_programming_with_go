package main

/*
type person struct {
	name string
	age  int
}

// methid receivers are similar to parameters.  Below uses a pointer for the receiver
// which allows the method to mutate a persons attributes
func (p *person) birthday() {
	p.age++
}

// pointers are used to enable mutation across function and method boundaries

func main() {

	terry := &person{
		name: "Terry",
		age:  14,
	}

	terry.birthday()
	fmt.Printf("%+v\n", terry) // &{name:Terry age:15}

	// the following doesn't use a pointer, yet it still works
	// GO will automatically determine the address (&) of a variable when calling methods with dot notation
	// therefore there is no need to write (&nathan).birthday()

	nathan := person{
		name: "Nathan",
		age:  17,
	}
	nathan.birthday()
	fmt.Printf("%+v\n", nathan) // {name:Nathan age:18}
}
*/
