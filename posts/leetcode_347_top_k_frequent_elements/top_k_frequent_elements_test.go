package top_k_frequent_elements

import (
	"testing"
)

func assertArrays(a, b []int) bool {
	m := make(map[int]int)

	for _, i := range a {
		m[i]++
	}

	for _, i := range b {
		m[i]--
		if m[i] == 0 {
			delete(m, i)
		}
		if m[i] < 0 {
			return false
		}
	}

	return len(m) == 0
}

var functions = []struct {
	Name string
	F    func([]int, int) []int
}{
	{"Naive", Naive},
	{"Bucket Sort", BucketSort},
}

var cases = []struct {
	Nums     []int
	K        int
	Expected []int
}{
	{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	{[]int{1}, 1, []int{1}},
	{[]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
}

func TestTopKFrequent(t *testing.T) {
	for _, function := range functions {
		t.Run(function.Name, func(t *testing.T) {
			for _, tt := range cases {
				got := function.F(tt.Nums, tt.K)

				if !assertArrays(got, tt.Expected) {
					t.Errorf("got %v want %v", got, tt.Expected)
				}
			}
		})
	}
}
