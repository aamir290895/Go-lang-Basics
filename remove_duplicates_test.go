package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	arr := []int{0, 1, 1, 1, 2, 2, 2, 3, 3, 3}

	arr = removeDuplicates(arr)
	t.Fatal(arr)

}

func TestRemoveDuplicatesAnotherWay(t *testing.T) {

	arr := []int{0, 1, 1, 1, 2, 2, 2, 3, 3, 3}

	arr = removeDuplicates(arr)
	fmt.Println(arr)

}

func TestRemoveDuplicateFromString(t *testing.T) {

	str := "hello world"

	str = removeDuplicateFromString(str)

	fmt.Println(str)
}
