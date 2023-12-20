package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 200ms"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			c2 <- "every 2 secods"
			time.Sleep(time.Second * 2)
		}
	}()

	for {

		select {
		case msg1 := <-c1:
			fmt.Println("msg from c1 chan----", msg1)

		case msg2 := <-c2:
			fmt.Println("msg from c2 chan ---", msg2)
		}

	}

}
