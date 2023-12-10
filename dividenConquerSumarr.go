package main

import "fmt"

func sumArr(arr []int, low int, high int) int {
	if low == high {
		return arr[low]
	} else {
		mid := (low + high) / 2
		leftSum := sumArr(arr, low, mid)
		rightSum := sumArr(arr, mid+1, high)
		return leftSum + rightSum

	}

}

func main() {
	// Example usage
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:", arr)

	res := sumArr(arr, 0, len(arr)-1)

	fmt.Println("Sum array:", res)
}
