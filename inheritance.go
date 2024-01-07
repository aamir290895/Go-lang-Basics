package main

import "fmt"

// Animal represents a generic animal
type Animal struct {
	Name string
}

// Speak defines a method for animals to make a sound
func (a *Animal) Speak() {
	fmt.Println("Generic animal sound")
}

// Dog represents a specific type of animal with additional behavior
type Dog struct {
	// Embedding the Animal type
	Animal
	Breed string
}

// Speak overrides the Speak method for a Dog
func (d *Dog) Speak() {
	fmt.Println("Woof!")
}

func testInheritance() {
	// Create a Dog instance
	myDog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Labrador",
	}

	// Access fields from embedded type
	fmt.Println("Name:", myDog.Name)
	fmt.Println("Breed:", myDog.Breed)

	// Call overridden method
	myDog.Speak() // Output: Woof!

	// Call method from the embedded type
	myDog.Animal.Speak() // Output: Generic animal sound
}
