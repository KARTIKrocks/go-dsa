// Golang program for implementation of Binary Search

/*
 *A binary search is a search strategy used to find elements within a list by consistently reducing the amount of data to be searched and thereby increasing the rate at which the search term is found. To use a binary search algorithm,
 *the list to be operated on must have already been sorted.
 */

 package main

 import "fmt"
 
 func binarySearch(array []int, key int) (bool, int) {
 
	 low := 0               // first index
	 high := len(array) - 1 // last index
 
	 if key < array[low] || key > array[high] {
		 return false, -1
	 }
 
	 for low <= high {
		 median := (low + high) / 2
 
		 if array[median] == key {
			 // scan backwards for start of value range, finds first occurence of the key
			 guess := median
			 for array[guess-1] == key {
				 guess--
			 }
			 return true, guess
		 } else if array[median] < key {
			 low = median + 1
		 } else {
			 high = median - 1
		 }
	 }
 
	 return false, -1
 }
 
 func main() {
	 items := []int{2, 2, 2, 2, 4, 5}
	 fmt.Println(binarySearch(items, 2))
 }
 