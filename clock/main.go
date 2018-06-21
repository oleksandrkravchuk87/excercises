package main

import (
	"fmt"
)

type I interface {
	some()
}

type T struct {
}

func (t T) some() {
	fmt.Println("some")
}

func main() {
	var i I
	i = &T{}
	i.some()

}
