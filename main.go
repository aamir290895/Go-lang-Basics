package main

import (
	"go-lang-basics/ginframework"
)

func main() {

	// ch := make(chan string)

	// go ping(ch)
	//labstack.Setup()
	ginframework.Setup()
	// go pong(ch)

	// for {
	// 	select {
	// 	case val, ok := <-ch:
	// 		if ok {
	// 			fmt.Println(val)

	// 		}
	// 	default:
	// 		fmt.Println("no signal")
	// 	}
	// }

}
