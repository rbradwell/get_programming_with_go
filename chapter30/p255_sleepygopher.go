package main

/**
An independently runnint task is known as a goroutine.  They are similar to coroutines, threads etc but not the same.
They can communcate with each other through channels.
**/
/*
func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
}

func main() {
	// start goroutines
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}
*/
