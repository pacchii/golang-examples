package _3_concurrency

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func AtomicExample() {
	var counter int64

	const n = 10000

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			//counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
