package commands

import (
	"github.com/spf13/cobra"

	"github.com/fjboy/magic-pocket/pkg/global/logging"
	localHttp "github.com/fjboy/magic-pocket/pkg/http"
)

func GetHTTPServerCommand() cobra.Command {
	var (
		port int16
	)

	var command = &cobra.Command{
		Use:   "http-server",
		Short: "简单的 HTTP 服务器",
		Long:  "启动一个简单的 HTTP 服务器",
		Run: func(cmd *cobra.Command, args []string) {
			logging.BasicConfig(logging.LogConfig{
				Level: logging.INFO,
			})
			logging.Info("启动服务 :%d", port)
			localHttp.SimpleHttpServer(port)
		},
	}
	command.Flags().Int16VarP(&port, "port", "p", 80, "监听端口")
	return *command
}
