package group_anagrams

import "slices"

func Sorting(strs []string) [][]string {
	sortString := func(s string) string {
		b := []byte(s)
		slices.Sort(b)

		return string(b)
	}

	m := make(map[string][]string)

	for _, s := range strs {
		sorted := sortString(s)
		m[sorted] = append(m[sorted], s)
	}

	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func MapArray(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int

		for _, c := range s {
			count[c-'a']++
		}

		m[count] = append(m[count], s)
	}

	res := make([][]string, 0, len(strs))
	for _, v := range m {
		res = append(res, v)
	}

	return res
}
