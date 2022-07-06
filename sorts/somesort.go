package sorts

func BubbleSort(nums []int) []int {
	ok := true
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	for ok {
		ok = false
		for i := 0; i < len(sortedNums)-1; i++ {
			if sortedNums[i] > sortedNums[i+1] {
				sortedNums[i], sortedNums[i+1] = sortedNums[i+1], sortedNums[i]
				ok = true
			}
		}
	}
	return sortedNums
}

func InsertSort(nums []int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	for i := 1; i < len(sortedNums); i++ {
		itemToInsert := sortedNums[i]
		j := i - 1
		for j >= 0 && sortedNums[j] > itemToInsert {
			sortedNums[j+1] = sortedNums[j]
			j--
		}
		sortedNums[j+1] = itemToInsert
	}
	return sortedNums
}

func QuickSort(nums []int) []int {
	if l := len(nums); l < 2 {
		return nums
	}
	bigger := make([]int, 0)
	less := make([]int, 0)
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	item := sortedNums[0]

	for _, elem := range sortedNums[1:] {
		if elem > item {
			bigger = append(bigger, elem)
		} else {
			less = append(less, elem)
		}
	}
	sortedNums = append(QuickSort(less), item)
	sortedNums = append(sortedNums, QuickSort(bigger)...)
	return sortedNums
}
