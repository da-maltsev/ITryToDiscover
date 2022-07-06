package main

import (
	"ITryToDiscover/sorts"
	"fmt"
)

func main() {

	nums := []int{-1, 4, 3, 6, 1, 11, 0, 4, 9, 2}

	fmt.Println("Original slice:", nums)
	fmt.Println("BubbleSorted slice:", sorts.BubbleSort(nums))
	fmt.Println("InsertSorted slice:", sorts.InsertSort(nums))
	fmt.Println("QuickSorted slice:", sorts.QuickSort(nums))
	fmt.Println("Original slice once again:", nums)

}
