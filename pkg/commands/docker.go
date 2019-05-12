package commands

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jesseduffield/lazydocker/pkg/config"
	"github.com/jesseduffield/lazydocker/pkg/i18n"
	"github.com/sirupsen/logrus"
)

// DockerCommand is our main git interface
type DockerCommand struct {
	Log       *logrus.Entry
	OSCommand *OSCommand
	Tr        *i18n.Localizer
	Config    config.AppConfigurer
	Client    *client.Client
}

// NewDockerCommand it runs git commands
func NewDockerCommand(log *logrus.Entry, osCommand *OSCommand, tr *i18n.Localizer, config config.AppConfigurer) (*DockerCommand, error) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		return nil, err
	}

	return &DockerCommand{
		Log:       log,
		OSCommand: osCommand,
		Tr:        tr,
		Config:    config,
		Client:    cli,
	}, nil
}

// GetContainers returns a slice of docker containers
func (c *DockerCommand) GetContainers() ([]*Container, error) {
	containers, err := c.Client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	ownContainers := make([]*Container, len(containers))

	for i, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
		ownContainers[i] = &Container{ID: container.ID, Name: "test", Container: container}
	}

	return ownContainers, nil
}
