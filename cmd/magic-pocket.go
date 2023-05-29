package main

import (
	"github.com/spf13/cobra"

	"github.com/fjboy/magic-pocket/cmd/commands"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "magic-pocket",
		Short: "常用工具合集",
		Long:  "Golang 实现的工具合集",
	}
	bingImgDownload := commands.GetCommand()

	rootCmd.AddCommand(&bingImgDownload)
	rootCmd.Execute()
}
