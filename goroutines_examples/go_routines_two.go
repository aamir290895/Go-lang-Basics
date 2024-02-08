// You can edit this code!
// Click here and start typing.
package goroutinesexamples

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	ch1 := make(chan int)

	go func() {
		for i := 0; i <= 20; i++ {
			if i%2 == 0 {
				ch <- i
			} else {
				ch1 <- i
			}
		}
	}()

	go even(ch)
	go odd(ch1)

	time.Sleep(time.Millisecond * 100)

}

func even(ch chan int) {

	for value := range ch {

		fmt.Println("even", value)

	}
	close(ch)

}

func odd(ch chan int) {

	for value := range ch {
		fmt.Println("odd", value)

	}
	close(ch)

}
