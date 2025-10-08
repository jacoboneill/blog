package longest_substring_without_repeating_characters

func Naive(s string) int {
	maxLength := 0
	start := 0

	seen := make(map[rune]bool)
	for i, c := range s {
		if _, ok := seen[c]; ok {
			if i-start > maxLength {
				maxLength = i - start
			}
			start = i + 1
			seen = make(map[rune]bool)
		} else if i == len(s)-1 && len(s)-start > maxLength {
			maxLength = len(s) - start
		} else {
			seen[c] = true
		}
	}
	return maxLength
}

func NaiveLetsTryThisAgain(s string) int {
	maxLength := 0

	for i := range s {
		if len(s)-i <= maxLength {
			break
		}
		seen := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			char := s[j]
			if seen[char] {
				break
			}
			seen[char] = true
			currentLength := j - i + 1
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}
	return maxLength
}

func AfterSomeResearch(s string) int {
	maxLength := 0
	seen := make(map[byte]bool)

	start := 0

	for end := 0; end < len(s); end++ {
		for _, ok := seen[s[end]]; ok; _, ok = seen[s[end]] {
			delete(seen, s[start])
			start++
		}

		seen[s[end]] = true
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

func AfterSomeResearchWithPruning(s string) int {
	maxLength := 0
	seen := make(map[byte]bool)

	start := 0

	for end := 0; end < len(s) && len(s)-start > maxLength; end++ {
		for _, ok := seen[s[end]]; ok; _, ok = seen[s[end]] {
			delete(seen, s[start])
			start++
		}

		seen[s[end]] = true
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}
