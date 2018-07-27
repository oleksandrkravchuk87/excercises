package main

import (
	"fmt"
)

// Clock API:
//
type Clock struct {
	m int
}

func New(hour, minute int) Clock {

	c := Clock{(hour*60 + minute) % (24 * 60)}

	if c.m < 0 {
		c.m = 24*60 + c.m
	}

	return c
}

func (c Clock) String() string {
	h := c.m / 60
	m := c.m % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, c.m+minutes)
}

func main() {
	c := New(-25, -160)
	fmt.Println(c)
}
