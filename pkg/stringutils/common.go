package stringutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/text/width"
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
func FormatLen(s string) int {
	length := len(s)
	for _, r := range s {
		fmt.Println("RuneLen", utf8.RuneLen(r))
		if (r >= 0xFF10 && r <= 0xFF19) || (r >= 0xFF01 && r <= 0xFF60) {
			length -= 1
			continue
		}
		if unicode.Is(unicode.Han, r) {
			length -= 1
			continue
		}
		if r >= 0x1F300 {
			length -= 3
			continue
		}
		if unicode.Is(unicode.Symbol, r) {
			length -= 2
			continue
		}
	}
	return length
}
func TextWidth(s string) int {
	w := 0
	for _, r := range s {
		switch width.LookupRune(r).Kind() {
		case width.EastAsianFullwidth, width.EastAsianWide:
			w += 2
		case width.EastAsianHalfwidth, width.EastAsianNarrow,
			width.Neutral, width.EastAsianAmbiguous:
			w += 1
		}
	}
	return w
}

// Parse a string to substrings
//
// e.g. if width is 2:
//
//	"a"    -> ["a"]
//	"abc"  -> ["ab", "c"]
//	"abcd" -> ["ab", "cd"]
func SubStrings(s string, width int) []string {
	result := []string{}
	subS := []rune{}
	for _, r := range s {
		subS = append(subS, r)
		if len(subS) >= width {
			result = append(result, string(subS))
			subS = []rune{}
		}
	}
	if len(subS) != 0 {
		result = append(result, string(subS))
	}
	return result
}
