package group_anagrams

import (
	"fmt"
	"testing"
)

func assertArray[T comparable](a, b []T) bool {
	m := make(map[T]int)
	for _, a_elem := range a {
		m[a_elem]++
	}

	for _, b_elem := range b {
		if _, ok := m[b_elem]; !ok {
			return false
		} else if m[b_elem] == 1 {
			delete(m, b_elem)
		} else {
			m[b_elem]--
		}
	}

	return len(m) == 0
}

func assert2DArray[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	for _, a_elem := range a {
		used := make([]bool, len(b))
		found := false
		for j, b_elem := range b {
			if !used[j] && assertArray(a_elem, b_elem) {
				found = true
				used[j] = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

var functions = []struct {
	Name string
	F    func([]string) [][]string
}{
	{"Sorting", Sorting},
	{"Map Array", MapArray},
}

var cases = []struct {
	Input    []string
	Expected [][]string
}{
	{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
	{[]string{""}, [][]string{{""}}},
	{[]string{"a"}, [][]string{{"a"}}},
}

func TestGroupAnagrams(t *testing.T) {
	for _, function := range functions {
		for _, tt := range cases {
			t.Run(fmt.Sprintf("%s:%v", function.Name, tt.Input), func(t *testing.T) {
				got := function.F(tt.Input)
				if !assert2DArray(got, tt.Expected) {
					t.Errorf("got %v, want %v", got, tt.Expected)
				}
			})
		}
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	for _, function := range functions {
		b.Run(function.Name, func(b *testing.B) {
			for _, tt := range cases {
				for range b.N {
					function.F(tt.Input)
				}
			}
		})
	}
}
