package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-7, -5, -4, 3, 6, 8, 9}
	fmt.Printf("Squares of the array %v is %v\n", arr, sortedSquares(arr))
}

// Time Comlpexity - O(n) | Space Complexity - O(n)
// It seems like Two Pointers
func sortedSquares(arr []int) []int {
	squares := make([]int, len(arr))
	smallIdx := 0
	largeIdx := len(arr) - 1
	smallVal := arr[smallIdx]
	largeVal := arr[largeIdx]
	for i := largeIdx; i >= 0; i-- {
		smallVal = arr[smallIdx]
		largeVal = arr[largeIdx]
		if math.Abs(float64(smallVal)) > math.Abs(float64(largeVal)) {
			squares[i] = smallVal * smallVal
			smallIdx++
			continue
		}
		squares[i] = largeVal * largeVal
		largeIdx--
	}
	return squares
}
