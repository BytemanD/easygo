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

type myError struct {
	Message string
}

func (e myError) Error() string {
	return e.Message
}

func TestIsTypeCustomError(t *testing.T) {
	if IsType[myError](myError{Message: "error"}) != true {
		t.Errorf(`IsType[MyError](MyError{Message: "error"}) is false`)
		return
	}
	if IsType[myError](fmt.Errorf("error")) != false {
		t.Errorf(`IsType[MyError](fmt.Errorf("error")) is true`)
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

func TestIsStructOf(t *testing.T) {
	if IsStructOf("demoStruct", demoStruct{}) != true {
		t.Errorf(`IsStructOf(demoStruct, demoStruct{}) is false`)
		return
	}
	if IsStructOf("demoStruct", &demoStruct{}) != true {
		t.Errorf(`IsStructOf(demoStruct, &demoStruct{}) is false`)
		return
	}
	if IsStructOf("demoStruct", myError{}) != false {
		t.Errorf(`IsStructOf(demoStruct, myError{}) is true`)
		return
	}
	if IsStructOf("demoStruct", "demoStruct") != false {
		t.Errorf(`IsStructOf(demoStruct, "demoStruct) is true`)
		return
	}
	if IsStructOf("demoStruct", nil) != false {
		t.Errorf(`IsStructOf(demoStruct, "demoStruct) is true`)
		return
	}
}

func TestEqualsStruct(t *testing.T) {
	if EqualsStruct(nil, nil) != true {
		t.Errorf(`EqualsStruct(nil, nil) is false`)
		return
	}
	if EqualsStruct(nil, demoStruct{}) != false {
		t.Errorf(`EqualsStruct(nil, demoStruct{}) is true`)
		return
	}
	if EqualsStruct(demoStruct{}, nil) != false {
		t.Errorf(`EqualsStruct(demoStruct{}, nil) is true`)
		return
	}
	if EqualsStruct(demoStruct{}, demoStruct{}) != true {
		t.Errorf(`EqualsStruct(demoStruct{}, demoStruct{}) is false`)
		return
	}
	if EqualsStruct(demoStruct{}, &demoStruct{}) != true {
		t.Errorf(`EqualsStruct(demoStruct{}, &demoStruct{}) is false`)
		return
	}
	if EqualsStruct(&demoStruct{}, demoStruct{}) != true {
		t.Errorf(`EqualsStruct(&demoStruct{}, demoStruct{}) is false`)
		return
	}
	if EqualsStruct(&demoStruct{}, &demoStruct{}) != true {
		t.Errorf(`EqualsStruct(&demoStruct{}, &demoStruct{}) is false`)
		return
	}
	if EqualsStruct(demoStruct{}, myError{}) != false {
		t.Errorf(`EqualsStruct(demoStruct{}, myError{}) is true`)
		return
	}
}
