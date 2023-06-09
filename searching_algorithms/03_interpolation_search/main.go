// Golang program for implementation of Interpolation Search

/*The Interpolation Search is an improvement over Binary Search for instances,
 *where the values in a sorted array are uniformly distributed.
 *Binary Search always goes to middle element to check. On the other hand interpolation search may go to different locations according the value of key being searched.
 *Here is the source code of the Go program to search element in an integer array using Interpolation search algorithm. The output shows result and the position of element in array returns (false , -1) if not present.
 */

package main

import "fmt"

func interpolationSearch(array []int, key int) (bool, int) {

	min, max := array[0], array[len(array)-1]

	low, high := 0, len(array)-1

	for {
		if (key < min || key > max) {
			return false, -1
		}

		// make a guess of the location
		var guess int
		if high == low {
			if array[low] == key {
				return true, low
			}
			return false, -1
		} else {
			size := high - low
			guess = low + int(float64(size-1) * (float64(key-min) / float64(max-min)))
		}

		// maybe we found it?
		if array[guess] == key {
			// scan backwards for start of value range, finds first occurence of the key
			for guess > 0 && array[guess-1] == key {
				guess--
			}
			return true, guess
		}

		// if we guessed to high, guess lower or vice versa
		if array[guess] > key {
			high = guess - 1
			max = array[high]
		} else {
			low = guess + 1
			min = array[low]
		}
	}
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(interpolationSearch(items, 45))
}
