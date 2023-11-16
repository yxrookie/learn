package sorts

func bubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}