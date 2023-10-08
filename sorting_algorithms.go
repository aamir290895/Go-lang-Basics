package main

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func insertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Move elements of arr[0..i-1] that are greater than key
		// to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr

}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left, right := []int{}, []int{}
	for _, element := range arr {
		if element < pivot {
			left = append(left, element)
		} else if element > pivot {
			right = append(right, element)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}
