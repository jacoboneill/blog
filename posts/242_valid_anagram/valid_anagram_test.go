package valid_anagram

import "testing"

var functions = []struct {
	Name string
	F    func(string, string) bool
}{
	{"Naive", Naive},
	{"Trick Solution", TrickSolution},
}

var cases = []struct {
	S        string
	T        string
	Expected bool
}{
	{"anagram", "nagaram", true},
	{"rat", "car", false},
}

func TestValidAnagram(t *testing.T) {
	for _, function := range functions {
		t.Run(function.Name, func(t *testing.T) {
			for _, tt := range cases {
				got := function.F(tt.S, tt.T)
				if got != tt.Expected {
					t.Errorf("got %t, want %t", got, tt.Expected)
				}
			}
		})
	}
}

func BenchmarkValidAnagram(b *testing.B) {
	for _, function := range functions {
		b.Run(function.Name, func(b *testing.B) {
			for _, tt := range cases {
				for range b.N {
					function.F(tt.S, tt.T)
				}
			}
		})
	}
}
