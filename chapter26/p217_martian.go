package main

/*
import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

func main() {
	// NOTE both martian and a pointer to martian satisfy the talker interface
	shout(martian{})  // NACK NACK
	shout(&martian{}) // NACK NACK
}
*/
