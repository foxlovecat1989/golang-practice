package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	d1 := rand.Intn(250)
	d2 := rand.Intn(250)

	go func() {
		time.Sleep(time.Millisecond * time.Duration(d1))
		ch1 <- d1
	}()

	go func() {
		time.Sleep(time.Millisecond * time.Duration(d2))
		ch2 <- d2
	}()

	show(ch1, ch2)
}

func show(ch1 chan int, ch2 chan int) {
	select {
	case v1 := <-ch1:
		fmt.Printf("value from channel 1: %d", v1)
	case v2 := <-ch2:
		fmt.Printf("value from channel 2: %d", v2)
	}
}
