package http

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/BytemanD/go-console/console"

	"github.com/BytemanD/easygo/pkg/fileutils"
	"github.com/BytemanD/easygo/pkg/progress"
	"github.com/BytemanD/easygo/pkg/stringutils"
	"github.com/BytemanD/easygo/pkg/syncutils"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
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

func Download(url string, output string, showProgress bool) error {
	_, fileName := filepath.Split(url)
	resp, err := http.Get(url)
	if err != nil {
		console.Error("下载 %s 失败, 原因: %s", url, err)
		return err
	}

	defer resp.Body.Close()

	fp := fileutils.FilePath{Path: output}
	if err := fp.MakeDirs(); err != nil {
		return err
	}

	outputPath := path.Join(output, fileName)
	outputFile, _ := os.Create(outputPath)
	defer outputFile.Close()

	if showProgress {
		if resp.Header.Get("Content-Length") != "" {
			size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
			console.Info("size: %s", stringutils.HumanBytes(size))
			pw := progress.NewProgressWriter(outputFile, size)
			pw.SetProgressColor(color.FgCyan)
			io.Copy(pw, resp.Body)
			pw.Wait()
			return nil
		} else {
			console.Warn("content-length is none for url: %s", url)
		}
	}

	wt := bufio.NewWriter(outputFile)
	wt.Flush()
	io.Copy(wt, resp.Body)
	return nil
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
	taskGroup := syncutils.TaskGroup{
		Items:        downloadLinks,
		Title:        fmt.Sprintf("下载: %s", url),
		ShowProgress: true,
		Func: func(item interface{}) error {
			url := item.(string)
			if err := Download(url, output, false); err != nil {
				console.Error("下载失败: %s", url)
				return err
			} else {
				console.Success("下载完成: %s", url)
				return nil
			}
		},
	}
	console.Info("开始下载(总数: %d), 保存路径: %s ...", len(downloadLinks), output)
	taskGroup.Start()
	console.Info("下载完成")
}
