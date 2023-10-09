package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 5, 4, 3, 2}

	// arr = bubbleSort(arr)

	arr = quickSort(arr)

	str := reverseString("hello darling")

	anagon := isAnagram("hello", "lloeh")
	fmt.Println(str)
	fmt.Println(anagon)

}
