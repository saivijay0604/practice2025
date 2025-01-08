package concurrencyexercise

import (
	"fmt"
	"sync"
)

func Waitgroup() {
	fmt.Println("Wait Group:")
	var data int
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		data++
	}()

	wg.Wait()

	fmt.Println("Data: ", data)

}
