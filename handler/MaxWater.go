package handler

import (
	"math"
)

func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		h := int(math.Min(float64(height[left]), float64(height[right])))
		w := right - left
		area := h * w

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
