package commands

import (
	"github.com/spf13/cobra"

	"github.com/fjboy/magic-pocket/pkg/global/logging"
	localHttp "github.com/fjboy/magic-pocket/pkg/http"
)

var (
	port int16
)
var SimpleHttpServer = &cobra.Command{
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

func init() {
	SimpleHttpServer.Flags().Int16VarP(&port, "port", "p", 80, "监听端口")
}
