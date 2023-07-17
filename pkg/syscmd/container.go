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

type Docker struct {
	cmd string
}

type containerCli ContainerCli

func (ctl Docker) GetImageRepo(repo string, image string) string {
	repo = strings.TrimSuffix(repo, "/")
	image = strings.TrimPrefix(image, "/")
	return fmt.Sprintf("%s/%s", repo, image)
}
func (ctl Docker) Pull(image string) error {
	_, err := GetOutput(ctl.cmd, "pull", image)
	if err != nil {
		return err
	}
	return nil
}
func (ctl Docker) Push(image string) error {
	_, err := GetOutput(ctl.cmd, "push", image)
	if err != nil {
		return err
	}
	return nil
}
func (ctl Docker) Tag(image string, tag string) error {
	_, err := GetOutput(ctl.cmd, "tag", image, tag)
	if err != nil {
		return err
	}
	return nil
}
func (ctl Docker) GetCmd() string {
	return ctl.cmd
}

type Podman struct {
	Docker
	cmd string
}

func NewDocker() *Docker {
	return &Docker{
		cmd: "docker",
	}
}

func NewPodman() *Podman {
	return &Podman{
		cmd: "podman",
	}
}

var SUPPORT_CONTAINER_CLIENT = []string{
	"podman", "docker",
}

// func (ctl *ContainerCtl) GetImageRepo(repo string, image string) string {
// 	repo = strings.TrimSuffix(repo, "/")
// 	image = strings.TrimPrefix(image, "/")
// 	return fmt.Sprintf("%s/%s", repo, image)
// }

// func (ctl *ContainerCtl) Pull(image string) error {
// 	_, err := GetOutput(ctl.Cmd, "pull", image)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (ctl *ContainerCtl) Push(image string) error {
// 	_, err := GetOutput(ctl.Cmd, "push", image)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (ctl *ContainerCtl) Tag(image string, tag string) error {
// 	_, err := GetOutput(ctl.Cmd, "tag", image, tag)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetContainerCtl() (ContainerCtl, error) {
// 	for _, client := range SUPPORT_CONTAINER_CLIENT {
// 		_, err := GetOutput(client, "-v")
// 		if err == nil {
// 			return ContainerCtl{Cmd: client}, nil
// 		}
// 	}
// 	return ContainerCtl{}, fmt.Errorf("无容器客户端，支持的客户端: %v", SUPPORT_CONTAINER_CLIENT)
// }

func GetDefaultContainerCli() (ContainerCli, error) {
	for _, client := range SUPPORT_CONTAINER_CLIENT {
		_, err := GetOutput(client, "-v")
		if err == nil {
			if client == "podman" {
				return Podman{cmd: "podman"}, nil
			}
			return &Docker{cmd: "docker"}, nil
		}
	}
	return nil, fmt.Errorf("无容器客户端，支持的客户端: %v", SUPPORT_CONTAINER_CLIENT)
}
