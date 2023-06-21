package stringutils

import (
	"fmt"
	"regexp"

	"github.com/wxnacy/wgo/arrays"
)

func MustInStringChoises(arg string, choises []string) error {
	if arrays.Contains(choises, arg) < 0 {
		return fmt.Errorf("%s not in %s", arg, choises)
	}
	return nil
}

func MustMatch(arg string, pattern string) error {
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		return fmt.Errorf("invalid pattern %s", pattern)
	}
	if matched {
		return nil
	} else {
		return fmt.Errorf("'%s' is not match '%s'", arg, pattern)
	}
}
