package commands

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"

	"github.com/fjboy/magic-pocket/pkg/global/logging"
	httpUtils "github.com/fjboy/magic-pocket/pkg/http"
)

func GetHttpDownloadCommand() cobra.Command {
	var (
		regex  string
		output string
		direct bool
		debug  bool
	)

	var command = &cobra.Command{
		Use:   "http-download",
		Short: "http 下载",
		Long:  "指定http地址下载",
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
			level := logging.INFO
			if debug {
				level = logging.DEBUG
			}
			logging.BasicConfig(logging.LogConfig{Level: level})
			url := strings.TrimRight(args[0], "/")
			if direct {
				httpUtils.Download(url, output)
			} else {
				httpUtils.DownloadLinksInHtml(url, regex, output)
			}
		},
	}
	command.Flags().StringVarP(&regex, "regex", "r", "", "匹配正则表达式，例如: .rpm")
	command.Flags().StringVarP(&output, "output", "o", "./", "保存路径")
	command.Flags().BoolVar(&direct, "direct", false, "直接下载连接")
	command.Flags().BoolVarP(&debug, "debug", "d", false, "显示Debug信息")
	return *command
}
