package main

import "fmt"

func main() {
	arr := []int{10, 84, 8, 48, 89, -1, -4, 5, 8, 567, 89, 894, 45}
	subsequence := []int{84, -4, 5, 8, 45}
	fmt.Println(traverseSimultaneously(arr, subsequence))
}

func traverseSimultaneously(arr, subseq []int) bool {
	subseqIdx := 0
	fmt.Printf("Subsequence items: %v\n", subseq)
	for _, val := range arr {
		if subseqIdx == len(subseq) {
			break
		}
		if val == subseq[subseqIdx] {
			subseqIdx++
			fmt.Printf("Value: %v\n", val)
			fmt.Printf("Ramaining items: %v\n", subseq[subseqIdx:])
		}
	}
	fmt.Printf("Idx: %d and Len: %d\n", subseqIdx, len(subseq))
	return subseqIdx == len(subseq)
}
