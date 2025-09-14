package problems

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

func threeSum(nums []int) [][]int {

	nums = bubbleSort(nums)
	var left, right int
	var total int
	var results [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left = i + 1
		right = len(nums) - 1

		for left < right {
			total = nums[i] + nums[left] + nums[right]
			if total == 0 {
				// append

				results = append(results, []int{nums[i], nums[left], nums[right]})

				// avoid dup
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// avoid dup
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if total < 0 {
				left++
			} else {
				right--
			}

		}
	}
	return results

}
