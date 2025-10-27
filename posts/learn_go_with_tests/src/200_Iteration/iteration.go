package iteration

import (
	"fmt"
	"strings"
)

const n = 5

// Repeat takes an input string, repeats it five times, and returns the string.
func Repeat(input string) string {
	var out string

	for range n {
		out += input
	}

	return out
}

// RepeatWithStringBuilder takes an input string, repeats it five times, and returns the string.
func RepeatWithStringBuilder(input string) string {
	var b strings.Builder

	for range n {
		b.WriteString(input)
	}

	return b.String()
}

// RepeatNTimes takes an input string and number of times to repeat, repeats it n times, and returns the string.
func RepeatNTimes(input string, n int) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("n has to be larger than 0")
	}

	var b strings.Builder

	for range n {
		b.WriteString(input)
	}

	return b.String(), nil
}
