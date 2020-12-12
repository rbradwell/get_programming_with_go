package main

/*
Use structural embedding below.  This allows go to do method forwarding for us rather than doing it manually.

All methods of the temperature type are automatically made accessible through the report type
*/
/*
type report struct {
	sol int
	//	temperature temperature
	//	location    location
	temperature
	location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {

	report := report{sol: 15,
		temperature: temperature{high: -1.0, low: -78.0},
		location:    location{-4.5895, 137.4417}}

	fmt.Printf("Average %v degrees C\n", report.average())

	// note: although no field name was specified, a field exists with the same name as the embedded type.
	// the temperature field can be accessed as follows

	fmt.Printf("Average %v degrees C\n", report.temperature.average())

	fmt.Printf("%v degrees C\n", report.high)

	report.high = 32

	fmt.Printf("%v degrees C\n", report.temperature.high)

}
*/
