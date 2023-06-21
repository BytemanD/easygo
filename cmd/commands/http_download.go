package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	httpUtils "github.com/BytemanD/easygo/pkg/http"
	"github.com/BytemanD/easygo/pkg/stringutils"
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
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		if err := stringutils.MustMatch(args[0], "^http(s)*://.+"); err != nil {
			return fmt.Errorf("invalid flag 'url': %s", err)
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
