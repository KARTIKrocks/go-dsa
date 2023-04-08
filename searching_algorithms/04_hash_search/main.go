// Golang program for implementation of Hash Search

/*In Go, hash tables are implemented as maps. A map is an unordered collection of key-value pairs where each key is unique.
 *To perform a hash search in Go, you can use the built-in map data type. Here is an example of using a map to perform a hash search:
 *The output shows result and the value of given key in map, returns (false, -1 ) if not present .
 */

package main

import "fmt"

func hashSearch(m map[string]int, key string) (bool, int) {

	if val, ok := m[key]; ok {
		return true, val
	}
	return false, -1
}

func main() {
	m := make(map[string]int)
	m["Alice"] = 27
	m["Bob"] = 35
	m["Charlie"] = 42

	fmt.Println(hashSearch(m, "Bolb"))
}
