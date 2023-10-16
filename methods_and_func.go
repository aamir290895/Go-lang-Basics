package main

type Calculator struct {
	x, y int
}

func add(x, y int) int {
	return x + y
}

func (c Calculator) add() int {
	return c.x + c.y
}

// func main() {
//     calc := Calculator{3, 4}
//     result := calc.add()

// result := add(3, 4)

//     fmt.Println(result)
// }

// In Go, both methods and functions are used to encapsulate blocks of code that can be executed. However, there are significant differences between methods and functions in terms of their behavior, where they are defined, and how they are called. Here are the key differences between methods and functions in Go:

// 1. Receiver:

// Function: Functions in Go do not have a receiver. They are standalone blocks of code that can be called from anywhere in your program without being associated with a specific data type.

// Method: Methods are associated with a specific data type, known as the receiver type. They operate on instances of that type and can access their data and properties.

// 2. Definition:

// Function: Functions are defined at the package level and can be called from anywhere within the package they are defined in.

// Method: Methods are defined as part of a type's definition. They are associated with a specific type and are usually defined in the same package as the type. Methods can be called on instances of the type.

// 3. Invocation:

// Function: Functions are called like regular functions in other programming languages, without a receiver. You simply provide the function name and its arguments.

// Method: Methods are called on instances of the receiver type using dot notation. You use the instance variable, followed by a dot, and then the method name. This is called a method call.

// 4. Scope:

// Function: Functions are generally global in scope within the package in which they are defined. They can be accessed from any part of the package.

// Method: Methods are only accessible within the package where the type and its methods are defined. They are not globally visible, and you cannot call methods on types from other packages unless the methods are exported.
