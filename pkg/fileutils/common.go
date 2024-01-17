package fileutils

import (
	"path"
	"strings"
)

func PathExtSplit(file string) (string, string) {
	ext := path.Ext(path.Base(file))
	name := strings.TrimSuffix(path.Base(file), ext)
	return name, ext
}
