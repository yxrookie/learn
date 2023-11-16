package sorts


func insertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		nowVal := nums[i]
		j := i-1
		for ; j >= 0; j-- {
			if nowVal < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
        nums[j+1] = nowVal 
		//fmt.Println(nums)
	}
}
