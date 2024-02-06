package main

import "math"

func arrangeLowestNextToHighest(arr []int) []int {

	// arr := []int {1,2,25,8,7}

	s := 25
	h := 0
	hi := 0
	si := 0

	for i, v := range arr {

		if v > h {

			h = v
			hi = i
		}

		if s > v {
			s = v
			si = i

		}
	}

	// temp := []int{}

	// for _, v := range arr {

	// 	if v == h && v != s {
	// 		temp = append(temp, v)
	// 		temp = append(temp, s)

	// 	} else if v != s {
	// 		temp = append(temp, v)

	// 	}
	// }

	arr = append(arr[:si], arr[si+1:]...)

	temp := []int{}

	temp = append(temp, arr[s+1:]...)
	arr = append(arr[:hi], s)

	arr = append(arr, temp...)

	return arr
}

func findSecondLargest(arr []int) int {
	h := math.MinInt

	h2 := math.MinInt

	for _, v := range arr {

		if v > h {
			h2 = h
			h = v
		} else if v > h2 && h2 != h {
			h2 = v
		}
	}
	return h2

}
func findSecondNonRepeatingCharacter(str string) string {
	m := make(map[rune]int)

	for _, v := range str {
		m[v]++
	}

	var firstNonRepeating, secondNonRepeating rune

	for _, v := range str {
		if m[v] == 1 {
			if firstNonRepeating == 0 {
				firstNonRepeating = v
			} else {
				secondNonRepeating = v
				break
			}
		}
	}

	if secondNonRepeating != 0 {
		return string(secondNonRepeating)
	}

	return "" // Return an empty string if no second non-repeating character is found
}
