package main

import "fmt"

func SelectionSort() {

	arr := []int{25, 5, 19, 45, 3, 57}
	n := len(arr)

	for i := 0; i < n-1; i++ {
		pos := i // Initialize pos with the current index

		// Find the index of the minimum element in the unsorted part of the array
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[pos] {
				pos = j
			}
		}

		// Swap the found minimum element with the first element
		arr[i], arr[pos] = arr[pos], arr[i]
	}

	fmt.Println(arr)
}
