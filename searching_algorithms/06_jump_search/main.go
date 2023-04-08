// Golang program for implementation of Jump Search

/*Jump search is a searching algorithm that works on sorted arrays.
 *It involves jumping ahead by a fixed number of steps (the jump size) instead of scanning every element.
 *Once a block is found that contains the search value, a linear search is performed to find the value within that block.
 *Here is an example to perform a jump search:
 *The output shows result and the value of given key, returns (false, -1 ) if not present.
 */

package main

import (
	"fmt"
	"math"
)

func jumpSearch(array []int, key int) (bool, int) {
	n := len(array)
	step := int(math.Sqrt(float64(n)))

	prev := 0
	for array[min(step, n)-1] < key {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return false, -1
		}
	}

	// Linear search in the block starting from prev
	for array[prev] < key {
		prev++
		if prev == min(step, n) {
			return false, -1
		}
	}

	if array[prev] == key {
		return true, prev
	}
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
	fmt.Println(jumpSearch(items, -1))
}
