package main

import "testing"

func TestIsPrime(t *testing.T) {

	n := []int{}

	for i := 0; i < 100; i++ {
		if isPrime(i) {
			n = append(n, i)
		}

	}

	t.Fatal(n)

}
