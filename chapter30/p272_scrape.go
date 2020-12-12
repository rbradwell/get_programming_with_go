package main

import "sync"

var mu sync.Mutex

// visited tracks whether web pages have been visited.  Usable from concurrent subroutines
type visited struct {
	// mu guards the visited map
	mu      sync.Mutex
	visited map[string]int
}

func (v *visited) visitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func main() {

}
