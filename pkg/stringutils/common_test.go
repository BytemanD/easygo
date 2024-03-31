package stringutils

import (
	"testing"
)

func TestTextWidth(t *testing.T) {
	tester := map[string]int{
		"a":  0,
		"你":  2,
		"()": 2,
	}
	for s, expect := range tester {
		result := TextWidth(s)
		if result != expect {
			t.Errorf("UnicodeNum(%s) = %v, not %v", s, result, expect)
		}
	}
}
