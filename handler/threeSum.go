package handler

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {

		// Skip duplicate i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for left
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				// Skip duplicates for right
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--

			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
