package commands

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"

	httpUtils "github.com/fjboy/magic-pocket/pkg/http"
)

var (
	regex      string
	outputPath string
	direct     bool
)
var HttpDownloadCmd = &cobra.Command{
	Use:   "http-download <url>",
	Short: "HTTP下载器",
	Long:  "下载指定http地址连接",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires a url argument")
		}
		url := args[0]
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			return errors.New("url must starts with http:// or https://")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		url := strings.TrimRight(args[0], "/")
		if direct {
			httpUtils.Download(url, outputPath)
		} else {
			httpUtils.DownloadLinksInHtml(url, regex, outputPath)
		}
	},
}

func init() {
	HttpDownloadCmd.Flags().StringVarP(&regex, "regex", "r", "", "匹配正则表达式，例如: .rpm")
	HttpDownloadCmd.Flags().StringVarP(&outputPath, "output", "o", "./", "保存路径")
	HttpDownloadCmd.Flags().BoolVar(&direct, "direct", false, "直接下载连接")
}
