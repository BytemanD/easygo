package commands

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"

	"github.com/BytemanD/easygo/pkg/global/logging"
	httpUtils "github.com/BytemanD/easygo/pkg/http"
	"github.com/BytemanD/easygo/pkg/stringutils"
)

const SCHEME string = "http"
const HOST string = "www.bingimg.cn"
const FILE_NAME_MAX_SIZE int = 50
const URL_GET_IMAGES_PAGE string = "%s://%s/list%s"

const UHD_ONLY string = "only"
const UHD_INCLUDE string = "include"
const UHD_NO string = "no"

var UHD_CHOICES = []string{UHD_ONLY, UHD_INCLUDE, UHD_NO}

func bingImgDownload(page int8, uhd string, output string) {
	logging.Info("下载页面 %d, 保存位置: %s", page, output)
	url := fmt.Sprintf(URL_GET_IMAGES_PAGE, SCHEME, HOST, strconv.Itoa(int(page)))
	doc := httpUtils.GetHtml(url)
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
	for link := links.Front(); link != nil; link = link.Next() {
		logging.Debug("下载 %s", link.Value)
		httpUtils.Download(fmt.Sprintf("%s", link.Value), output)
	}
}

func bingImgDownloadPages(page int8, endPage int8, uhd string, output string) {
	if endPage <= 0 {
		endPage = page
	}
	for i := page; i <= endPage; i++ {
		bingImgDownload(i, uhd, output)
	}
}

var (
	uhd     string
	output  string
	endPage int8
)

var BingImgDownloadCmd = &cobra.Command{
	Use:              "get-bing-img <page>",
	Short:            "下载bing高质量壁纸",
	Long:             "下载 www.bingimg.cn 网站下的高质量壁纸",
	TraverseChildren: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		if _, err := stringutils.MustGreaterThan(args[0], 1); err != nil {
			return fmt.Errorf("invalid arg 'page': %s", err)
		}
		if err := stringutils.MustInStringChoises(uhd, UHD_CHOICES); err != nil {
			return fmt.Errorf("invalid flag 'uhd': %s", err)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		page, _ := strconv.Atoi(args[0])
		bingImgDownloadPages(int8(page), endPage, uhd, output)
	},
}

func init() {
	BingImgDownloadCmd.Flags().StringVarP(&uhd, "uhd", "u", "only", "下载4K分辨率")
	BingImgDownloadCmd.Flags().StringVarP(&output, "output", "o", "./", "保存路径")
	BingImgDownloadCmd.Flags().Int8VarP(&endPage, "end-page", "e", 0, "结束的页面, 默认和page相同")
}
