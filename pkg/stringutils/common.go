package stringutils

import (
	"bytes"
	"encoding/json"
	"fmt"
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

const GB = 1 << 30
const MB = 1 << 20
const KB = 1 << 10

func HumanBytes(size int) string {
	if size >= GB {
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	} else if size >= MB {
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	} else if size >= KB {
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	} else {
		return fmt.Sprintf("%.2f B", float64(size))
	}
}

func JsonDumpsIndent(v interface{}) (string, error) {
	jsonBytes, _ := json.Marshal(v)
	var buffer bytes.Buffer
	json.Indent(&buffer, jsonBytes, "", "  ")
	return buffer.String(), nil
}
