package bsearch

// 查找第一个值大于或等于给定值的元素
func bsearchVariant3(nums []int, low, high, target int) int {
	for low <= high {
		mid := low + (high-low) >> 1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				high = mid-1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}