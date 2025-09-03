package problems

import "fmt"

// cách giải :
// 1. Bài toán đã cho 2 sorted arr, cách giải chay là merge 2 array lại
// 2. Dùng 2 con trỏ p1, p2 để so sánh 2 array, nếu thằng nào bé hơn cho vào array merge trc và P min + 1
// 3. Nếu Len lẻ thì lấy int (m+n / 2) , nếu Len chẵn thì lấy avarage( int(m+n/2), int(m+n/2) + 1 ))
//4. Có vẻ là cách nhanh nhất

func MergeSortedArray(nums1, nums2 []int) float64 {

	var p1, p2 int
	m := len(nums1)
	n := len(nums2)

	mergeArr := make([]int, m+n)
	for i := range m + n {
		if p1 >= m {
			mergeArr[i] = nums2[p2]
			p2++
			continue
		} else if p2 >= n {
			mergeArr[i] = nums1[p1]
			p1++
			continue
		} else {

			if nums1[p1] <= nums2[p2] {
				mergeArr[i] = nums1[p1]
				p1++
			} else {
				mergeArr[i] = nums2[p2]
				p2++
			}

		}
	}
	fmt.Println(mergeArr)
	fmt.Println((m + n) % 2)

	if (m+n)%2 != 0 {
		return float64(mergeArr[(m+n-1)/2])
	}

	return float64((mergeArr[(m+n)/2] + mergeArr[(m+n)/2-1])) / 2
}
