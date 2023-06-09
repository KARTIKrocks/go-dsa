// Golang program for implementation of Linear Search

// This technique pass over the list of elements, by using the index to move from the beginning of the list to the end. Each element is examined and if it does not match the search item, the next item is examined. By hopping from one item to its next, the list is passed over sequentially.
package main

import "fmt"

func linearsearch(array []int, key int) (bool, int) {
	for index, item := range array {
		if item == key {
			return true, index
		}
	}
	return false, -1
}

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println(linearsearch(items, 58))
}
