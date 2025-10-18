package encode_and_decode_string

import (
	"fmt"
	"strconv"
	"strings"
)

func ContentLengthEncode(strs []string) string {
	var b strings.Builder

	for _, str := range strs {
		fmt.Fprintf(&b, "%03d%s", len(str), str)
	}

	return b.String()
}

func ContentLengthDecode(s string) []string {
	var strs []string

	i := 0
	for i < len(s)-1 {
		l, err := strconv.Atoi(s[i : i+3])
		if err != nil {
			return []string{}
		}

		strs = append(strs, s[i+3:i+3+l])
		i += 3 + l
	}

	return strs
}

var delimeter = '#'

func LengthDelimeterEncode(strs []string) string {
	var b strings.Builder

	for _, str := range strs {
		fmt.Fprintf(&b, "%d%c%s", len(str), delimeter, str)
	}

	return b.String()
}

func LengthDelimeterDecode(s string) []string {
	var strs []string

	i := 0
	for i <= len(s)-1 {
		var dp int
		for j := i; j < len(s); j++ {
			if rune(s[j]) == delimeter {
				dp = j
				break
			}
		}

		l, err := strconv.Atoi(s[i:dp])
		if err != nil {
			return []string{}
		}

		strs = append(strs, s[dp+1:dp+1+l])
		i = dp + l + 1
	}

	return strs
}
