// Command Part of AetCLI.
// TODO: Move to AetCLI-Admin.
// This command will present the layout of the platform through the CLI.
package aetcli

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func SystemOverview() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer cli.Close()

	// List the containers.
	containers, err := cli.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, container := range containers {
		fmt.Println(container.Names)
	}
}
