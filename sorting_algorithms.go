package main

// Bubble Sort:

// Definition: Bubble sort is a simple comparison-based sorting algorithm that repeatedly steps through the list to be sorted, compares each pair of adjacent items, and swaps them if they are in the wrong order.
// Key Difference: It repeatedly compares adjacent elements and swaps them if they are in the wrong order, causing the largest unsorted element to "bubble up" to its correct position. It has poor performance for large datasets.
// Selection Sort:

// Definition: Selection sort is a simple comparison-based sorting algorithm that divides the input array into a sorted and an unsorted region. It repeatedly selects the minimum (or maximum) element from the unsorted region and places it at the end of the sorted region.
// Key Difference: It scans the entire unsorted region to find the minimum (or maximum) element and moves it to the sorted region. It performs poorly on large datasets.
// Insertion Sort:

// Definition: Insertion sort is a simple comparison-based sorting algorithm that builds the final sorted array one item at a time. It takes each element from the unsorted portion and inserts it into its correct position in the sorted portion.
// Key Difference: It works well for small datasets or when the input data is mostly sorted, as it has a time complexity of O(n^2) in the worst case.
// Quick Sort:

// Definition: Quick sort is a fast, efficient, and widely used comparison-based sorting algorithm. It uses a divide-and-conquer strategy to partition the array into two subarrays, recursively sorts them, and combines them to produce a sorted array.
// Key Difference: It efficiently handles large datasets and has an average time complexity of O(n log n). It selects a pivot element and partitions the array based on the pivot's value.
// Merge Sort:

// Definition: Merge sort is an efficient, stable, and comparison-based sorting algorithm that divides the array into two halves, sorts each half, and then merges the two sorted halves into one.
// Key Difference: It is stable, meaning it preserves the relative order of equal elements. It has a consistent time complexity of O(n log n) but requires additional memory for the merging step.
// Heap Sort:

// Definition: Heap sort is an in-place, efficient, and comparison-based sorting algorithm. It builds a binary heap from the input array and repeatedly extracts the maximum element (for ascending order) and places it at the end of the array.
// Key Difference: It has a time complexity of O(n log n) and doesn't require additional memory for sorting, making it suitable for large datasets. It is not stable.

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

// selectionSort sorts an array using the selection sort algorithm by repeatedly finding the minimum element and swapping it with the current element.

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// performs a merge sort on the array using a recursive divide-and-conquer approach. The merge function combines two sorted subarrays into a single sorted array.

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
