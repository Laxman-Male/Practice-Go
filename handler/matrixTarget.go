package handler

func SearchMatrix(matrix [][]int, target int) bool {
	// binary search
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])

	left, right := 0, m*n-1

	for left <= right {
		mid := (left + right) / 2

		// Convert 1D index → 2D
		row := mid / n
		col := mid % n

		val := matrix[row][col]

		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
