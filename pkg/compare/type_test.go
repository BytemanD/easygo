package compare

import (
	"fmt"
	"testing"
)

func TestIsTypeInt(t *testing.T) {
	if IsType[int](1) != true {
		t.Errorf("IsType[int](1) is false")
		return
	}
	if IsType[int]("1") != false {
		t.Errorf(`IsType[int]("1") is true`)
		return
	}
}
func TestIsTypeError(t *testing.T) {
	if IsType[error](fmt.Errorf("error")) != true {
		t.Errorf(`IsType[error](fmt.Errorf("error")) is false`)
		return
	}
	if IsType[error](1) != false {
		t.Errorf(`IsType[error](1) is true`)
		return
	}
}

type demoStruct struct {
	Id string
}

func TestIsTypeStruct(t *testing.T) {
	if IsType[demoStruct](demoStruct{}) != true {
		t.Errorf(`IsType[demoStruct](demoStruct{}) is false`)
		return
	}
	if IsType[demoStruct](1) != false {
		t.Errorf(`IsType[demoStruct](1) is true`)
		return
	}
}
