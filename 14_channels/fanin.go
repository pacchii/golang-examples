package _4_channels

import (
	"fmt"
	"sync"
)

func FanINExample() {

	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go sendData(even, odd)

	go receiveData(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Done!")
}

func receiveData(e, o <-chan int, fanin chan<- int) {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range e {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

func sendData(e, o chan int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}
