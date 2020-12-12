package main

/*
import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

// the variable t can hold any value of any type that satisfies the interface.
// a type will satisfy the interface if it declares a method named talk that accepts no arguments and returns a string

// both the martian and laser type below meet these requirements
// therefore both can be assigned to the variable t

// NOTE: neither martian or laser type explicitly declare that they implement tghe interface
type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}
func main() {
	t = martian{}
	fmt.Println(t.talk()) // nack nack
	t = laser(3)          // pew pew pew
	fmt.Println(t.talk())
}
*/
