package commands

import (
	"github.com/spf13/cobra"

	"github.com/BytemanD/easygo/pkg/global/logging"
	myHttp "github.com/BytemanD/easygo/pkg/http"
)

var (
	port int16
	root string
)

var SimpleHttpFS = &cobra.Command{
	Use:   "fs-server",
	Short: "一个简单的 HTTP 文件服务器",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		myHttp.FSConfig.Port = port
		myHttp.FSConfig.Root = root
		err := myHttp.SimpleHttpFS()
		if err != nil {
			logging.Fatal("启动HTTP服务失败: %s", err)
		}
	},
}

func init() {
	SimpleHttpFS.Flags().Int16VarP(&port, "port", "p", 80, "监听端口")
	SimpleHttpFS.Flags().StringVarP(&root, "root", "r", ".", "根目录")
}
