package main

// use channels to coordinate so that main waits until all goroutines done
/*
func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id // send a value back to main through channel
}

func main() {
	// make communication channel
	c := make(chan int)
	// start goroutines
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
		gopherID := <-c // receive a value from channel
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}
*/
