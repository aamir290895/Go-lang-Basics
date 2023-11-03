package main

import "fmt"

func removeDuplicates(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	uniqueIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}

	// The unique elements are in the slice nums[:uniqueIndex+1]
	return nums[:uniqueIndex+1]
}

func removeDuplicatesAnotherWay(nums []int) []int {

	aMap := make(map[int]int)

	arr2 := []int{}
	for _, v := range nums {
		aMap[v]++
	}

	for nums, counts := range aMap {
		fmt.Println(nums, ":", counts)
		arr2 = append(arr2, nums)

	}

	return arr2
}

func removeDuplicateFromString(str string) string {
	m := make(map[rune]bool)

	str2 := []rune{}

	for _, v := range str {
		if !m[v] {
			str2 = append(str2, v)
			m[v] = true
		}
	}

	return string(str2)
}
