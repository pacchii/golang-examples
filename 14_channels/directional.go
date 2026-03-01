package _4_channels

import (
	"fmt"
	"time"
)

// In this, not mandatorely both run in different goroutine.
// Here for receive-onyly func i have not added go, so it will run with main goroutine.
// So we will have 2 goroutines (g1, main) thats enough for channels to communicate
// NOTE : But if i reverse the go keyword (add to send function instead of receive then deadlock will occur.)
// And also if case :  main receive then go send() this also deadlock, because receive is blocked but no sender ready
func DirectionalExample() {

	ch := make(chan int)

	//send-only type channel
	go func(ch chan<- int) {
		//fmt.Println(<- ch) //Compile error : Invalid operation: <- ch (receive from the send-only type chan<- int)

		ch <- 1
	}(ch)

	time.Sleep(5 * time.Second)

	//receive-only type channel
	func(ch <-chan int) {
		//ch <- 1  //Compie error : Invalid operation: ch <- 1 (send to the receive-only type <-chan int)
		fmt.Println(<-ch)
	}(ch)

	time.Sleep(1 * time.Second)
}
