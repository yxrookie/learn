package bsearch

import (
	"fmt"
	"testing"
)

func TestBsearch(t *testing.T) {
	nums := []int{-1,2,5,67,790}
    fmt.Println(bsearch(nums, 0, len(nums)-1, -1))
	fmt.Println(bsearch(nums, 0, len(nums)-1,  2))
	fmt.Println(bsearch(nums, 0, len(nums)-1, 790))
	fmt.Println(bsearch(nums, 0, len(nums)-1, 67))
	fmt.Println(bsearch(nums, 0, len(nums)-1, 100))
}