package _4_channels

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func FanOutExample() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	//go FanOut(c1, c2)
	go FanOutThrottle(c1, c2)
	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("Done!")
}

func populate(c chan int) {
	for i := 0; i < 1000; i++ {
		c <- i
	}
	close(c)
}

func FanOut(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c2 <- task(v)
		}()
	}
	fmt.Println("Goroutines", runtime.NumGoroutine())
	wg.Wait()
	close(c2)
}

func FanOutThrottle(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- task(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)

}

func task(v int) int {
	time.Sleep(time.Duration(v) * 500 * time.Millisecond)
	return v * 100
}
