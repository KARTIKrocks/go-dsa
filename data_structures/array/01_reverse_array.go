package main

import "fmt"

func reverseArray(arr []int) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", numbers)

	reverseArray(numbers)
	fmt.Println("Reversed array:", numbers)
}
