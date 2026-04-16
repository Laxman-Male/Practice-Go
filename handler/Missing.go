package handler

import (
	"math"
	"sort"
)

func FindMissingAndRepeatedValues(grid [][]int) []int {

	var ConcarArray = []int{}
	maxCount := 0
	maxKey := 0
	missingKey := 0
	Result := []int{}
	sortKey := []int{}
	for i := range len(grid) {
		for j := range len(grid[0]) {
			var a = grid[i][j]
			ConcarArray = append(ConcarArray, a)
		}
	}
	freq := make(map[int]int)
	for _, value := range ConcarArray {
		freq[value]++
	}
	for key, value := range freq {
		if value > maxCount {
			maxCount = value
			maxKey = key
		}

		sortKey = append(sortKey, key)

	}
	Result = append(Result, maxKey)

	//this sort the array
	sort.Ints(sortKey)

	for key := 0; key < len(sortKey)-1; key++ {

		if sortKey[key+1]-sortKey[key] == 1 {

		} else {

			missingKey = sortKey[key] + 1
		}

		if (key+1 == len(sortKey)-1) && missingKey == 0 {
			missingKey = int(math.Pow(float64(len(grid)), 2))
			var found = false
			for i, _ := range ConcarArray {
				if missingKey == ConcarArray[i] {
					found = true
					break
				}
			}
			if found == true {
				// missingKey = int(math.Sqrt(float64(missingKey)))
				missingKey = 1
			}

		}
	}
	Result = append(Result, missingKey)

	return Result
}
