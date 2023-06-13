package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/BytemanD/easygo/pkg/global/ini"
)

var IniCrud = &cobra.Command{
	Use:   "crud-ini",
	Short: "ini 配置文件编辑器",
}
var IniGet = &cobra.Command{
	Use:   "get <file path> <section> <key>",
	Short: "获取配置",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		filePath, section, key := args[0], args[1], args[2]
		err := ini.CONF.Load(filePath)
		if err != nil {
			fmt.Printf("配置文件加载异常, %s\n", err)
			os.Exit(1)
		}
		ini.CONF.SetBlockMode(true)
		fmt.Println(ini.CONF.Get(section, key))
	},
}
var IniSet = &cobra.Command{
	Use:   "set <file path> <section> <key> <value>",
	Short: "添加/更新配置",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		filePath, section, key, value := args[0], args[1], args[2], args[3]
		err := ini.CONF.Load(filePath)
		if err != nil {
			fmt.Printf("配置文件加载异常, %s\n", err)
			os.Exit(1)
		}
		ini.CONF.Set(section, key, value)
	},
}
var IniDelete = &cobra.Command{
	Use:   "del <file path> <section> <key>",
	Short: "获取配置",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		filePath, section, key := args[0], args[1], args[2]
		err := ini.CONF.Load(filePath)
		if err != nil {
			fmt.Printf("配置文件加载异常, %s\n", err)
			os.Exit(1)
		}
		ini.CONF.Delete(section, key)
	},
}

func init() {
	IniCrud.AddCommand(IniGet)
	IniCrud.AddCommand(IniSet)
	IniCrud.AddCommand(IniDelete)
}
