package numstr

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

func LessThanFold(v1, v2 string) bool {
	var s1, s2 string
	var r rune
	var n int
	for {
		// Search for numbers, keeping non-numeric
		for { // Read v1 until number
			if r, n = utf8.DecodeRuneInString(v1); n == 0 || isNum(r) {
				break
			}
			s1, v1 = s1+string(r), v1[n:]
		}
		for { // Read v2 until number
			if r, n = utf8.DecodeRuneInString(v2); n == 0 || isNum(r) {
				break
			}
			s2, v2 = s2+string(r), v2[n:]
		}
		if s1 != s2 {
			return strings.Compare(s1, s2) < 0
		}
		if v1 == "" || v2 == "" {
			return v1 == ""
		}
		s1, s2 = "", ""

		// Search for non-numeric, keeping numeric
		for { // Read v1 until non-number
			if r, n = utf8.DecodeRuneInString(v1); n == 0 || !isNum(r) {
				break
			}
			s1, v1 = s1+string(r), v1[n:]
		}
		for { // Read v2 until non-number
			if r, n = utf8.DecodeRuneInString(v2); n == 0 || !isNum(r) {
				break
			}
			s2, v2 = s2+string(r), v2[n:]
		}
		s1 = strings.TrimPrefix(s1, "0")
		s2 = strings.TrimPrefix(s2, "0")
		if s1 != s2 {
			n1, _ := strconv.Atoi(s1)
			n2, _ := strconv.Atoi(s2)
			return n1 < n2
		}
		if v1 == "" || v2 == "" {
			return v1 == ""
		}
		s1, s2 = "", ""
	}
}