package fileutils

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

type FilePath struct {
	Path     string
	fileInfo *fs.FileInfo
}

func (fp *FilePath) getFileInfo() (fs.FileInfo, error) {
	if fp.fileInfo == nil {
		fi, err := os.Stat(fp.Path)
		fp.fileInfo = &fi
		return *fp.fileInfo, err
	}
	return *fp.fileInfo, nil
}

func (fp *FilePath) Exists() bool {
	_, err := fp.getFileInfo()
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
func (fp *FilePath) IsDir() bool {
	if fp.Exists() {
		return true
	}
	return (*fp.fileInfo).IsDir()
}

func (fp *FilePath) IsFile() bool {
	return !fp.IsDir()
}

func (fp *FilePath) ReadLines(file string) ([]string, error) {
	if !fp.Exists() {
		return nil, fmt.Errorf("%s 不存在", fp.Path)
	}
	if !fp.IsFile() {
		return nil, fmt.Errorf("%s 不是文件", fp.Path)
	}

	f, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("文件打开失败, %s", err)
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("文件读取失败 %s", err)
	}
	return strings.Split(string(bytes), "\n"), nil
}

func (fp *FilePath) MakeDirs() error {
	if !fp.Exists() {
		return os.MkdirAll(fp.Path, os.ModePerm)
	}
	if fp.IsFile() {
		return fmt.Errorf("已存在文件 %s", fp.Path)
	}
	return nil
}
