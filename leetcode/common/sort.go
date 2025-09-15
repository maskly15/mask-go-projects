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

func quickSort(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}
	pivot := nums[int(len(nums)/2)]
	left := []int{}
	right := []int{}
	middle := []int{}

	for i := range len(nums) {
		if nums[i] < pivot {
			left = append(left, nums[i])
		} else if nums[i] == pivot {
			middle = append(middle, pivot)
		} else {
			right = append(right, nums[i])
		}
	}

	return append(append(quickSort(left), middle...), quickSort(right)...)

}
