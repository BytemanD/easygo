package commands

import (
	"container/list"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	httpLib "github.com/BytemanD/easygo/pkg/http"
	"github.com/BytemanD/easygo/pkg/stringutils"
	"github.com/BytemanD/easygo/pkg/syncutils"
	"github.com/BytemanD/easygo/pkg/table"
	"github.com/BytemanD/go-console/console"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	FILE_NAME_MAX_SIZE int = 50

	UHD_ONLY           string = "only"
	UHD_INCLUDE        string = "include"
	UHD_NO             string = "no"
	URL_WWW_BINGIMG_CN string = "http://www.bingimg.cn/list%s"

	URL_BING_WDBYTE_COM string = "https://bing.wdbyte.com"
)

var UHD_CHOICES = []string{UHD_ONLY, UHD_INCLUDE, UHD_NO}

func bingImgDownload(page int8, uhd string, output string) {
	url := fmt.Sprintf(URL_WWW_BINGIMG_CN, strconv.Itoa(int(page)))
	console.Info("解析页面 %s", url)
	doc := httpLib.GetHtml(url)
	links := list.New()

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		link := s.AttrOr("href", "")
		if strings.HasSuffix(link, ".jpg") {
			if uhd == UHD_INCLUDE {
				links.PushBack(link)
			} else {
				if uhd == UHD_ONLY && strings.Contains(link, "/uhd/") {
					links.PushBack(link)
				} else if uhd == UHD_NO && !strings.Contains(link, "/uhd/") {
					links.PushBack(link)
				}
			}
		}
	})
	if links.Len() == 0 {
		console.Warn("页面 %s 无图片链接", url)
		os.Exit(0)
	}
	console.Info("开始下载, 保存路径: %s", output)
	for link := links.Front(); link != nil; link = link.Next() {
		console.Debug("下载 %s", link.Value)
		httpLib.Download(fmt.Sprintf("%s", link.Value), output, false)
	}
}

var FetchWallpaperCmd = &cobra.Command{
	Use:   "fetch-wallpaper",
	Short: "下载壁纸",
}

var BingImgCmd = &cobra.Command{
	Use:              "bingimg <page>",
	Short:            "下载 www.bingimg.cn 网站的壁纸",
	TraverseChildren: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		if _, err := stringutils.MustGreaterThan(args[0], 1); err != nil {
			return fmt.Errorf("invalid arg 'page': %s", err)
		}
		uhd, _ := cmd.Flags().GetString("uhd")
		if err := stringutils.MustInStringChoises(uhd, UHD_CHOICES); err != nil {
			return fmt.Errorf("invalid flag 'uhd': %s", err)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		uhd, _ := cmd.Flags().GetString("uhd")
		output, _ := cmd.Flags().GetString("output")
		endPage, _ := cmd.Flags().GetInt8("end-page")
		pageInt, _ := strconv.Atoi(args[0])
		page := int8(pageInt)
		if endPage <= 0 {
			endPage = page
		}
		for i := page; i <= endPage; i++ {
			bingImgDownload(i, uhd, output)
		}
	},
}
var BingImgWdbyteCmd = &cobra.Command{
	Use:              "bingimg-wdbyte",
	Short:            "下载 bing.wdbyte.com 网站的壁纸",
	TraverseChildren: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(0)(cmd, args); err != nil {
			return err
		}
		date, _ := cmd.Flags().GetString("date")
		if date != "" {
			if err := stringutils.MustMatch(date, "^[0-9]+-[0-9]+$"); err != nil {
				return fmt.Errorf("invalid flag 'date', %s", err)
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
		output, _ := cmd.Flags().GetString("output")
		workers, _ := cmd.Flags().GetInt("workers")
		limit, _ := cmd.Flags().GetInt("limit")

		reqUrl := fmt.Sprintf("%s/%s", URL_BING_WDBYTE_COM, date)
		console.Info("解析页面 %s", reqUrl)
		doc := httpLib.GetHtml(reqUrl)
		files := []httpLib.HttpFile{}
		doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			href := s.AttrOr("href", "")
			if err := stringutils.MustMatch(href, `.*id=.+\.jpg`); err == nil {
				if limit > 0 && len(files) >= limit {
					return
				}
				urlParsed, _ := url.Parse(href)
				fileName := urlParsed.Query().Get("id")
				files = append(files, httpLib.HttpFile{Name: fileName, Url: href})
			}
		})
		if len(files) == 0 {
			console.Warn("页面 %s 无图片链接", reqUrl)
			os.Exit(0)
		}
		type webFile struct {
			Name   string
			Size   int
			Result string
			Spend  time.Duration
		}
		webFiles := []webFile{}
		downloader := httpLib.HttpDownloader{Output: output}
		taskGroup := syncutils.TaskGroup{
			Items:        files,
			ShowProgress: true,
			Title:        fmt.Sprintf("download %d picture(s)", len(files)),
			MaxWorker:    workers,
			Func: func(item interface{}) error {
				file := item.(httpLib.HttpFile)
				startTime := time.Now()
				err := downloader.Download(&file)
				var result string
				if err != nil {
					console.Error("下载失败: %s", file.Url)
					result = "😞"
				} else {
					console.Info("下载完成: %s, Size: %d", file.Name, file.GetSize())
					result = "👌"
				}
				webFiles = append(webFiles, webFile{
					Name: file.Name, Size: file.GetSize(),
					Result: result,
					Spend:  time.Since(startTime),
				})
				return err
			},
		}
		console.Info("开始下载(总数: %d), 保存路径: %s ...", len(files), output)
		taskGroup.Start()

		result, _ := table.NewItemsTable(
			[]string{"Result", "Name", "Size", "Spend"}, webFiles,
		).SetStyle(
			table.StyleLight, color.FgCyan,
		).EnableAutoIndex().Render()
		fmt.Println(result)
	},
}

func init() {
	BingImgCmd.Flags().StringP("uhd", "u", "only", "下载4K分辨率")
	BingImgCmd.Flags().StringP("output", "o", "./", "保存路径")
	BingImgCmd.Flags().Int8P("end-page", "e", 0, "结束的页面, 默认和page相同")

	BingImgWdbyteCmd.Flags().StringP("output", "o", "./", "保存路径")
	BingImgWdbyteCmd.Flags().Int("workers", 0, "并发数")
	BingImgWdbyteCmd.Flags().Int("limit", 0, "下载条数")
	BingImgWdbyteCmd.Flags().String("date", "", "日期, 例如: 2023-09")

	FetchWallpaperCmd.AddCommand(BingImgCmd, BingImgWdbyteCmd)
}
