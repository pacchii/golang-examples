package _3_concurrency

import (
	"fmt"
	"runtime"
)

func RaceCondition() {

	fmt.Println("No of CPUs", runtime.NumCPU())
	fmt.Println("No of Goroutines", runtime.NumGoroutine())
	counter := 0

	const gs = 100000

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			runtime.Gosched()
			//time.Sleep(100 * time.Millisecond)
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("No of Goroutines", runtime.NumGoroutine())
	fmt.Println(counter)
}
