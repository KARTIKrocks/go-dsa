// Golang program for implementation of Selection Sort

/*
This sorting algorithm begins by finding the smallest element in an array and interchanging it with data at, for instance, array index [0].
Starting at index 0, we search for the smallest item in the list that exists between index 1 and the index of the last element.
When this element has been found, it is exchanged with the data found at index 0. We simply repeat this process until the list becomes sorted.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(array []int) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func main() {
	slice := generateSlice(10)
	fmt.Println("--- Unsorted --- \n", slice)
	selectionSort(slice)
	fmt.Println("--- Sorted ---\n", slice)
}
