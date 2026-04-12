package handler

func MajorityElementArray(nums []int) int {
	// uses Boyer-Moore Voting Algorithm
	// The Boyer–Moore Voting Algorithm is a smart way to find the majority element in an array using:
	// O(n) time (one pass)
	// O(1) space (no extra memory)

	retunCount := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			retunCount = num
		}
		if num == retunCount {
			count++
		} else {
			count--
		}
	}

	return retunCount
}
