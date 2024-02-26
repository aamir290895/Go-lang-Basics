package goroutinesexamples

import (
	"fmt"
	"time"
)

func testExample() {

	ch := make(chan int)

	go one(ch)
	go two(ch)
	go three(ch)

	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Hello, 世界")

}

func one(ch chan int) {
	val := <-ch
	fmt.Println("val from go 1", val)

	ch <- val + 1
}

func two(ch chan int) {
	ch <- 1

	val := <-ch
	fmt.Println("val from go 2", val)

	ch <- val + 10
}

func three(ch chan int) {
	val := <-ch
	fmt.Println("val from go 3", val)

	ch <- val + 1

}
