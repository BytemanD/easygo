package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

func GetVersionCommand() cobra.Command {

	var command = &cobra.Command{
		Use:   "version",
		Short: "获取版本",
		Long:  "获取工具版本",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Version)
		},
	}
	return *command
}
