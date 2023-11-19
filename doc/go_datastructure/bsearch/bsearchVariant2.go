package bsearch

// 查找最后一个值等于给定值的元素
func bsearchVariant2 (nums []int, low, high, target int) int{
	for low <= high {
		mid := low + (high-low) >> 1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}