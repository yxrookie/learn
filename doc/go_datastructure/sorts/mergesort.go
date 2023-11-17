package sorts

func mergeSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	Merge(nums, l, mid, r)
}

func Merge(nums []int, l int, mid int, r int) {
	tep := make([]int, r-l+1)
	//index := 0
	index, i, j := 0, l, mid+1
	for i < mid+1 && j < r+1 {
		if nums[i] < nums[j] {
			tep[index] = nums[i]
			i++
		} else {
			tep[index] = nums[j]
			j++
		}
		index++
	}
	for i < mid+1 {
		tep[index] = nums[i]
		index++
		i++
	}
	for j < r+1 {
		tep[index] = nums[j]
		index++
		j++
	}

	for i = 0; i < r-l+1; i++ {
		nums[l+i] = tep[i]
	}
}
