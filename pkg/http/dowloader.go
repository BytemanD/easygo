package http

import (
	"bufio"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/BytemanD/easygo/pkg/fileutils"
	"github.com/BytemanD/easygo/pkg/global/logging"
)

type HttpFile struct {
	Url  string
	Name string
}

type HttpDownloader struct {
	Output string
}

func (downloader HttpDownloader) Download(files []HttpFile) error {
	for _, file := range files {
		if file.Name == "" {
			urlParsed, err := url.Parse(file.Url)
			if err != nil {
				return err
			}
			_, file.Name = filepath.Split(urlParsed.Path)
		}
		logging.Debug("下载地址: %s", file.Url)
		logging.Info("下载文件 %s", file.Name)
		resp, err := http.Get(file.Url)
		if err != nil {
			logging.Error("下载 %s 失败, %s", file.Url, err)
		}
		defer resp.Body.Close()
		fp := fileutils.FilePath{Path: downloader.Output}
		if err := fp.MakeDirs(); err != nil {
			return err
		}
		outputPath := path.Join(downloader.Output, file.Name)
		outputFile, _ := os.Create(outputPath)
		defer outputFile.Close()
		wt := bufio.NewWriter(outputFile)
		io.Copy(wt, resp.Body)
		wt.Flush()
	}
	return nil
}
