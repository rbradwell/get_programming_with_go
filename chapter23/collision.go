package main

/*
// any types can be embedded in a structure, not just structures
type sol int

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type report struct {
	sol
	location
	temperature
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (l location) days(l2 location) int {
	// TODO - complicated calculation
	return 5
}

// this was necessary to fix ambiguous delegation
// try commenting it out
func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

func main() {

	report := report{sol: 15}
	fmt.Println(report.sol.days(1446))
	// this is ambiguous now because both sol and location has days
	// GO doesn't know which method to delegate to
	// This has been resolved by implementing a delegation method days on the report type
	fmt.Println(report.days(1446))
}
*/
