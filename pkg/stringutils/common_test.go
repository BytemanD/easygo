package stringutils

import (
	"testing"
)

func TestUnicodeNum(t *testing.T) {
	tester := map[string]int{
		"a":  0,
		"你":  1,
		"()": 0,
	}
	for s, expect := range tester {
		result := UnicodeNum(s)
		if result != expect {
			t.Errorf("UnicodeNum(%s) = %v, not %v", s, result, expect)
		}
	}
}
