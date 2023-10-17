package main

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {

	str := "abcdef"
	str = reverseString(str)

	if str == "ffdcba" {
		fmt.Println("test_passed")
	} else {
		t.Errorf("expected %s", "fedcab")

	}

}
