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
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/fjboy/magic-pocket/pkg/global/logging"
)

func GetHtml(url string) goquery.Document {
	resp, _ := http.Get(url)
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	return *doc
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

	reg := regexp.MustCompile(regex)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href := s.AttrOr("href", "")
		if regex == "" || reg.FindString(href) != "" {
			links.PushBack(UrlJoin(url, href))
		}
	})
	return *links
}

func Download(url string, output string) {
	_, fileName := filepath.Split(url)
	resp, err1 := http.Get(url)
	if err1 != nil {
		logging.Error("下载 %s 失败, 原因: %s", url, err1)
		return
	}

	defer resp.Body.Close()

	_, err := os.Stat(output)
	if !os.IsExist(err) {
		os.MkdirAll(output, os.ModePerm)
	}
	outputPath := path.Join(output, fileName)
	outputFile, _ := os.Create(outputPath)
	defer outputFile.Close()

	wt := bufio.NewWriter(outputFile)
	io.Copy(wt, resp.Body)
	wt.Flush()
}

func DownloadLinksInHtml(url string, regex string, output string) {
	links := GetLinks(url, regex)

	var (
		wg sync.WaitGroup
	)
	link := links.Front()
	wg.Add(links.Len())
	logging.Info("开始下载, 数量 %d", links.Len())
	for i := 0; i < links.Len(); i++ {
		go func(aaa list.Element) {
			Download(fmt.Sprintf("%s", aaa.Value), output)
			logging.Debug("下载完成: %s ", aaa.Value)
			wg.Done()
		}(*link)
		link = link.Next()
		if link == nil {
			break
		}
	}
	wg.Wait()
	logging.Info("下载完成")
}
