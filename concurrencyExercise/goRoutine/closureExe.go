package concurrencyexercise

import (
	"fmt"
	"sync"
)

func ClosureExe() {
	fmt.Println("Closure Exercise:")

	var wg sync.WaitGroup

	inc := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Println("The value of i:", i)
		}()
		fmt.Println("return from function")
		return
	}

	inc(&wg)
	wg.Wait()
	fmt.Println("End of closure exercise")

}

func ClosureExe2() {
	fmt.Println("Closure Exercise 2:")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("The value of i: ", i)
		}(i)
	}
	wg.Wait()

}
