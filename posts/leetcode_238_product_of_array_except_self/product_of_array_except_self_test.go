package main

import (
	"slices"
	"testing"
)

var cases = []struct {
	Input    []int
	Expected []int
}{
	{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	{[]int{1, 2, 4, 6}, []int{48, 24, 12, 8}},
	{[]int{-1, 0, 1, 2, 3}, []int{0, -6, 0, 0, 0}},
}

var functions = []struct {
	Name string
	F    func([]int) []int
}{
	{"Cheating", Cheating},
	{"BruteForce", BruteForce},
	{"Postfix Prefix v1", PostfixPrefixV1},
	{"Postfix Prefix v2", PostfixPrefixV2},
}

func TestProductExceptSelf(t *testing.T) {
	for _, function := range functions {
		for _, tt := range cases {
			got := function.F(tt.Input)

			if !slices.Equal(got, tt.Expected) {
				t.Errorf("%s: got %v, want %v", function.Name, got, tt.Expected)
			}
		}
	}
}
