package commands

import (
	"github.com/BytemanD/easygo/pkg/global/logging"
	myHttp "github.com/BytemanD/easygo/pkg/http"
	"github.com/spf13/cobra"
)

var HttpFS = &cobra.Command{
	Use:   "httpfs",
	Short: "一个简单的 HTTP 文件服务器",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt16("port")
		root, _ := cmd.Flags().GetString("root")
		server := myHttp.VuetifyHttpFS{
			Port:       port,
			Root:       root,
			StaticPath: "./static",
		}
		if err := server.Run(); err != nil {
			logging.Fatal("启动HTTP服务失败: %s", err)
		}
	},
}

func init() {
	HttpFS.Flags().Int16("port", 8080, "监听端口")
	HttpFS.Flags().String("root", "./", "根目录")
}
