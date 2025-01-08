package concurrencyexercise

import "fmt"

func UnBufferedChannel() {
	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		ch <- c
	}(2, 3)
	result := <-ch
	fmt.Println("Computation value: ", result)
}

func BufferedChannel() {
	ch := make(chan int, 6)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("sending values into channel: ", i)
			ch <- i
		}
		defer close(ch)
	}()
	for v := range ch {
		fmt.Println("revecing values from channel: ", v)
	}
}

// Channael Direction - Relaying f message with channel direction

func ChannelDirection() {
	var ch1 = make(chan string)
	var ch2 = make(chan string)
	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	r := <-ch2
	fmt.Println("recieve message on channel 2: ", r)
}

// send message on ch1
func genMsg(ch1 chan<- string) {
	ch1 <- "message1"
}

// recieve message on ch1
// send it on ch2
func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	msg := <-ch1
	ch2 <- msg
}

// channel ownership
// owner : create, return channel and spin the func then write the date and close it.
// consumer: relay the channel though  then iterate though channel and print value
func ChannelOwnership() {
	owner := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()
		return ch
	}
	consumer := func(ch <-chan int) {
		for v := range ch {
			fmt.Println("RECEIVED: ", v)
		}
		fmt.Println("done!")
	}
	ch := owner()
	consumer(ch)
}
