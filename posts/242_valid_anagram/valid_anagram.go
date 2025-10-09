package valid_anagram

func Naive(s, t string) bool {
	m := make(map[rune]int)

	for _, r := range s {
		if _, ok := m[r]; !ok {
			m[r] = 1
		} else {
			m[r]++
		}
	}

	for _, r := range t {
		if _, ok := m[r]; !ok {
			return false
		} else if m[r] == 1 {
			delete(m, r)
		} else {
			m[r]--
		}
	}

	return len(m) == 0
}

func TrickSolution(s string, t string) bool {
	var arr [26]int

	for _, ch := range s {
		arr[ch-'a']++
	}

	for _, ch := range t {
		arr[ch-'a']--
	}

	for _, count := range arr {
		if count != 0 {
			return false
		}
	}

	return true
}
