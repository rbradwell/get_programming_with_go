package main

/*
Functions are first class.
In GO you can assign functions to variables, pass functions to functions, write functions that return functions
*/
/*
type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func main() {
	var sensor func() kelvin // declaring the variable explicitly i.e. without type inference
	sensor = fakeSensor      // pointer to a function
	fmt.Println(sensor())
	sensor = realSensor // pointer to a function
	fmt.Println(sensor())
}
*/
