package http

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
)

func GetHtml(url string) goquery.Document {
	resp, _ := http.Get(url)
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	return *doc
}

func Download(url string, output string) {
	_, fileName := filepath.Split(url)
	log.Println("下载", url)
	resp, _ := http.Get(url)

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
