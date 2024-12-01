package slice

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	result := make([]int, 0)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if len(result) == 0 || x > result[len(result)-1] {
				result = append(result, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}

	}
	return result
}
