package handler

import (
	"fmt"
	"sort"
)

func Merge(nums1 []int, m int, nums2 []int, n int) {

	var mergingArr1 = []int{}
	var mergingArr2 = []int{}
	for i, v := range nums1 {
		if i+1 <= m {
			mergingArr1 = append(mergingArr1, v)
		} else {

		}
	}
	mergingArr2 = append(mergingArr1, nums2...)
	sort.Ints(mergingArr2)

	// nums1 = mergingArr2
	copy(nums1, mergingArr2)
	fmt.Println("final", nums1)

}
