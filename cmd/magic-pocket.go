package main

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/fjboy/magic-pocket/cmd/commands"
)

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
		Long:  "从 www.bingimg.cn 网站下载高质量壁纸",
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
			commands.BingImgDownload(page, uhd, output)
		},
	}
	bingImgDownload.Flags().BoolVar(&uhd, "uhd", false, "仅下载4K壁纸")
	bingImgDownload.Flags().StringVarP(&output, "output", "o", "./", "保存路径")
	rootCmd.AddCommand(bingImgDownload)
	rootCmd.Execute()
}
