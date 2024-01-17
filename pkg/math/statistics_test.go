package math

import (
	"testing"
)

func TestSumInt(t *testing.T) {
	result := SumInt(1, 2, 3)
	if int(result) != 6 {
		t.Errorf("expect 6 but got %v", result)
	}
}
