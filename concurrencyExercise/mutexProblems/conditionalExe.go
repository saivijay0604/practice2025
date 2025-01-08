package concurrencyexercise

import (
	"fmt"
	"sync"
)

//condition variable is one of the synchronization mechanisms
// A condition variable is basically a container of Goroutines that are waiting for a certain condition

//Three functions: Wait() ; signal(); Boardcast()

// In ConditionSignal() function we used conditioal variable to coordinate the producer and consumer goroutine

var sharedRecord = make(map[string]interface{})

func ConditionSignal() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	c := sync.NewCond(&mu)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()

		for len(sharedRecord) == 0 {
			//time.Sleep(1 * time.Millisecond)  Instad of sleep use wait method
			c.Wait()
		}
		fmt.Println("Mutex Condition: share data- ", sharedRecord["record0"])
		c.L.Unlock()
	}()
	c.L.Lock()
	sharedRecord["record1"] = "zero Record"
	c.Signal()
	c.L.Unlock()

	wg.Wait()
}

func ConditionalBoardcast() {
	var wg sync.WaitGroup

	var mu sync.Mutex
	c := sync.NewCond(&mu)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRecord) < 1 {
			c.Wait()
		}

		fmt.Println(sharedRecord["record1"])
		c.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRecord) < 2 {
			c.Wait()
		}
		fmt.Println(sharedRecord["record2"])
		c.L.Unlock()
	}()
	c.L.Lock()
	c.Broadcast() // We need to wake up both of our go routines, so we use Boardcast method.

	sharedRecord["record1"] = "Frirst Record"
	sharedRecord["record2"] = "Second Record"
	c.L.Unlock()
	wg.Wait()

}
