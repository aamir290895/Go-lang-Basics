package main

import "fmt"

func pointers() {

	a := 10

	wg.Add(1)
	go increment(&a) // &a represents the memory 0xc000016088 for variable a

	wg.Wait()

	b := &a

	fmt.Println(*b) // *b represents the value of 0xc000016088 this memory location

}

func increment(a *int) {

	for *a < 100 {
		*a++

	}

	wg.Done()
}
