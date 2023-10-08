package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 5, 4, 3, 2}

	// arr = bubbleSort(arr)

	arr = insertionSort(arr)

	str := reverseString("hello darling")
	fmt.Println(str)
	fmt.Println(arr)

}
