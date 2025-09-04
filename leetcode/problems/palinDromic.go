package problems

// 1. Init an matrix n*n with false value, the ans
// 2. Init the True Value for dep[i][i]
// 3. Loop all s[i] & s[i+1] if s[i] == s[i+1] => Set value to true
// 4. Iterations again with Diff , compare if s[i] == s[i] && dep[i+1][j-1] == true => dep[i][j] = True  , Update i-j if > ans
func LongestPalindrome(s string) string {
	n := len(s)
	var ans = []int{0, 0}
	dep := make([][]bool, n)

	for i := range n {
		dep[i] = make([]bool, n)
		dep[i][i] = true

	}

	for i := range (n) - 1 {

		if s[i] == s[i+1] {
			dep[i][i+1] = true
			ans = []int{i, i + 1}
		}
	}

	for diff := 2; diff < n; diff++ {

		for i := 0; i < n-diff; i++ {
			j := i + diff

			if s[i] == s[j] && dep[i+1][j-1] == true {
				dep[i][j] = true
				ans = []int{i, j}
			}
		}

	}

	return s[ans[0] : ans[1]+1]

}
