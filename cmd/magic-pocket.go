package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/fjboy/magic-pocket/cmd/commands"
	"github.com/fjboy/magic-pocket/pkg/global/gitutils"
)

var Version string

func GetVersionCommand() cobra.Command {

	var command = &cobra.Command{
		Use:   "version",
		Short: "版本",
		Long:  "获取工具版本",
		Run: func(cmd *cobra.Command, args []string) {
			if Version == "" {
				Version = gitutils.GetVersion()
			}
			fmt.Println(Version)
		},
	}
	return *command
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "magic-pocket",
		Short: "常用工具合集",
		Long:  fmt.Sprintf("Golang 实现的工具合集(版本: %s)", Version),
	}

	bingImgDownload := commands.GetCommand()
	simpleHttpServer := commands.GetHTTPServerCommand()
	version := GetVersionCommand()

	rootCmd.AddCommand(&bingImgDownload)
	rootCmd.AddCommand(&simpleHttpServer)
	rootCmd.AddCommand(&version)
	rootCmd.Execute()
}
