package problems

func TwoSum(nums []int, target int) []int {
	nums_map := make(map[int]int)
	for i, v := range nums {
		if j, ok := nums_map[target-v]; ok {
			return []int{i, j}
		}
		nums_map[v] = i
	}
	return []int{}

}
