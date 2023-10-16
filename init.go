package main

var message string

// In Go, the init function is a special function that can be used for package-level initialization. When a Go program or package is built, the init function is automatically executed before the main function or any other functions in the package.

func init() {
	message = "Initialization complete"
}

// func main() {
// 	fmt.Println(message)
// }
