// Golang program for implementation of Exponential Search

/*Exponential search is a searching algorithm that works on sorted arrays.
 *It involves two steps: first, we find a range where the search key may be present, and then we do a binary search in that range.
 *The range is found by repeatedly doubling the index i starting from 1 until arr[i] is greater than or equal to the search key.
 *Once the range is found, a binary search is performed in that range to find the key.
 *Here is an example to perform a exponential search:
 *The output shows result and the value of given key, returns (false, -1 ) if not present.
 */

package main

import "fmt"

func exponentialSearch(array []int, key int) (bool, int) {
	n := len(array)

	if key < array[0] || key > array[n-1] {
		return false, -1
	}
	if array[0] == key {
		return true, 0
	}

	// Find range for binary search by repeated doubling
	i := 1
	for i < n && array[i] <= key {
		i = i * 2
	}

	// Call binary search for the found range
	return binarySearch(array, i/2, min(i, n-1), key)
}

func binarySearch(array []int, low int, high int, key int) (bool, int) {

	if key < array[low] || key > array[high] {
		return false, -1
	}

	for low <= high {
		median := (low + high) / 2

		if array[median] == key {
			return true, median
		} else if array[median] < key {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	// recursive method 
	// if low <= high {
	// 	median := (low + high) / 2
	// 	if array[median] == key {
	// 		return true, median
	// 	}else if array[median] > key {
	// 		return binarySearch(array, low, median-1, key)
	// 	}
	// 	return binarySearch(array, median+1, high, key)
	// }

	return false, -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	items := []int{1, 2, 3, 4, 5}
	fmt.Println(exponentialSearch(items, 5))
}
