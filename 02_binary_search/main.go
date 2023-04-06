// Golang program for implementation of Binary Search

/*
 *A binary search is a search strategy used to find elements within a list by consistently reducing the amount of data to be searched and thereby increasing the rate at which the search term is found. To use a binary search algorithm,
 *the list to be operated on must have already been sorted.
 */

package main

import "fmt"

func binarySearch(key int, array []int) (bool, int) {

	low := 0               // first index
	high := len(array) - 1 // last index

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

	return false, -1
}

func main() {
	items := []int{1, 2, 3, 4, 5}
	fmt.Println(binarySearch(7, items))
}
