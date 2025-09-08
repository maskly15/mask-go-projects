package problems

import "fmt"

func isDuplicate(checkString string) bool {
	dctWord := make(map[byte]int)

	for i := 0; i < len(checkString); i++ {
		if _, ok := dctWord[checkString[i]]; ok {
			return false
		} else {
			dctWord[checkString[i]] = i
		}
	}
	return true

}

func LengthOfLongestSubstring(s string) int {
	// start := 0
	// end := 0
	if len(s) == 0 {
		return 0
	}
	for i := len(s); i > 0; i-- {
		for j := 0; i+j <= len(s); j++ {
			// fmt.Println(s[j:i+j])
			if isDuplicate(s[j : i+j]) {
				// fmt.Println("Longest String ",text[j:i+j])
				return i
			}
		}
	}
	return 1

}

func LengthOfLongestSubstringSlideWindow(s string) int {
	var res int
	if len(s) <= 1 {
		return len(s)
	}
	left := 0
	hm := make(map[byte]int)

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	// "abcabcbb"
	for right := 0; right < len(s); right++ {
		fmt.Println(s[left : right+1])
		fmt.Println("left", left)
		fmt.Println("right", right)
		fmt.Println("=====")

		if v, ok := hm[s[right]]; ok && v >= left {

			left = v + 1
		}
		hm[s[right]] = right

		res = max(res, right-left+1)

	}
	return res
}
