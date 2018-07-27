package main

import (
	"log"
	"math/rand"
	"time"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeConsumerAtSeat(customerId int, seat Seat) {
	log.Print("consumer#", customerId, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("<- consumer#", customerId, " frees seat#", seat)
	bar <- seat // free the seat and leave the bar
}

func main() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10) // the bar has 10 seats
	// Place seats in an bar.
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId) // none of the sends will block
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		seat := <-bar24x7 // need a seat to serve next consumer
		go bar24x7.ServeConsumerAtSeat(customerId, seat)
	}
	for {
		time.Sleep(time.Second)
	} // sleeping != blocking
}
