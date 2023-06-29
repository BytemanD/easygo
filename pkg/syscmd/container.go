package syscmd

import (
	"fmt"
	"strings"
)

type ContainerCtl struct {
	Cmd string
}

var SUPPORT_CONTAINER_CLIENT = []string{
	"podman", "docker",
}

func (ctl *ContainerCtl) GetImageRepo(repo string, image string) string {
	repo = strings.TrimSuffix(repo, "/")
	image = strings.TrimPrefix(image, "/")
	return fmt.Sprintf("%s/%s", repo, image)
}

func (ctl *ContainerCtl) Pull(image string) error {
	_, err := GetOutput(ctl.Cmd, "pull", image)
	if err != nil {
		return err
	}
	return nil
}
func (ctl *ContainerCtl) Push(image string) error {
	_, err := GetOutput(ctl.Cmd, "push", image)
	if err != nil {
		return err
	}
	return nil
}
func (ctl *ContainerCtl) Tag(image string, tag string) error {
	_, err := GetOutput(ctl.Cmd, "tag", image, tag)
	if err != nil {
		return err
	}
	return nil
}

func GetContainerCtl() (ContainerCtl, error) {
	for _, client := range SUPPORT_CONTAINER_CLIENT {
		_, err := GetOutput(client, "-v")
		if err == nil {
			return ContainerCtl{Cmd: client}, nil
		}
	}
	return ContainerCtl{}, fmt.Errorf("无容器客户端，支持的客户端: %v", SUPPORT_CONTAINER_CLIENT)
}
