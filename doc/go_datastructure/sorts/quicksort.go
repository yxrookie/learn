package sorts

func quickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	q := partition(nums, l ,r)
	quickSort(nums, l, q-1)
	quickSort(nums, q+1, r)
}

func partition(nums []int, l int, r int) int {
	//基础做法：每次的基准元素选择为第 r 个元素
	i, j := l, l
	for ; j < r+1; j=j+1 {
		if nums[j] < nums[r] {
			nums[i], nums[j] = nums[j], nums[i]
			i ++
		}
	}  
	nums[i], nums[r] = nums[r], nums[i]
	return i
}