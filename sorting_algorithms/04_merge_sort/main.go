package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(array []int) []int {
	n := len(array)
	if n <= 1 {
		return array
	}
	mid := n / 2
	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
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
	slice = mergeSort(slice)
	fmt.Println("--- Sorted ---\n", slice)
}
