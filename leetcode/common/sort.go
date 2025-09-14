package common

func bubbleSort(nums []int) []int {
	var temp int
	var isSwap bool
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				isSwap = true
				temp = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}

		}
		if isSwap {
			isSwap = false
		} else {
			break
		}
	}

	return nums

}
