package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

//var sem chan struct{} // Semaphore using a channel

func sayHello(ch chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		ch <- "hello chan"

		time.Sleep(time.Millisecond * 10)
	}

	close(ch)

	wg.Done()

}

func receiveHello(ch chan string) {

	for <-ch == "hello chan" {

		fmt.Println(<-ch + " received")

	}

	wg.Done()

}

func goroutines(t *testing.T) {

	ch := make(chan string)
	wg.Add(2)

	t.Log("Main function starts")

	go sayHello(ch)

	go receiveHello(ch)

	select {
	case val := <-ch:
		t.Log(val)
	}

	wg.Wait()
	//time.Sleep(time.Second)
}

func incrementByOne(a *int) {
	// mu.Lock()
	// defer mu.Unlock()

	// sem <- struct{}{} // Acquire semaphore

	for i := 0; i < 100; i++ {
		*a++

	}

	// <-sem // Release semaphore

	wg.Done()
}

func decrementByOne(a *int) {
	// mu.Lock()
	// defer mu.Unlock()

	// sem <- struct{}{} // Acquire semaphore

	for i := 0; i < 10; i++ {
		*a--

	}

	wg.Done()

	// sem <- struct{}{} // Acquire semaphore

}

func raceConditionManagement(t *testing.T) {
	var x int

	wg.Add(2)
	go incrementByOne(&x)
	go decrementByOne(&x)

	wg.Wait()

	// Log the final value of x after both goroutines have completed
	t.Log("Final Value of x:", x)
}
