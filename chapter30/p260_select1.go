package main

/*
// use channels to coordinate so that main waits until all goroutines done
// only wait for so long before giving up
func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id // send a value back to main through channel
}

func main() {
	// make communication channel
	c := make(chan int)

	// After returns a channel that receives a value after some time has passed
	timeout := time.After(3 * time.Second)

	// start goroutines
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
		// select holds a channel receive or send.  select waits until on case is ready then run that associated code
		// it's as if select is looking at both channels at once and takes action when it seems something happening on either of the channels
		// when waiting to send or receive on a channel, we say it's blocked
		select {
		case gopherID := <-c: // receive a value from gopher channel
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout:
			fmt.Println("my patience has ran out")
			return
		}

	}
}
*/
