package main

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {

	str := "abcdeyf"
	str = reverseString(str)

	if str == "fedcba" {
		fmt.Println("test_passed")
	} else {
		t.Errorf("expected %s", "fedcab")

	}

	if str != "ab" {
		t.Fail()
	}

	t.Fatal(str)

}

func TestCountVowelsInString(t *testing.T) {

	str := "aaAuoperIot"

	m := countVowelsInString(str)

	fmt.Println(m)
}
