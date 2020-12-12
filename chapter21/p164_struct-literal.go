package main

/*
// initialize structures with composite literals
func main() {
	type location struct {
		lat, long float64
	}
	// first technique ...
	// fields may be in any order, fields that aren't listed will retain zero value
	// this work if extra fields are added later or fields reordered
	opportunity := location{lat: -1.9462, long: 354.4734}

	insight := location{lat: 4.5, long: 135.9}

	// second technique ..
	// provide all values in correcr order.
	// this doesn't work if fields are added or reordered
	spirit := location{-14.5684, 175.472636}

	// the %v format verb can be modified with a plus sign + to print out field names

	fmt.Printf("%+v\n", opportunity)
	fmt.Printf("%+v\n", insight)
	fmt.Printf("%+v\n", spirit)

}
*/
