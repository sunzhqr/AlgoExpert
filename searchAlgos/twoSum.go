package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 13
	fmt.Println(doubleLoops(arr, target))
	fmt.Println(hashTableForSearch(arr, target))
	fmt.Println(twoPointers(arr, target))
}

// First way use double for loops with Tine Complexity O(n^2) | Space Complexity O(1)
func doubleLoops(arr []int, target int) ([]int, bool) {
	firstNum, secondNum := 0, 0
	for i := 0; i < len(arr)-1; i++ {
		firstNum = arr[i]
		for j := i + 1; j < len(arr); j++ {
			secondNum = arr[j]
			if firstNum+secondNum == target {
				return []int{firstNum, secondNum}, true
			}
		}
	}
	return nil, false
}

// Second way use hash-table with Tine Complexity O(n) | Space Complexity O(n)
func hashTableForSearch(arr []int, target int) ([]int, bool) {
	hashTable := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if ok := hashTable[target-arr[i]]; ok {
			return []int{arr[i], target - arr[i]}, true
		}
		hashTable[arr[i]] = true
	}
	return nil, false
}

// Third way use Two Pointers only for sorted array with Tine Complexity O(n*log(n)) | Space Complexity O(1)
func twoPointers(arr []int, target int) ([]int, bool) {
	sort.Ints(arr) // It will be MergeSort or QuickSort
	left, right := 0, len(arr)-1
	currentSum := arr[left] + arr[right]
	for left < right {
		if currentSum == target {
			return []int{arr[left], arr[right]}, true
		} else if currentSum < target {
			left++
			currentSum = arr[left] + arr[right]
		} else {
			right--
			currentSum = arr[left] + arr[right]
		}
	}
	return nil, false
}
