package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sayHello(ch chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		ch <- "hello chan"

		time.Sleep(time.Millisecond * 100)
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

// func main() {

// 	ch := make(chan string)
// 	wg.Add(2)

// 	fmt.Println("Main function starts")

// 	go sayHello(ch)

// 	go receiveHello(ch)

// 	select {
// 	case val := <-ch:
// 		fmt.Println(val)
// 	}

// 	wg.Wait()
// 	//time.Sleep(time.Second)
// 	fmt.Println("Main function")
// }
