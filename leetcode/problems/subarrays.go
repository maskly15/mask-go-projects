package problems

func FindSubArray(arr, subArr string) int {

	hm := make(map[byte][]int)
	for i := 0; i < len(arr); i++ {
		hm[arr[i]] = append(hm[arr[i]], i)
	}
	flag := true
	for _, start := range hm[subArr[0]] {
		if start+len(subArr) > len(arr) {
			return -1
		}
		for i := 0; i < len(subArr); i++ {
			if arr[start+i] != subArr[i] {
				flag = false
				break
			}
		}
		if flag {
			return start
		}

	}
	return -1

}
