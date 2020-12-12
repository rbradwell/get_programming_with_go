package main

/*
// creating a pipepline using goroutines

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream) // close the channel to signify no more values will be sent
}

func filterGopher(upstream, downstream chan string) {
	for {
		// second argument here indicates if we successfully read from channel.  If channel is closed this will be false
		item, ok := <-upstream
		if !ok {
			close(downstream)
			return
		}
		// filter bad apple
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

// this is a rewrite of the filterGopher method above using the range statement
func filterGopher2(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher2(c0, c1)
	printGopher(c1)
}
*/
