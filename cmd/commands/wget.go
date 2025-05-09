package commands

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/BytemanD/easygo/pkg/http"
	"github.com/BytemanD/go-console/console"

	"github.com/BytemanD/easygo/pkg/progress"
	"github.com/BytemanD/easygo/pkg/stringutils"
)

type ProgressWriter struct {
	Writer *bufio.Writer
	bar    *progress.ProgressBar
}

func (pw ProgressWriter) Write(p []byte) (n int, err error) {
	pw.bar.Increment(len(p))
	return pw.Writer.Write(p)
}
func (pw ProgressWriter) Flush() error {
	return pw.Writer.Flush()
}

var Wget = &cobra.Command{
	Use:   "wget",
	Short: "get web file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		output, _ := cmd.Flags().GetString("output")
		quiet, _ := cmd.Flags().GetBool("quiet")

		console.Debug("download %s", url)
		err := http.Download(url, output, !quiet)
		if err != nil {
			console.Error("download %s failed: %s", url, err)
			return
		}
	},
}
var WgetLinks = &cobra.Command{
	Use:   "wget-links <url>",
	Short: "下载网页上的链接",
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
		regex, _ := cmd.Flags().GetString("regex")
		output, _ := cmd.Flags().GetString("output")
		http.DownloadLinksInHtml(url, regex, output)
	},
}

func init() {
	Wget.Flags().StringP("output", "O", "./", "保存路径")
	Wget.Flags().BoolP("quiet", "q", false, "quiet")

	WgetLinks.Flags().StringP("regex", "r", "", "匹配正则表达式，例如: .rpm")
	WgetLinks.Flags().StringP("output", "O", "./", "保存路径")
}
