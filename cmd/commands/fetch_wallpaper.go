package commands

import (
	"container/list"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/BytemanD/easygo/cli"
	httpLib "github.com/BytemanD/easygo/pkg/http"
	"github.com/BytemanD/easygo/pkg/stringutils"
	"github.com/BytemanD/easygo/pkg/syncutils"
	"github.com/BytemanD/easygo/pkg/table"
	"github.com/BytemanD/go-console/console"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

type bingImgCmdFlags struct {
	Uhd     *string
	Output  *string
	EndPage *int8
}
type bingImgWdbyteCmdFlags struct {
	Output  *string
	Workers *int
	Limit   *int
	Date    *string
}

func bingImgDownload(page int8, uhd string, output string) {
	url := fmt.Sprintf(URL_WWW_BINGIMG_CN, strconv.Itoa(int(page)))
	console.Info("解析页面 %s", url)
	doc := httpLib.GetHtml(url)
	if doc == nil {
		console.Error("get url %s failed", url)
		return
	}
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

var BingImgCmd = cli.NewCommand(cli.Cli[bingImgCmdFlags]{
	Use:              "bingimg <page>",
	Short:            "下载 www.bingimg.cn 网站的壁纸",
	TraverseChildren: true,
	RegistryFlags: func(fs *pflag.FlagSet) bingImgCmdFlags {
		return bingImgCmdFlags{
			Uhd:     fs.StringP("uhd", "u", "only", "下载4K分辨率"),
			Output:  fs.StringP("output", "o", "./", "保存路径"),
			EndPage: fs.Int8P("end-page", "e", 0, "结束的页面, 默认和page相同"),
		}
	},
	ValidArgs: []cobra.PositionalArgs{
		cobra.ExactArgs(1),
		func(cmd *cobra.Command, args []string) error {
			if _, err := stringutils.MustGreaterThan(args[0], 1); err != nil {
				return fmt.Errorf("invalid arg 'page': %s", err)
			}
			uhd, _ := cmd.Flags().GetString("uhd")
			if err := stringutils.MustInStringChoises(uhd, UHD_CHOICES); err != nil {
				return fmt.Errorf("invalid flag 'uhd': %s", err)
			}
			return nil
		},
	},
	Run: func(args []string, flags bingImgCmdFlags) error {
		pageInt, _ := strconv.Atoi(args[0])
		page := int8(pageInt)
		if *flags.EndPage <= 0 {
			flags.EndPage = &page
		}
		for i := page; i <= *flags.EndPage; i++ {
			bingImgDownload(i, *flags.Uhd, *flags.Output)
		}
		return nil
	},
})

var BingImgWdbyteCmd = cli.NewCommand(cli.Cli[bingImgWdbyteCmdFlags]{
	Use:              "bingimg-wdbyte",
	Short:            "下载 www.bingimg.cn 网站的壁纸",
	TraverseChildren: true,
	RegistryFlags: func(fs *pflag.FlagSet) bingImgWdbyteCmdFlags {
		return bingImgWdbyteCmdFlags{
			Output:  fs.StringP("output", "O", "./", "保存路径"),
			Workers: fs.Int("workers", 0, "并发数"),
			Limit:   fs.Int("limit", 0, "下载条数"),
			Date:    fs.String("date", "", "日期, 例如: 2023-09"),
		}
	},
	ValidArgs: []cobra.PositionalArgs{
		cobra.ExactArgs(0),
		func(cmd *cobra.Command, args []string) error {
			date, _ := cmd.Flags().GetString("date")
			if date != "" {
				if err := stringutils.MustMatch(date, "^[0-9]+-[0-9]+$"); err != nil {
					return fmt.Errorf("invalid flag 'date', %s", err)
				}
			}
			return nil
		},
	},
	Run: func(args []string, flags bingImgWdbyteCmdFlags) error {
		reqUrl := fmt.Sprintf("%s/%s", URL_BING_WDBYTE_COM, *flags.Date)
		console.Info("解析页面 %s", reqUrl)
		doc := httpLib.GetHtml(reqUrl)
		if doc == nil {
			console.Warn("页面 %s 请求失败", reqUrl)
			os.Exit(1)
		}
		files := []httpLib.HttpFile{}
		doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			href := s.AttrOr("href", "")
			if err := stringutils.MustMatch(href, `.*id=.+\.jpg`); err == nil {
				if *flags.Limit > 0 && len(files) >= *flags.Limit {
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
		downloader := httpLib.HttpDownloader{Output: *flags.Output}

		syncutils.StartTasks(
			syncutils.TaskOption{
				TaskName:     fmt.Sprintf("download %d picture(s)", len(files)),
				MaxWorker:    *flags.Workers,
				ShowProgress: true,
				PreStart: func() {
					console.Info("开始下载(总数: %d), 保存路径: %s ...", len(files), *flags.Output)
				},
			},
			files,
			func(file httpLib.HttpFile) error {
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
		)

		result, _ := table.NewItemsTable(
			[]string{"Result", "Name", "Size", "Spend"}, webFiles,
		).SetStyle(
			table.StyleLight, color.FgCyan,
		).EnableAutoIndex().Render()
		fmt.Println(result)
		return nil
	},
})

func init() {
	FetchWallpaperCmd.AddCommand(BingImgCmd, BingImgWdbyteCmd)
}
