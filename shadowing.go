package main

import "fmt"

func shadowing() {
	// Outer scope
	x := 10
	fmt.Println("Outer x:", x) // Prints "Outer x: 10"

	// Inner scope
	{
		// Shadowing the outer variable 'x' with a new variable 'x'
		x := 20
		fmt.Println("Inner x:", x) // Prints "Inner x: 20"
	}

	// The outer 'x' is still accessible outside the inner scope
	fmt.Println("Outer x after inner scope:", x) // Prints "Outer x after inner scope: 10"
}
