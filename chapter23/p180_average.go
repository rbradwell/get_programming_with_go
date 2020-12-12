package main

/*
type report struct {
	sol         int
	temperature temperature
	location    location
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

// this isn't needed, see embed.go
func (r report) average() celsius {
	return r.temperature.average()
}

func main() {

	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v degrees C\n", report.temperature.high)
	fmt.Printf("Average %v degrees C\n", t.average())

	// can also access average like this
	fmt.Printf("Average %v degrees C\n", report.temperature.average())

	// can also write a function thtat forwards to the real implementation e.g. delegation
	fmt.Printf("Average %v degrees C\n", report.average())

}
*/
