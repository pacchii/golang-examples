package _4_channels

import (
	"fmt"
	"time"
)

func BlockExample() {

	ch := make(chan int)

	go func() {
		fmt.Println("sending...")
		ch <- 1
		fmt.Println("sent")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Receiving...")
	fmt.Println(<-ch)
	fmt.Println("received")

	time.Sleep(2 * time.Second)
}

/*

sending...
Receiving...
1
received
sent

*/
