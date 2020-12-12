package main

// creating a pipepline using goroutines
/*
func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	// here we are using an empty string to signify no more, but empty string may be a value we wish to process
	// the next example pipeline2.go will solve this issue
	downstream <- ""
}

func filterGopher(upstream, downstream chan string) {
	for {
		item := <-upstream
		if item == "" {
			downstream <- ""
			return
		}
		// filter bad apple
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

func printGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
*/
