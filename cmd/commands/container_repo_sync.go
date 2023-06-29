package commands

import (
	"os"

	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/BytemanD/easygo/pkg/syscmd"
	"github.com/spf13/cobra"
)

var (
	destRepos []string
)
var ContainerImageSync = &cobra.Command{
	Use:   "container-repo-sync <source repo> <image>",
	Short: "同步容器镜像",
	Long:  "同步两个镜像仓库之间的镜像",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(2)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		srcRepo := args[0]
		image := args[1]

		ctl, err := syscmd.GetContainerCtl()
		if err != nil {
			logging.Error("获取客户端失败: %s", err)
			os.Exit(1)
		}
		logging.Info("使用容器客户端: %s", ctl.Cmd)
		srcImage := ctl.GetImageRepo(srcRepo, image)

		logging.Info("拉取镜像: %s", srcImage)
		if err := ctl.Pull(srcImage); err != nil {
			logging.Error("拉取失败: %s", err)
			os.Exit(1)
		}
		for _, destRepo := range destRepos {
			destImage := ctl.GetImageRepo(destRepo, image)
			if err := ctl.Tag(srcImage, destImage); err != nil {
				logging.Error("Tag 设置失败: %s", err)
				os.Exit(1)
			}
			logging.Info("推送镜像: %s", destImage)
			if err := ctl.Push(destImage); err != nil {
				logging.Error("失败: %s", err)
				os.Exit(1)
			}
		}
		logging.Info("同步完成")
	},
}

func init() {
	// ContainerImageSync.Flags().StringVarP(&destRepo, "dest", "D", "", "目标镜像仓库地址,例如: docker.io")
	ContainerImageSync.Flags().StringArrayVarP(
		&destRepos, "dest", "D", []string{}, "目标镜像仓库地址,例如: docker.io")
}
