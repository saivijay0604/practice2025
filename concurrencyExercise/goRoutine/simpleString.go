package concurrencyexercise

import (
	"fmt"
	"time"
)

func simpleString(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func GoRoutines() {
	// simple function call.
	simpleString("simple function cal")

	//goroutine function cal
	go simpleString("gorutine functiona cal - 1")

	// goroutine with anonymous function call
	go func() {
		simpleString("anonymous functiona cal - 2")
	}()

	//goroutine with function value call
	fv := simpleString
	go fv("function value cal - 3")

	//wait for the goroutine
	time.Sleep(100 * time.Millisecond)

	fmt.Println("done!")

}
