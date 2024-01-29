package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func ping(ch chan string) {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)

		ch <- "ping"

	}
}

func pong(ch chan string) {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- "pong"
	}
}

func pingPong() {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 1
		for {
			select {
			case val := <-ch:

				if val == 2 {
					fmt.Println("ping", val)
					ch <- 1
					time.Sleep(time.Millisecond * 1)

				}
			}
		}

	}(ch)

	go func(ch chan int) {

		for {
			select {
			case val := <-ch:

				if val == 1 {
					fmt.Println("pong", val)
					ch <- 2
					time.Sleep(time.Millisecond * 1)

				}

			}

		}

	}(ch)

	time.Sleep(time.Millisecond * 10)

}
