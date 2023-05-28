package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

const SCHEME string = "http"
const HOST string = "www.bingimg.cn"
const FILE_NAME_MAX_SIZE int = 50
const URL_GET_IMAGES_PAGE string = "%s://%s/list%s"

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

func BingImgDownload(page int, uhd bool, output string) {
	log.Printf("下载页面 %d, 保存位置: %s", page, output)
	url := fmt.Sprintf(URL_GET_IMAGES_PAGE, SCHEME, HOST, strconv.Itoa(page))
	resp, _ := http.Get(url)
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	links := list.New()

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		link := s.AttrOr("href", "")
		if strings.HasSuffix(link, ".jpg") {
			if !uhd || strings.Contains(link, "/uhd/") {
				links.PushBack(link)
			}
		}
	})
	for link := links.Front(); link != nil; link = link.Next() {
		Download(fmt.Sprintf("%s", link.Value), output)
		fmt.Println("下载", link.Value)
	}

}

func main() {
	var uhd bool
	var output string

	var rootCmd = &cobra.Command{
		Use:   "magic-pocket",
		Short: "常用工具合集",
		Long:  "Golang 实现的工具合集",
	}
	var bingImgDownload = &cobra.Command{
		Use:   "get-bing-img",
		Short: "下载bing高质量壁纸",
		Long:  "从www.bingimg.cn网站下载高质量壁纸",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("requires a page argument")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			page, err := strconv.Atoi(args[0])
			if err != nil {
				log.Println("ERROR", err)
				os.Exit(1)
			}
			BingImgDownload(page, uhd, output)
		},
	}
	bingImgDownload.Flags().BoolVar(&uhd, "uhd", false, "仅下载4K壁纸")
	bingImgDownload.Flags().StringVarP(&output, "output", "o", "./", "保存路径")
	rootCmd.AddCommand(bingImgDownload)
	rootCmd.Execute()
}
