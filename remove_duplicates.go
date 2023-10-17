package main

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
