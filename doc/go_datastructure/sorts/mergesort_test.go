package sorts

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{100, 23, 56, 213, 2, 3, 9, -1}
	fmt.Println("排序前:", nums)
	mergeSort(nums, 0, len(nums)-1)
	fmt.Println("排序后:", nums)
}
