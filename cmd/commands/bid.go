package commands

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	httpUtils "github.com/fjboy/magic-pocket/pkg/http"
)

const SCHEME string = "http"
const HOST string = "www.bingimg.cn"
const FILE_NAME_MAX_SIZE int = 50
const URL_GET_IMAGES_PAGE string = "%s://%s/list%s"

func BingImgDownload(page int, uhd bool, output string) {
	log.Printf("下载页面 %d, 保存位置: %s", page, output)
	url := fmt.Sprintf(URL_GET_IMAGES_PAGE, SCHEME, HOST, strconv.Itoa(page))
	doc := httpUtils.GetHtml(url)
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
		httpUtils.Download(fmt.Sprintf("%s", link.Value), output)
	}
}
