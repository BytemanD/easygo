package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/BytemanD/easygo/cmd/commands"
	"github.com/BytemanD/easygo/pkg/global/gitutils"
	"github.com/BytemanD/go-console/console"
)

var Version string

func getVersion() string {
	if Version == "" {
		return gitutils.GetVersion()
	}
	return fmt.Sprint(Version)
}

func main() {
	var rootCmd = &cobra.Command{
		Use:     "easygo",
		Short:   "Golang 工具集",
		Long:    "Golang 实现的工具合集",
		Version: getVersion(),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			if debug {
				console.EnableLogDebug()
			}
			logFile, _ := cmd.Flags().GetString("log-file")
			if logFile != "" {
				console.SetLogFile(logFile)
			}
		},
	}

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "显示Debug信息")
	rootCmd.PersistentFlags().String("log-file", "", "日志文件")
	rootCmd.PersistentFlags().Bool("enable-log-color", false, "启用日志颜色")

	rootCmd.AddCommand(
		commands.FetchWallpaperCmd,
		commands.WgetLinks,
		commands.HttpFS,
		commands.IniCrud,
		commands.ContainerImageSync,
		commands.MDRender,
		commands.Wget,
		commands.CSVRender,
	)

	rootCmd.Execute()
}
