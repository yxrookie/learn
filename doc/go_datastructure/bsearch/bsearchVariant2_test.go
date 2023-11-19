package bsearch

import (
	"fmt"
	"testing"
)

func TestBsearchVariant2(t *testing.T) {
	nums := []int{0, 0, 0, 0, 1, 1, 1, 23, 34, 546, 546, 546, 789}
	fmt.Println(bsearchVariant2(nums, 0 ,len(nums)-1, 0))
	fmt.Println(bsearchVariant2(nums, 0 ,len(nums)-1, 1))
	fmt.Println(bsearchVariant2(nums, 0 ,len(nums)-1, 546))
}