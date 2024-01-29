package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)

	go ping(ch)

	go pong(ch)

	for {
		select {
		case val, ok := <-ch:
			if ok {
				fmt.Println(val)

			}
		default:
			fmt.Println("no signal")
		}
	}

}
