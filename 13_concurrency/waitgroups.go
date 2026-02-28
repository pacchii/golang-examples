package _3_concurrency

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func WgExample() {

	fmt.Println("In WgExample func")
	wg.Add(1)

	wg.Go(foo)
	//or
	//go foo()

	for i := 0; i < 5; i++ {
		fmt.Println("Main", i)
	}

	wg.Wait()
	fmt.Println("Main Method ends")

}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("FOOOOOOOOOOOOOOOOOOOOOO", i)
		time.Sleep(2 * time.Second)
	}

	wg.Done()
}
