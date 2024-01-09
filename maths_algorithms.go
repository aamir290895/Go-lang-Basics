package main

import "math"

func isPrime(n int) bool {

	if n <= 1 {
		return false
	}

	if n == 2 || n == 5 {
		return true
	}

	if n%2 == 0 || n%5 == 0 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))

	for i := 3; i < sqrt; i = i + 2 {

		if n%i == 0 {
			return false
		}

	}

	return true

}
