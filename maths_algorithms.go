package main

import (
	"fmt"
	"math"
)

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

func hcf(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / hcf(a, b)
}

func generatePascalsTriangle(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1

		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	fmt.Println(result)
	return result
}
