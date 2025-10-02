package two_sum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	functions := []struct {
		Name     string
		Function func([]int, int) []int
	}{
		{"naiveSolution", naiveSolution},
		{"optimalSolution", optimalSolution},
	}

	cases := []struct {
		Input    []int
		Target   int
		Expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, f := range functions {
		for _, tt := range cases {
			got := f.Function(tt.Input, tt.Target)
			if !slices.Equal(got, tt.Expected) {
				t.Errorf("%s(%v, %v)[]int = %v, expected: %v", f.Name, tt.Input, tt.Target, got, tt.Expected)
			}
		}
	}
}
