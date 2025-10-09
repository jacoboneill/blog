package contains_duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	functions := []struct {
		Name string
		f    func([]int) bool
	}{
		{"Brute Force", BruteForce},
		{"Hash Set", HashSet},
	}

	cases := []struct {
		Numbers  []int
		Expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, function := range functions {
		for _, tt := range cases {
			got := function.f(tt.Numbers)
			if got != tt.Expected {
				t.Errorf("%s: got %t want %t", function.Name, got, tt.Expected)
			}
		}
	}

}
