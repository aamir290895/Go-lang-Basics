package main

import "fmt"

func closures() {
	// Outer function defines a variable
	outerVar := 10

	// Inner function (closure) captures the outer variable
	closure := func() {
		// The inner function can access the outer variable
		fmt.Printf("Outer variable: %d\n", outerVar)
	}

	// Call the closure
	closure()

	// The outer variable is still accessible even though it's outside its scope
	fmt.Printf("Outer variable outside closure: %d\n", outerVar)
}
