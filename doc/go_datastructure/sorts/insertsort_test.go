package sorts

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	nums := []int{100,23,56,213,2,3,9,-1}
	fmt.Println("排序前:", nums)
	insertSort(nums)
	fmt.Println("排序后:", nums)
}