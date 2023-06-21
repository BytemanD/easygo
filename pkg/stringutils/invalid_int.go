package stringutils

import (
	"fmt"
	"strconv"

	"github.com/wxnacy/wgo/arrays"
)

func validMin(value int, min int) error {
	if value < min {
		return fmt.Errorf("%d is less than %d", value, min)
	}
	return nil
}
func validMax(value int, max int) error {
	if value > max {
		return fmt.Errorf("%d is greater than %d", value, max)
	}
	return nil
}

func MustLessThen(arg string, min int) (int, error) {
	value, err := strconv.Atoi(arg)
	if err != nil {
		return value, err
	}
	return value, validMax(value, min)
}

func MustGreaterThan(arg string, min int) (int, error) {
	value, err := strconv.Atoi(arg)
	if err != nil {
		return value, err
	}
	return value, validMin(value, min)
}

func MustInRange(arg string, min int, max int) (int, error) {
	value, err := MustLessThen(arg, min)
	if err != nil {
		return value, err
	}
	return value, validMax(value, min)
}

func MustInIntChoises(arg string, choises []int) (int, error) {
	value, err := strconv.Atoi(arg)
	if err != nil {
		return value, err
	}
	if arrays.Contains(choises, value) < 0 {
		return value, fmt.Errorf("%d not in %v", value, choises)
	}
	return value, nil
}
