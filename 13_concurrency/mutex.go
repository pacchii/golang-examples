package _3_concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

//To solve the race condition

func MutexExample() {

	counter := 0

	const n = 100000

	var m sync.Mutex

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			m.Lock()
			runtime.Gosched()
			counter++
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
