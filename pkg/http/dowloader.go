package http

import (
	"bufio"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/BytemanD/easygo/pkg/fileutils"
	"github.com/BytemanD/easygo/pkg/global/logging"
)

type HttpFile struct {
	Url  string
	Name string
	size int
}

func (f HttpFile) GuessName() string {
	if f.Name != "" {
		return f.Name
	}
	urlParsed, err := url.Parse(f.Url)
	if err != nil {
		return ""
	}
	_, fileName := filepath.Split(urlParsed.Path)
	return fileName
}

func (f HttpFile) GetSize() int {
	return f.size
}

type HttpDownloader struct {
	Output string
}

func (downloader HttpDownloader) Download(file *HttpFile) error {
	if file.Name == "" {
		file.Name = file.GuessName()
	}
	logging.Debug("下载: %s -> %s", file.Url, file.Name)
	resp, err := http.Get(file.Url)
	if err != nil {
		logging.Error("下载 %s 失败, %s", file.Url, err)
	}
	defer resp.Body.Close()
	logging.Debug("size: %v", resp.Header.Get("Content-Length"))
	file.size, _ = strconv.Atoi(resp.Header.Get("Content-Length"))
	fp := fileutils.FilePath{Path: downloader.Output}
	if err := fp.MakeDirs(); err != nil {
		return err
	}
	outputPath := path.Join(downloader.Output, file.Name)
	outputFile, _ := os.Create(outputPath)
	defer outputFile.Close()
	wt := bufio.NewWriter(outputFile)
	io.Copy(wt, resp.Body)
	logging.Debug("下载 %s 完成", file.Url)
	wt.Flush()
	return nil
}
