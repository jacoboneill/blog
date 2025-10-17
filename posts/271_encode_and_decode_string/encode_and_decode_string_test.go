package encode_and_decode_string

import "testing"

func assertArray(a, b []string) bool {
	m := make(map[string]int)

	for _, elem := range a {
		m[elem]++
	}

	for _, elem := range b {
		m[elem]--
		if m[elem] == 0 {
			delete(m, elem)
		}
		if m[elem] < 0 {
			return false
		}
	}

	return len(m) == 0
}

var functions = []struct {
	Name   string
	Encode func([]string) string
	Decode func(string) []string
}{
	{"Content Length", ContentLengthEncode, ContentLengthDecode},
	{"Length Delimeter", LengthDelimeterEncode, LengthDelimeterDecode},
}

var cases = [][]string{
	{"Hello", "World"},
	{"Hello,", "World"},
}

func TestEncodeAndDecode(t *testing.T) {
	for _, function := range functions {
		t.Run(function.Name, func(t *testing.T) {
			for _, tt := range cases {
				encoded := function.Encode(tt)

				got := function.Decode(encoded)
				if !assertArray(got, tt) {
					t.Errorf("got: %v, want: %v", got, tt)
				}
			}
		})
	}
}

func BenchmarkEncodeAndDecode(b *testing.B) {
	for _, function := range functions {
		b.Run(function.Name, func(b *testing.B) {
			for _, tt := range cases {
				for range b.N {
					function.Decode(function.Encode(tt))
				}
			}
		})
	}
}
