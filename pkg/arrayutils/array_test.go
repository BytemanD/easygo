package arrayutils

import (
	"strings"
	"testing"
)

func testSplitStrings(t *testing.T, testCase []string, expect [][]string, splitSize int) {
	result := SplitStrings(testCase, splitSize)
	if len(result) != len(expect) {
		t.Errorf("SplitStrings(%s, 2) = %v, not %v", testCase, result, expect)
		return
	}
	for j, resultSplice := range result {
		if len(resultSplice) != len(expect[j]) {
			t.Errorf("SplitStrings(%s, 2) = %v, not %v", testCase, result, expect)
			return
		}
		if strings.Join(result[j], ",") != strings.Join(expect[j], ",") {
			t.Errorf("SplitStrings(%s, 2) = %v, not %v", testCase, result, expect)
			return
		}
	}
}

func TestSplitStrings(t *testing.T) {
	cases := [][]string{
		{"1", "2", "3", "4"},
		{"1", "2", "3", "4", "5"},
		{"1", "2", "3", "4", "5", "6"},
	}
	expects := [][][]string{
		{{"1", "2"}, {"3", "4"}},
		{{"1", "2"}, {"3", "4"}, {"5"}},
		{{"1", "2"}, {"3", "4"}, {"5", "6"}},
	}
	for i, testCase := range cases {
		testSplitStrings(t, testCase, expects[i], 2)
	}
	expects2 := [][][]string{
		{{"1", "2", "3"}, {"4"}},
		{{"1", "2", "3"}, {"4", "5"}},
		{{"1", "2", "3"}, {"4", "5", "6"}},
	}
	for i, testCase := range cases {
		testSplitStrings(t, testCase, expects2[i], 3)
	}

}
