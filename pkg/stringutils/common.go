package stringutils

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func ContainsString(stringList []string, s string) bool {
	for _, str := range stringList {
		if s == str {
			return true
		}
	}
	return false
}

func IsUUID(s string) bool {
	uuid.NewV4()
	if _, err := uuid.FromString(s); err != nil {
		return false
	} else {
		return true
	}
}

func PathJoin(path ...string) string {
	return strings.Join(path, "/")
}
