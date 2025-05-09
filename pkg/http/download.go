package http

import (
	"container/list"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/BytemanD/go-console/console"
	"github.com/go-resty/resty/v2"

	"github.com/BytemanD/easygo/pkg/fileutils"
	"github.com/BytemanD/easygo/pkg/progress"
	"github.com/BytemanD/easygo/pkg/syncutils"
	"github.com/PuerkitoBio/goquery"
)

func GetHtml(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		console.Warn("get url failed: %s", err)
		return nil
	}
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	return doc
}

func UrlJoin(url1 string, url2 string) string {
	re := regexp.MustCompile("^http(s)://")
	if re.FindString(url2) != "" {
		return url2
	} else {
		return strings.TrimRight(url1, "/") + "/" + strings.TrimLeft(url2, "/")
	}
}

func GetLinks(url string, regex string) list.List {
	doc := GetHtml(url)
	links := list.New()
	if doc == nil {
		return *links
	}
	reg := regexp.MustCompile(regex)
	selection := doc.Find("a")
	if selection == nil {
		return *links
	}
	selection.Each(func(_ int, s *goquery.Selection) {
		href := s.AttrOr("href", "")
		if regex == "" || reg.FindString(href) != "" {
			links.PushBack(UrlJoin(url, href))
		}
	})
	return *links
}

type ProgressWriter struct {
	ToalBytes int64
	completed int
}

func (w *ProgressWriter) Write(p []byte) (n int, err error) {
	w.completed += len(p)
	fmt.Printf("=== > %d/%d\n", w.completed, w.ToalBytes)
	return len(p), nil
}

func Download(url string, output string, showProgress bool) error {
	_, fileName := filepath.Split(url)
	client := resty.New().SetRetryCount(3).SetRetryWaitTime(time.Second)
	if !showProgress {
		_, err := client.SetOutputDirectory(output).R().SetOutput(fileName).Get(url)
		return err
	}

	console.Debug("fetch %s ...", url)
	resp, err := client.SetDoNotParseResponse(true).R().Get(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("[%d] %s %s", resp.StatusCode(), resp.Status(), resp.Body())
	}

	fp := fileutils.FilePath{Path: output}
	if err := fp.MakeDirs(); err != nil {
		return err
	}
	outputFile, err := os.OpenFile(path.Join(output, fileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	console.Debug("saving to %s", path.Join(output, fileName))
	writers := []io.Writer{outputFile}

	fmt.Println(resp.RawResponse.ContentLength)
	if resp.RawResponse.ContentLength > 0 {
		pbrWriter := progress.DefaultBytesWriter(fileName, resp.RawResponse.ContentLength)
		writers = append(writers, pbrWriter)
	}
	return resp.RawResponse.Write(io.MultiWriter(writers...))
}

func DownloadLinksInHtml(url string, regex string, output string) {
	console.Info("开始解析: %s", url)
	links := GetLinks(url, regex)

	console.Info("链接数量: %d", links.Len())
	if links.Len() <= 0 {
		os.Exit(0)
	}
	downloadLinks := []string{}
	link := links.Front()
	for i := 0; i < links.Len(); i++ {
		downloadLinks = append(downloadLinks, fmt.Sprintf("%s", link.Value))
		link = link.Next()
	}
	console.Info("开始下载(总数: %d), 保存路径: %s ...", len(downloadLinks), output)
	taskGroup := syncutils.TaskGroup[string]{
		Items:        downloadLinks,
		ShowProgress: true,
		Title:        fmt.Sprintf("下载 %d 个文件", len(downloadLinks)),
		Func: func(item string) error {
			console.Debug("开始下载: %s", item)
			if err := Download(item, output, false); err != nil {
				console.Error("下载失败: %s", item)
				return err
			} else {
				console.Info("下载完成: %s", item)
				return nil
			}
		},
	}
	taskGroup.Start()
	console.Info("下载结束")
}
