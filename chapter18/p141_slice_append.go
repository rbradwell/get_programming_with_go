package main

/*
import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	// slice on an array
	dwarfs1 := []string{
		"Ceres",
		"Pluto",
		"Haumea",
		"Makemake",
		"Eris",
	}

	//the array backing dwarf1 doesn't have enough roon (capacity) to append "Orcus", so append copies the contents
	// of dwarf1 to a freshly allocated array with twice the capacity
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	dump("dwarfs", dwarfs1)  // dwarfs: length 5, capacity 5 [Ceres Pluto Haumea Makemake Eris]
	dump("dwarfs2", dwarfs2) // dwarfs2: length 6, capacity 10 [Ceres Pluto Haumea Makemake Eris Orcus]
	dump("dwarfs3", dwarfs3) // dwarfs3: length 9, capacity 10 [Ceres Pluto Haumea Makemake Eris Orcus Salacia Quaoar Sedna]

	// append allocates a new array with increased capacity when necessary

	// just to prove that views are now operating on different underlying arrays i.e. the array was copied when the
	// append took it over capacity
	// the following mutation only appears in the array that dwarfs3 slice is viewing
	dwarfs3[1] = "Pluto!"
	dump("dwarfs", dwarfs1)  // dwarfs: length 5, capacity 5 [Ceres Pluto Haumea Makemake Eris]
	dump("dwarfs2", dwarfs2) // dwarfs2: length 6, capacity 10 [Ceres Pluto Haumea Makemake Eris Orcus]
	dump("dwarfs3", dwarfs3) // dwarfs3: length 9, capacity 10 [Ceres Pluto! Haumea Makemake Eris Orcus Salacia Quaoar Sedna]

}
*/
