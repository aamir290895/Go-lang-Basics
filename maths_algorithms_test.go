package main

import "testing"

func TestIsPrime(t *testing.T) {

	test := []struct {
		num      int
		expected bool
	}{

		{2, true},
		{5, true},
		{6, false},
	}

	for _, v := range test {
		if isPrime(v.num) {
			t.Logf("yes the num %d is prime", v.num)
		} else {
			t.Logf("No, the num %d is not prime", v.num)
		}

	}

}

func TestHcf(t *testing.T) {

	t.Log(" the start of the test.")

	tests := []struct {
		a, b, expected int
	}{
		{12, 18, 6},
		{25, 35, 5},
		{8, 12, 4},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := hcf(test.a, test.b)
		if result != test.expected {
			t.Errorf("GCD of %d and %d, expected: %d, got: %d", test.a, test.b, test.expected, result)
		}

		t.Log(result)
	}

}
