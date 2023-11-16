package sorts

func chooseSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		minindex := i
		for j := i+1; j < len(nums); j++ {
			if nums[j] < nums[minindex] {
				minindex = j
			}
		}
		nums[i], nums[minindex] = nums[minindex], nums[i]
	}
}