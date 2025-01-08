package concurrencyexercise

import (
	"fmt"
	"runtime"
	"sync"
)

func MutexPractice() {
	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amount int) {
		mu.Lock()
		balance = balance + amount
		mu.Unlock()
	}

	withdrawal := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance = balance - amount
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()

	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()

	}

	wg.Wait()
	fmt.Println("Balance : ", balance)
}
