package main

/*
import (
	"fmt"
	"strings"
)

type laser int

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

// an interface can be used anywhere other types are ussed

func shout(t talker) {
	louder := strings.ToLower(t.talk())
	fmt.Println(louder)
}

func main() {
	// you can use the shout function what any value that satisfies the talker interface
	shout(martian{})
	shout(laser(2))
}
*/
