package main

/*
func main() {

	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0}

	// Go doesn't have Sets! Therefore we improvise using maps
	set := make(map[float64]bool)

	for _, t := range temperatures {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("set member")
	}

	// map keys have an arbitrart order in GO (no hashMap like java?)
	// therefore convert back to slice before sorting

	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}

	sort.Float64s(unique)
	fmt.Println(unique)
}
*/
