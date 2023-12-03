package main

import (
	"sync"
	"time"
)

var mu sync.Mutex

func ping(ch chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- "ping"
		time.Sleep(time.Second)

		<-ch

	}
}

func pong(ch chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 10; i++ {
		<-ch
		time.Sleep(time.Second)

		ch <- "pong"
	}
}
