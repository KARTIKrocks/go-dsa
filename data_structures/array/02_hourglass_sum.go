package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 * 9 9 9
 *   9
 * 9 9 9
 */
func hourglassSum(arr [][]int32) int32 {
	rows := len(arr)
	cols := len(arr[0])
	maxSum := int32(-63) // Initialize with the lowest possible sum (-9 * 7 = -63)

	for i := 0; i < rows-2; i++ {
		for j := 0; j < cols-2; j++ {
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] +
				arr[i+1][j+1] +
				arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

func newArray(row, col int32) [][]int32 {
	arr := make([][]int32, row)
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator

	for i := int32(0); i < row; i++ {
		arr[i] = make([]int32, col)
		for j := int32(0); j < col; j++ {
			arr[i][j] = rand.Int31n(10) // Generate a random number between 0 and 9
		}
	}
	return arr
}

func main() {
	// arr := [][]int32{
	// 	{1, 1, 1, 0, 0, 0},
	// 	{0, 1, 0, 0, 0, 0},
	// 	{1, 1, 1, 0, 0, 0},
	// 	{0, 0, 2, 4, 4, 0},
	// 	{0, 0, 0, 2, 0, 0},
	// 	{0, 0, 1, 2, 4, 0},
	// }

	arr := newArray(7, 8)

	fmt.Println("array", arr)
	result := hourglassSum(arr)
	fmt.Println("Max hourglass sum:", result)
}
