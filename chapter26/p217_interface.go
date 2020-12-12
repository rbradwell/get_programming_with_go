package main

import (
	"fmt"
	"strings"
)

type laser int

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func main() {
	pew := laser(2)
	shout(&pew) // PEW PEW

	// NOTE: can't just do this
	//shout(pew)
	// as in example p217_martian.go because
	// laser doesn't satisfy the interface in this case because it defined as
	// func (l *laser) talk() string
	// not
	// func (l laser) talk() string

}
