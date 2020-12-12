package main

/*
// How we can reffactor code to use interface and make the code more generic

// this function is limited to earth dates
func stardate(t time.Time) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

// therefore use interface to make it more generic
type stardater interface {
	YearDay() int
	Hour() int
}

// stardate returns a fictional measure of time
func stardate2(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668) // there are 668 sols in a martian year
}

func (s sol) Hour() int {
	return 0
}

func main() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	// the new stardate2 function still works on earth dates because time.Time type satisfies the stardater interface
	// NOTE: code that you didn't write can satisfy an interface implicitly. Was this what default interfaces in java were introduced for?
	fmt.Printf("%.1f Curiosity has landed\n", stardate2(day))

	// stardate2 works on both Earth dates and Martian sols
	s := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", stardate2(s))
}
*/
