package two_sum

func naiveSolution(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1}
}

func optimalSolution(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		} else {
			seen[n] = i
		}
	}

	return []int{-1}
}
