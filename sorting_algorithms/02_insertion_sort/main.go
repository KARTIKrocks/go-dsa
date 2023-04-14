// Golang program for implementation of Insertion Sort

/* 
The idea of swapping adjacent elements to sort a list of items can also be used to implement the insertion sort.
In the insertion sort algorithm, we assume that a certain portion of the list has already been sorted, while the other portion remains unsorted.
With this assumption, we move through the unsorted portion of the list, picking one element at a time.
With this element, we go through the sorted portion of the list and insert it in the right order so that the sorted portion of the list remains sorted.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key
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
	insertionSort(slice)
	fmt.Println("--- Sorted ---\n", slice)
}
