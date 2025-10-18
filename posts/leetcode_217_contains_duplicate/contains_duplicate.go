package contains_duplicate

func BruteForce(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func HashSet(nums []int) bool {
	s := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := s[n]; ok {
			return true
		}

		s[n] = struct{}{}
	}

	return false
}
