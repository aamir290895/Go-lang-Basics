package main

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {

	str := "abcdef"
	str = reverseString(str)

	if str == "fedcba" {
		fmt.Println("test_passed")
	} else {
		t.Errorf("expected %s", "fedcab")

	}

}

func TestCountVowelsInString(t *testing.T) {

	str := "aaAuoperIot"

	m := countVowelsInString(str)

	fmt.Println(m)
}
