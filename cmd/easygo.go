package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/BytemanD/easygo/cmd/commands"
	"github.com/BytemanD/easygo/pkg/global/gitutils"
	"github.com/BytemanD/easygo/pkg/global/logging"
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
			level := logging.INFO
			debug, _ := cmd.Flags().GetBool("debug")
			if debug {
				level = logging.DEBUG
			}
			logFile, _ := cmd.Flags().GetString("log-file")
			logging.BasicConfig(logging.LogConfig{Level: level, EnableFileLine: true, Output: logFile})
		},
	}

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "显示Debug信息")
	rootCmd.PersistentFlags().String("log-file", "", "日志文件")

	rootCmd.AddCommand(
		commands.FetchWallpaperCmd,
		commands.WgetLinks,
		commands.SimpleHttpFS,
		commands.IniCrud,
		commands.ContainerImageSync,
		commands.MDRender,
		commands.Wget,
		commands.CSVRender,
	)

	rootCmd.Execute()
}
