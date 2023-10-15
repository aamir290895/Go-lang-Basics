package main

import "fmt"

func pointers() {

	a := 10

	increment(&a)

	b := &a

	fmt.Println(*b)

}

func increment(a *int) {

	for *a < 100 {
		*a++

	}
}
