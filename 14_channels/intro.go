package _4_channels

import "fmt"

// not working
func ChannelExample() {

	ch := make(chan int, 5)
	/*
		//NOT WORKING

			ch <- 42 //fatal error: all goroutines are asleep - deadlock!  goroutine 1 [chan send]:

			fmt.Println(<-ch)

		//

	*/

	//WORKING

	ch <- 42
	//ch <- 43

	fmt.Println(len(ch), cap(ch))

	fmt.Println(<-ch)

	var ch1 chan int
	ch1 <- 1 // blocks forever
	<-ch1    // blocks forever
}
