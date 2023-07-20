package syscmd

import (
	"fmt"
	"strings"
)

type ContainerCli interface {
	Pull(image string) error
	Push(image string) error
	Tag(image string, tag string) error
	GetImageRepo(repo string, image string) string
	GetCmd() string
}

type Impl struct {
	cmd string
}

func (ctl Impl) GetImageRepo(repo string, image string) string {
	repo = strings.TrimSuffix(repo, "/")
	image = strings.TrimPrefix(image, "/")
	return fmt.Sprintf("%s/%s", repo, image)
}
func (ctl Impl) Pull(image string) error {
	_, err := GetOutput(ctl.GetCmd(), "pull", image)
	if err != nil {
		return err
	}
	return nil
}
func (ctl Impl) Push(image string) error {
	_, err := GetOutput(ctl.GetCmd(), "push", image)
	if err != nil {
		return err
	}
	return nil
}
func (ctl Impl) Tag(image string, tag string) error {
	_, err := GetOutput(ctl.GetCmd(), "tag", image, tag)
	if err != nil {
		return err
	}
	return nil
}
func (ctl Impl) GetCmd() string {
	return ctl.cmd
}

func NewDocker() *Impl {
	return &Impl{cmd: "docker"}
}

func NewPodman() *Impl {
	return &Impl{cmd: "podman"}
}

var SUPPORT_CONTAINER_CLIENT = []string{
	"podman", "docker",
}

func GetDefaultContainerCli() (ContainerCli, error) {
	for _, client := range SUPPORT_CONTAINER_CLIENT {
		if _, err := GetOutput(client, "-v"); err != nil {
			continue
		}
		if client == "podman" {
			return NewPodman(), nil
		} else {
			return NewDocker(), nil
		}
	}
	return nil, fmt.Errorf("无容器客户端，支持的客户端: %v", SUPPORT_CONTAINER_CLIENT)
}
