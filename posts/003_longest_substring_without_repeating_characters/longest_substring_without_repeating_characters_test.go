package longest_substring_without_repeating_characters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	functions := []struct {
		Name        string
		Function    func(string) int
		MeantToPass bool
	}{
		{"Naive", Naive, false},
		{"Let's Try This Again", NaiveLetsTryThisAgain, true},
		{"After Some Research", AfterSomeResearch, true},
		{"After Some Research With Pruning", AfterSomeResearchWithPruning, true},
	}

	cases := []struct {
		Input    string
		Expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"aab", 2},
		{"dvdf", 3},
	}

	for _, f := range functions {
		if !f.MeantToPass {
			continue
		}

		for _, tt := range cases {
			got := f.Function(tt.Input)
			if got != tt.Expected {
				t.Errorf("%s: Input: %s, Output: %d, Expected: %d", f.Name, tt.Input, got, tt.Expected)
			}
		}
	}
}
