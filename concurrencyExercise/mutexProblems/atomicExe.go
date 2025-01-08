package concurrencyexercise

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//Low level atomic ioerations in memory
// lockless operation
// Used for atomic Operations on Coounters

func AtomicPractice() {
	runtime.GOMAXPROCS(4)
	var count uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&count, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Atomic practice: ", count)

}
