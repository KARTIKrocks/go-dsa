// Golang program for implementation of Bubble Sort

/* 
Given an unordered list, we compare adjacent elements in the list,
each time, putting in the right order of magnitude, only two elements. The algorithm hinges on a swap procedure. 
Knowing how many times to swap is important when implementing a bubble sort algorithm. To sort a list of numbers such as [3, 2, 1], we need to swap the elements a maximum of twice. 
This is equal to the length of the list minus 1
 */ 

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// swap elements
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
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
	bubbleSort(slice)
	fmt.Println("--- Sorted ---\n", slice)
}
