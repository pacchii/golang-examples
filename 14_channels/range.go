package _4_channels

import "fmt"

// If you dont close channel after send completed, then receiver will be still waiting for data expecting sender will send.
// If you close channel then it gets message that channel closed
func RangeExample() {

	ch := make(chan int)

	go foo(ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Program Exited")

}

func foo(c chan<- int) {

	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) //If you dont close then deadlock goroutine 1 [chan receive]:
}
